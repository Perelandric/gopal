package gopal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"path"
	"time"
)

func (self *connection) NewPaypalPayment(
	payer payer, urls Redirects) (*Payment, error) {

	// Make sure we're still authenticated. Will refresh if not.

	// TODO: Should I really authenticate here? Seems like I should wait for the
	// actual request to be made.
	if err := self.authenticate(); err != nil {
		return nil, err
	}

	var pymt = Payment{
		connection: self,
	}
	pymt.private.Payer = payer
	pymt.private.Intent = intent.sale
	pymt.private.Transactions = make([]*Transaction, 0)
	pymt.private.RedirectUrls = urls

	return &pymt, nil
}

/***************************

	Payment object methods

***************************/

func (self *Payment) AddTransaction(
	c CurrencyTypeEnum, shp *ShippingAddress) *Transaction {

	var t Transaction
	t.private.Amount = amount{}
	t.private.ItemList = &itemList{}

	t.private.Amount.private.Currency = c
	t.private.Amount.private.Total = 0

	t.private.ItemList.private.Items = make([]*Item, 0, 1)
	t.private.ItemList.private.ShippingAddress = shp

	self.private.Transactions = append(self.private.Transactions, &t)

	return &t
}

// Prices are assumed to use the CurrencyType passed to NewTransaction.
func (t *Transaction) AddItem(item *Item) (err error) {
	if item == nil {
		log.Println("Transaction received unexpected nil Item")
		return nil
	}
	if err = item.validate(); err != nil {
		return err
	}

	if item.Price == 0 {
		// TODO: should have some sort of non-zero flag to denote a zero amount
	}

	item.private.Currency = t.private.Amount.private.Currency

	t.private.ItemList.private.Items =
		append(t.private.ItemList.private.Items, item)
	return nil
}

// TODO: Should send a query string parameter `token=[some token]`
func (self *Payment) Authorize() (to string, code int, err error) {
	if err = self.validate(); err != nil {
		return "", 0, err
	}

	self.calculateToAuthorize()

	// Create Totals
	var pymt Payment

	err = self.send(&request{
		method:   method.post,
		path:     _paymentsPath,
		body:     self,
		response: &pymt,
	})

	if err == nil {
		switch pymt.private.State {
		case state.Created:
			// Set url to redirect to PayPal site to begin approval process
			to, _ = pymt.private.Links.get(relType.approvalUrl)
			code = 303
		default:
			// otherwise cancel the payment and return an error
			err = UnexpectedResponse
		}
	}

	return to, code, err
}

func (self *connection) Execute(u url.URL) error {
	var query = u.Query()

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

	if pymt.private.State != state.Approved {
		return fmt.Errorf(
			"Payment with ID %q for payer %q was not approved\n", pymtid, payerid)
	}

	return nil
}

// TODO: I'm returning a list of Sale objects because every payment can have multiple transactions.
//		I need to find out why a payment can have multiple transactions, and see if I should eliminate that in the API
func (self *Payment) FetchSale() []*Sale {
	var sales = []*Sale{}
	for _, trans := range self.private.Transactions {
		for _, related_resource := range trans.private.RelatedResources {
			if s, ok := related_resource.(*Sale); ok {
				sales = append(sales, s)
			}
		}
	}
	return sales
}

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

// Used to execute the approved payment.
type paymentExecution struct {
	// The ID of the Payer, passed in the return_url by PayPal.
	PayerID string `json:"payer_id,omitempty"`

	// Transactional details if updating a payment. Note that this instance of
	// the transactions object accepts only the Amount object.
	Transactions []*Transaction `json:"transactions,omitempty"`
}

/**********************

	Transaction objects

**********************/

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

//func (self *relatedResources) MarshalJSON() ([]byte, error) {
//	return []byte("[]"), nil
//}

func (self *relatedResources) UnmarshalJSON(b []byte) error {
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
			var t resource // for unmarshaling the current item

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
				log.Printf("Unexpected resource type: %s\n", name)
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
	baseQuery string
	nextId    string
	done      bool
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
	var qry = self.baseQuery

	if self.nextId != "" {
		qry = fmt.Sprintf("%s&start_id=%s", qry, self.nextId)
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
		self.nextId = ""
		return nil, ErrNoResults
	}

	self.nextId = pymt_list.NextId

	if self.nextId == "" {
		self.done = true
	}

	return pymt_list.Payments, nil
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
		baseQuery:  qry,
		nextId:     "",
		done:       false,
	}
}

// These provide a way to both get and set the `next_id`.
// This gives the ability to cache the ID, and then set it in a new Batcher.
// Useful if a session is not desired or practical

func (self *PaymentBatcher) GetNextId() string {
	return self.nextId
}
func (self *PaymentBatcher) SetNextId(id string) {
	self.nextId = id
}

type payment_list struct {
	Payments []*Payment `json:"payments"`
	Count    int        `json:"count"`
	NextId   string     `json:"next_id"`
	*identity_error
}
