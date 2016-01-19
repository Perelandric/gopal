package gopal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"time"
)

func (self *connection) FetchPayment(payment_id string) (*Payment, error) {
	var pymt = &Payment{
		connection: self,
	}
	if err := self.send(&request{
		method:   method.get,
		path:     path.Join(_paymentsPath, payment_id),
		body:     nil,
		response: pymt,
	}); err != nil {
		return nil, err
	}
	return pymt, nil
}

func CreatePayment(
	conn *connection,
	method PaymentMethodEnum,
	urls RedirectUrls) (*Payment, error) {

	if method == PaymentMethod.PayPal {
		// Make sure we're still authenticated. Will refresh if not.
		if err := conn.authenticate(); err != nil {
			return nil, err
		}

		return &Payment{
			payment_request: payment_request{
				Intent: intent.sale,
				Payer: payer{
					PaymentMethod: method,
				},
				Transactions: make([]*transaction, 0),
				RedirectUrls: urls,
			},
		}, nil
	}
	return nil, nil
}

func (self *Payment) Execute(req *http.Request) error {
	var query = req.URL.Query()

	// TODO: Is this right? Does URL.Query() ever return nil?
	if query == nil {
		return fmt.Errorf("Attempt to execute a payment that has not been approved")
	}

	var payerid = query.Get("PayerID")
	if payerid == "" {
		return fmt.Errorf("PayerID is missing\n")
	}

	var pymtid = query.Get("paymentId")
	if pymtid == "" {
		return fmt.Errorf("paymentId is missing\n")
	}

	var pymt Payment

	if err := self.send(&request{
		method:   method.post,
		path:     path.Join(_paymentsPath, pymtid, _execute),
		body:     &paymentExecution{PayerID: payerid},
		response: &pymt,
	}); err != nil {
		return err
	}

	if pymt.State != state.Approved {
		return fmt.Errorf(
			"Payment with ID %q for payer %q was not approved\n", pymtid, payerid)
	}

	return nil
}

/***************************

	Payment object methods

***************************/

func (self *Payment) AddTransaction(t transaction) {
	self.Transactions = append(self.Transactions, &t)
}

// TODO: The tkn parameter is ignored, but should send a query string parameter `token=tkn`
func (self *Payment) Authorize(tkn string) (to string, code int, err error) {
	err = self.send(&request{
		method:   method.post,
		path:     _paymentsPath,
		body:     self,
		response: self,
	})

	if err == nil {
		switch self.State {
		case state.Created:
			// Set url to redirect to PayPal site to begin approval process
			to, _ = self.Links.get(relType.approvalUrl)
			code = 303
		default:
			// otherwise cancel the payment and return an error
			err = UnexpectedResponse
		}
	}

	return to, code, err
}

// TODO: I'm returning a list of Sale objects because every payment can have multiple transactions.
//		I need to find out why a payment can have multiple transactions, and see if I should eliminate that in the API
//		Also, I need to find out why `related_resources` is an array. Can there be more than one per type?
func (self *Payment) GetSale() []*Sale {
	var sales = []*Sale{}
	for _, transaction := range self.Transactions {
		for _, related_resource := range transaction.RelatedResources {
			if s, ok := related_resource.(*Sale); ok {
				sales = append(sales, s)
			}
		}
	}
	return sales
}

// Pagination
//	Assuming `start_time`, `start_index` and `start_id` are mutually exclusive
//	...going to treat them that way anyhow until I understand better.

// I'm going to ignore `start_index` for now since I don't see its usefulness

func (self *connection) GetAllPayments(
	size int,
	sort_by sortByEnum, sort_order sortOrderEnum, time_range ...time.Time,
) *PaymentBatcher {

	if size < 0 {
		size = 0
	} else if size > 20 {
		size = 20
	}

	var qry = fmt.Sprintf("?sort_order=%s&sort_by=%s&count=%d", sort_by, sort_order, size)

	if len(time_range) > 0 {
		if time_range[0].IsZero() == false {
			qry = fmt.Sprintf("%s&start_time=%s", qry, time_range[0].Format(time.RFC3339))
		}
		if len(time_range) > 1 && time_range[1].After(time_range[0]) {
			qry = fmt.Sprintf("%s&end_time=%s", qry, time_range[1].Format(time.RFC3339))
		}
	}

	return &PaymentBatcher{
		connection: self,
		base_query: qry,
		next_id:    "",
		done:       false,
	}
}

// These provide a way to both get and set the `next_id`.
// This gives the ability to cache the ID, and then set it in a new Batcher.
// Useful if a session is not desired or practical

func (self *PaymentBatcher) GetNextId() string {
	return self.next_id
}
func (self *PaymentBatcher) SetNextId(id string) {
	self.next_id = id
}

