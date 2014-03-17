package gopal

import "encoding/json"
import "fmt"
import "net/http"
import "path"
import "time"

type Payments struct {
	connection *Connection
}

// Pagination
//	Assuming `start_time`, `start_index` and `start_id` are mutually exclusive
//	...going to treat them that way anyhow until I understand better.

// I'm going to ignore `start_index` for now since I don't see its usefulness

func (self *Payments) GetAll(size int, sort_by sort_by_i, sort_order sort_order_i, time_range ...time.Time) *PaymentBatcher {
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
		base_query: qry,
		next_id:    "",
		done:       false,
		connection: self.connection,
	}
}

/****************************************

	PaymentBatcher

Manages paginated requests for Payments

*****************************************/

type PaymentBatcher struct {
	base_query string
	next_id    string
	done       bool
	connection *Connection
}

func (self *PaymentBatcher) IsDone() bool {
	return self.done
}

// TODO: Should `.Next()` take an optional filter function?
func (self *PaymentBatcher) Next() ([]*PaymentObject, error) {
	if self.done {
		return nil, ErrNoResults
	}
	var pymt_list = new(payment_list)
	var qry = self.base_query

	if self.next_id != "" {
		qry = fmt.Sprintf("%s&start_id=%s", qry, self.next_id)
	}

	var err = self.connection.make_request("GET",
		"payments/payment"+qry,
		nil, "", pymt_list, false)
	if err != nil {
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

// These provide a way to both get and set the `next_id`.
// This gives the ability to cache the ID, and then set it in a new Batcher.
// Useful if a session is not desired or practical

func (self *PaymentBatcher) GetNextId() string {
	return self.next_id
}
func (self *PaymentBatcher) SetNextId(id string) {
	self.next_id = id
}

func (self *Payments) Get(payment_id string) (*PaymentObject, error) {
	var pymt = new(PaymentObject)
	var err = self.connection.make_request("GET",
		"payments/payment/"+payment_id,
		nil, "", pymt, false)
	if err != nil {
		return nil, err
	}
	return pymt, nil
}

func (self *Payments) Create(method payment_method_i, return_url, cancel_url string) (*PaymentObject, error) {

	if method == PayPal {
		// Make sure we're still authenticated. Will refresh if not.
		var err = self.connection.authenticate()
		if err != nil {
			return nil, err
		}

		return &PaymentObject{
			PaymentExecutor: PaymentExecutor {
				payments:     self,
			},
			Intent: Sale,
			Redirect_urls: redirects{
				Return_url: return_url,
				Cancel_url: cancel_url,
			},
			Payer: payer{
				Payment_method: method.payment_method(), // PayPal
			},
			Transactions: make([]*transaction, 0),
		}, nil
	}
	return nil, nil
}
func (self *Payments) ParseRawData(rawdata []byte) *PaymentObject {
	var po PaymentObject
	var err = json.Unmarshal(rawdata, &po)
	if err != nil {
		return nil
	}
	return &po
}

func (self *Payments) Execute(pymt PaymentFinalizer, req *http.Request) error {
	var query = req.URL.Query()

	// TODO: Is this right? Does URL.Query() ever return nil?
	if query == nil {
		return fmt.Errorf("Attempt to execute a payment that has not been approved")
	}

	var payerid = query.Get("PayerID")
	if payerid == "" {
		return fmt.Errorf("PayerID is missing\n")
	}

	if pymt == nil {
		return fmt.Errorf("Payment Object is missing\n")
	}

	var pymtid = pymt.GetId()
	if pymtid == "" {
		return fmt.Errorf("Payment ID is missing\n")
	}

	var pathname = path.Join("payments/payment", pymtid, "execute")

	var err = self.connection.make_request("POST", pathname, `{"payer_id":"`+payerid+`"}`, "execute_", pymt, false)
	if err != nil {
		return err
	}

	if pymt.GetState() != Approved {
/*
		var s, err = json.Marshal(pymt)
		if err != nil {
			return fmt.Errorf("JSON marshal error\n")
		}
		fmt.Println(string(s))
*/
		return fmt.Errorf("Payment with ID %q for payer %q was not approved\n", pymtid, payerid)
	}

	return nil
}

// TODO: Should this hold the `execute` path so that it doesn't need to be constructed in `Execute()`?
type PaymentExecutor struct {
	Id			 string			`json:"id,omitempty"`
	State		 State			`json:"state,omitempty"`

	RawData		  []byte		`json:"-"`
	*payment_error
	payments *Payments
}
func (self *PaymentExecutor) GetState() State {
	return self.State
}
func (self *PaymentExecutor) GetId() string {
	return self.Id
}
func (self *PaymentExecutor) Execute(r *http.Request) error {
	return self.payments.Execute(self, r)
}

type PaymentFinalizer interface {
	GetState() State
	GetId() string
	Execute(*http.Request) error

	errorable
}

/***************************

	Payment object methods

***************************/

func (self *PaymentObject) GetState() State {
	return self.State
}
func (self *PaymentObject) GetId() string {
	return self.Id
}
func (self *PaymentObject) MakeExecutor() *PaymentExecutor {
	return &PaymentExecutor{
		Id: self.Id,
		State: self.State,
		payments: self.payments,
	}
}
func (self *PaymentObject) AddTransaction(trans Transaction) {
	var t = transaction{
		Amount: Amount{
			Currency: trans.Currency.currency_type(),
			Total:    trans.Total,
		},
		Description: trans.Description,
	}
	if trans.Details != nil {
		t.Amount.Details = &Details{
			Subtotal: trans.Details.Subtotal,
			Tax:      trans.Details.Tax,
			Shipping: trans.Details.Shipping,
		}
		//	Fee: "0",	// This field is only for paypal response data
	}

	if trans.item_list != nil {
		var list = *trans.item_list
		t.Item_list = &list
	}

	if trans.ShippingAddress != nil {
		if t.Item_list == nil {
			t.Item_list = new(item_list)
		}
		t.Item_list.Shipping_address = trans.ShippingAddress
	}
	self.Transactions = append(self.Transactions, &t)
}

// TODO: The tkn parameter is ignored, but should send a query string parameter `token=tkn`
func (self *PaymentObject) Authorize(tkn string) (to string, code int, err error) {

	err = self.payments.connection.make_request("POST", "payments/payment", self, "send_", self, false)

	if err == nil {
		switch self.State {
		case Created:
			// Set url to redirect to PayPal site to begin approval process
			to, _ = self.Links.get("approval_url")
			code = 303
		default:
			// otherwise cancel the payment and return an error
			err = UnexpectedResponse
		}
	}

	return to, code, err
}

func (self *PaymentObject) Execute(req *http.Request) error {
	return self.payments.Execute(self, req)
}

func (t *Transaction) AddItem(qty uint, price float64, curr currency_type_i, name, sku string) {
	if t.item_list == nil {
		t.item_list = new(item_list)
	}
	t.item_list.Items = append(t.item_list.Items, &item{
		Quantity: qty,
		Name:     name,
		Price:    price,
		Currency: curr.currency_type(),
		Sku:      sku,
	})
}

// TODO: I'm returning a list of Sale objects because every payment can have multiple transactions.
//		I need to find out why a payment can have multiple transactions, and see if I should eliminate that in the API
//		Also, I need to find out why `related_resources` is an array. Can there be more than one per type?
func (self *PaymentObject) GetSale() []*SaleObject {
	var sales = []*SaleObject{}
	for _, transaction := range self.Transactions {
		for _, related_resource := range transaction.Related_resources {
			if related_resource.Sale != nil {
				sales = append(sales, related_resource.Sale)
			}
		}
	}
	return sales
}

// The _times are assigned by PayPal in responses
type _times struct {
	Create_time string `json:"create_time,omitempty"`
	Update_time string `json:"update_time,omitempty"`
}

type payment_list struct {
	Payments []*PaymentObject `json:"payments"`
	Count    int              `json:"count"`
	Next_id  string           `json:"next_id"`
	*identity_error
}

type PaymentObject struct {
	_times
	PaymentExecutor
	Intent        Intent         `json:"intent,omitempty"`
	Payer         payer          `json:"payer,omitempty"`
	Transactions  []*transaction `json:"transactions,omitempty"`
	Redirect_urls redirects      `json:"redirect_urls,omitempty"`
	Links         links          `json:"links,omitempty"`
}

// This is for simplified Transaction creation
// TODO: Should I keep this abstraction, or just require the PayPal model directly?
type Transaction struct {
	Total           float64				// maps to Amount.Total
	Currency        CurrencyType		// maps to Amount.Currency
	Description     string
	Details         *Details			// maps to Amount.Details
	ShippingAddress *ShippingAddress	// maps to item_list.Shipping_address
	*item_list
}
type transaction struct {
	Amount            Amount            `json:"amount"` // Required object
	Description       string            `json:"description,omitempty"`
	Item_list         *item_list        `json:"item_list,omitempty"`
	Related_resources related_resources `json:"related_resources,omitempty"`
}

// array of SaleObject, AuthorizationObject, CaptureObject, or RefundObject
type related_resources []related_resource

type related_resource struct {
	Sale          *SaleObject          `json:"sale,omitempty"`
	Authorization *AuthorizationObject `json:"authorization,omitempty"`
	Capture       *CaptureObject       `json:"capture,omitempty"`
	Refund        *RefundObject        `json:"refund,omitempty"`
}

type payment_execution struct {
	Payer_id     string         `json:"payer_id,omitempty"`
	Transactions []*Transaction `json:"transactions,omitempty"`
}

type item_list struct {
	Items            []*item          `json:"items,omitempty"`
	Shipping_address *ShippingAddress `json:"shipping_address,omitempty"`
}
type item struct {
	Quantity uint         `json:"quantity,omitempty,string"`
	Name     string       `json:"name,omitempty"`
	Price    float64      `json:"price,omitempty,string"`
	Currency CurrencyType `json:"currency,omitempty"`
	Sku      string       `json:"sku,omitempty"`
}

/*
	Currency and Total fields are required when making payments
*/
type Amount struct {
	Currency CurrencyType `json:"currency"`
	Total    float64      `json:"total,string"`
	Details  *Details     `json:"details,omitempty"`
}

type Details struct {
	Shipping float64 `json:"shipping,omitempty,string"`
	Subtotal float64 `json:"subtotal,omitempty,string"` // Apparently must be greater than 0
	Tax      float64 `json:"tax,omitempty,string"`
	Fee      float64 `json:"fee,omitempty,string"` // Response only
}

/*
	Amount is always required. TODO: verify this
*/
type _trans struct {
	_times
	Id     string `json:"id,omitempty"`
	Amount Amount `json:"amount"`

	Parent_payment string `json:"parent_payment,omitempty"`
}

type links []link

func (l links) get(s string) (string, string) {
	for i, _ := range l {
		if l[i].Rel == s {
			return l[i].Href, l[i].Method
		}
	}
	return "", ""
}

type link struct {
	Href   string `json:"href,omitempty"`
	Rel    string `json:"rel,omitempty"`
	Method string `json:"method,omitempty"`
}

type redirects struct {
	Return_url string `json:"return_url,omitempty"`
	Cancel_url string `json:"cancel_url,omitempty"`
}