type payment_list struct {
	Payments []*Payment `json:"payments"`
	Count    int        `json:"count"`
	Next_id  string     `json:"next_id"`
	*identity_error
}

type payment_request struct {
	// Payment intent. Must be set to sale for immediate payment, authorize to
	// authorize a payment for capture later, or order to create an order.
	// Allowed values: sale, authorize, order.
	Intent intentEnum `json:"intent,omitempty"`

	Payer payer `json:"payer,omitempty"`

	// Transactional details including the amount and item details. REQUIRED.
	Transactions []*transaction `json:"transactions,omitempty"`

	// Set of redirect URLs you provide only for PayPal-based payments. Returned
	// only when the payment is in created state. REQUIRED for PayPal payments.
	RedirectUrls RedirectUrls `json:"redirect_urls,omitempty"`
}

type RedirectUrls struct {
	ReturnUrl string `json:"return_url,omitempty"`
	CancelUrl string `json:"cancel_url,omitempty"`
}

type Payment struct {
	*connection
	payment_request

	// Payment creation time as defined in RFC 3339 Section 5.6. and the
	// time that the resource was last updated. Values assigned by PayPal.
	_times

	// ID of the created payment. Value assigned by PayPal.
	Id string `json:"id,omitempty"`

	// Payment state. Must be one of the following: created; approved; failed;
	// pending; canceled; expired, or in_progress. Value assigned by PayPal.
	State stateEnum `json:"state,omitempty"`

	// Identifier for the payment experience.
	ExperienceProfileId string `json:"experience_profile_id"`

	// HATEOAS links related to this call.
	Links links `json:"links,omitempty"`

	*payment_error
}

type paymentExecution struct {
	// The ID of the Payer, passed in the return_url by PayPal.
	PayerID string `json:"payer_id,omitempty"`

	// Transactional details if updating a payment. Note that this instance of
	// the transactions object accepts only the amount object.
	Transactions []*update_transaction `json:"transactions,omitempty"`
}

// Amount Object
//  A`Transaction` object also may have an `ItemList`, which has dollar amounts.
//  These amounts are used to calculate the `Total` field of the `Amount` object
//
//	All other uses of `Amount` do have `shipping`, `shipping_discount` and
// `subtotal` to calculate the `Total`.
type amount struct {
	// 3 letter currency code. PayPal does not support all currencies. REQUIRED.
	Currency CurrencyTypeEnum `json:"currency"`

	// Total amount charged from the payer to the payee. In case of a refund, this
	// is the refunded amount to the original payer from the payee. 10 characters
	// max with support for 2 decimal places. REQUIRED.
	Total string `json:"total"`

	Details *details `json:"details,omitempty"`
}

func (self *amount) setTotal(amt float64) error {
	if s, err := make10CharAmount(amt); err != nil {
		return err
	} else {
		self.Total = s
	}
	return nil
}

type details struct {
	// Amount charged for shipping. 10 chars max, with support for 2 decimal places
	Shipping float64 `json:"shipping,omitempty"`

	// Amount of the subtotal of the items. REQUIRED if line items are specified.
	// 10 chars max, with support for 2 decimal places
	Subtotal float64 `json:"subtotal,omitempty"`

	// Amount charged for tax. 10 chars max, with support for 2 decimal places
	Tax float64 `json:"tax,omitempty"`

	// Amount being charged for handling fee. When `payment_method` is `paypal`
	HandlingFee float64 `json:"handling_fee,omitempty"`

	// Amount being charged for insurance fee. When `payment_method` is `paypal`
	Insurance float64 `json:"insurance,omitempty"`

	// Amount being discounted for shipping fee. When `payment_method` is `paypal`
	ShippingDiscount float64 `json:"shipping_discount,omitempty"`
}

type itemList struct {
	Items           []*item          `json:"items,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

type item struct {
	// REQUIRED. Number of a particular item. 10 chars max.
	Quantity int `json:"quantity,string"`

	// REQUIRED. Item name. 127 chars max.
	Name string `json:"name"`

	// REQUIRED. Item cost. 10 chars max.
	Price float64 `json:"price,string"`

	// REQUIRED. 3-letter currency code.
	Currency CurrencyTypeEnum `json:"currency"`

	// Stock keeping unit corresponding (SKU) to item. 50 chars max.
	Sku string `json:"sku,omitempty"`

	// Description of the item. Only supported when the `payment_method` is `paypal`
	// 127 characters max.
	Description string `json:"description,omitempty"`

	// Tax of the item. Only supported when `payment_method` is `paypal`.
	Tax float64 `json:"tax,omitempty"`
}

// This is just used for the payment_execution, since it only needs an amount.
type update_transaction struct {
	// Amount being collected. REQUIRED.
	Amount amount `json:"amount"`
}

const descMax = 127

type transaction struct {
	update_transaction

	// Description of transaction. 127 characters max.
	Description string `json:"description,omitempty"`

	// Items and related shipping address within a transaction.
	ItemList *itemList `json:"item_list,omitempty"`

	// Financial transactions related to a payment. (Only in PayPal responses?)
	// Sale, Authorization, Capture, or Refund objects
	RelatedResources RelatedResources `json:"related_resources,omitempty"`

	// Invoice number used to track the payment. Only supported when the
	// `payment_method` is set to `paypal`. 256 characters max.
	InvoiceNumber string `json:"invoice_number,omitempty"`

	// Free-form field for the use of clients. Only supported when the
	// `payment_method` is set to `paypal`. 256 characters max.
	Custom string `json:"custom,omitempty"`

	// Soft descriptor used when charging this funding source. Only supported when
	// the `payment_method` is set to `paypal`. 22 characters max.
	SoftDescriptor string `json:"soft_descriptor,omitempty"`

	// Payment options requested for this purchase unit.
	PaymentOptions paymentOptions `json:"payment_options,omitempty"`
}

func NewTransaction(
	c CurrencyTypeEnum, desc string, shp *ShippingAddress) *transaction {

	if len(desc) > descMax { // Log and truncate if description is too long
		log.Printf("Description exceeds %d characters: %q\n", descMax, desc)
		desc = desc[0:descMax]
	}

	var t = &transaction{
		update_transaction: update_transaction{
			Amount: amount{
				Currency: c,
				Total:    "0",
			},
		},
		Description: desc,
	}

	if shp != nil {
		t.ItemList = &itemList{
			ShippingAddress: shp,
		}
	}

	return t
}

func (t *transaction) AddItem(
	qty int, price float64, curr CurrencyTypeEnum, name, sku string) {

	if qty < 1 {
		qty = 1
	}

	if t.ItemList == nil {
		t.ItemList = new(itemList)
	}
	t.ItemList.Items = append(t.ItemList.Items, &item{
		Quantity: qty,
		Name:     name,
		Price:    price,
		Currency: curr,
		Sku:      sku,
	})
}

type RelatedResources []Resource

//func (self *RelatedResources) MarshalJSON() ([]byte, error) {
//	return []byte("[]"), nil
//}

func (self *RelatedResources) UnmarshalJSON(b []byte) error {
	if self == nil || len(*self) == 0 {
		return nil
	}

	var a = []map[string]json.RawMessage{}
	err := json.Unmarshal(b, &a)
	if err != nil {
		return err
	}

	for _, m := range a {
		for name, rawMesg := range m {
			var t Resource // for unmarshaling the current item

			switch name {
			case SaleType:
				t = new(Sale)
			case AuthorizationType:
				t = new(Authorization)
			case CaptureType:
				t = new(Capture)
			case RefundType:
				t = new(Refund)
			default:
				log.Printf("Unexpected Resource type: %s\n", name)
				continue
			}
			if err = json.Unmarshal(rawMesg, t); err != nil {
				return err
			}

			*self = append(*self, t) // Add unmarshaled item
		}
	}
	return nil
}

//This object includes payment options requested for the purchase unit.
type paymentOptions struct {
	// Optional payment method type. If specified, the transaction will go through
	// for only instant payment. Allowed values: `INSTANT_FUNDING_SOURCE`. Only for
	// use with the `paypal` payment_method, not relevant for the `credit_card`
	// payment_method.
	AllowedPaymentMethod string `json:"allowed_payment_method,omitempty"`
}

/****************************************

	PaymentBatcher

Manages paginated requests for Payments

*****************************************/

type PaymentBatcher struct {
	*connection
	base_query string
	next_id    string
	done       bool
}

func (self *PaymentBatcher) IsDone() bool {
	return self.done
}

// TODO: Should `.Next()` take an optional filter function?
func (self *PaymentBatcher) Next() ([]*Payment, error) {
	if self.done {
		return nil, ErrNoResults
	}
	var pymt_list = new(payment_list)
	var qry = self.base_query

	if self.next_id != "" {
		qry = fmt.Sprintf("%s&start_id=%s", qry, self.next_id)
	}

	if err := self.send(&request{
		method:   method.get,
		path:     path.Join(_paymentsPath, qry),
		body:     nil,
		response: pymt_list,
	}); err != nil {
		return nil, err
	}

	if pymt_list.Count == 0 {
		self.done = true
		self.next_id = ""
		return nil, ErrNoResults
	}

	self.next_id = pymt_list.Next_id

	if self.next_id == "" {
		self.done = true
	}

	return pymt_list.Payments, nil
}
