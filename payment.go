package gopal

import "fmt"
import "net/url"
import "net/http"
import "path"
import "strconv"
import "time"

type Payments struct {
    pathGroup *PathGroup
    pending map[string]*PaymentObject
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

	return &PaymentBatcher {
		base_query: qry,
		next_id: "",
		done: false,
		pathGroup: self.pathGroup,
	}
}


/****************************************

	PaymentBatcher

Manages paginated requests for Payments

*****************************************/

type PaymentBatcher struct {
	base_query string
	next_id string
	done bool
	pathGroup *PathGroup
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

	var err = self.pathGroup.connection.make_request("GET",
													"payments/payment" + qry,
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



func (self *Payments) get(req *http.Request) (*PaymentObject, error) {
    var query = req.URL.Query()
    var uuid = query.Get("uuid")
    var pymt, _ = self.pending[uuid]

    if pymt == nil || pymt.uuid != uuid {
        return nil, fmt.Errorf("Unknown payment")
    }
    pymt.url_values = query
    return pymt, nil
}


func (self *Payments) Get(payment_id string) (*PaymentObject, error) {
    var pymt = new(PaymentObject)
    var err = self.pathGroup.connection.make_request("GET",
													"payments/payment/" + payment_id,
													nil, "", pymt, false)
    if err != nil {
        return nil, err
    }
    return pymt, nil
}

func (self *Payments) Create(method payment_method_i) (*PaymentObject, error) {

	if method == PayPal {
		// Make sure we're still authenticated. Will refresh if not.
		var err = self.pathGroup.connection.authenticate()
		if err != nil {
			return nil, err
		}

		var uuid string

		for { // TODO Clearly this needs to be improved
			uuid = strconv.FormatInt(time.Now().UnixNano(), 36)
			var _, ok = self.pending[uuid]
			if ok {
				continue
			}
			break
		}

		var return_url, cancel_url = self.pathGroup.return_url, self.pathGroup.cancel_url

		for _, uptr := range [...]*string{&return_url, &cancel_url} {
			u, _ := url.Parse(*uptr)
			var q = u.Query()
			q.Set("uuid", uuid)
			u.RawQuery = q.Encode()
			*uptr = u.String()
		}

		self.pending[uuid] = &PaymentObject{
			Intent: Sale,
			Redirect_urls: redirects{
				Return_url: return_url,
				Cancel_url: cancel_url,
			},
			Payer: payer{
				Payment_method: method.payment_method(), // PayPal
			},
			Transactions: make([]*transaction, 0),
			url_values: nil,
			payments: self,
			uuid:       uuid,
		}

		return self.pending[uuid], nil
	}
	return nil, nil
}

/***************************

	Payment object methods

***************************/

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

func (self *PaymentObject) Authorize() (to string, code int, err error) {
	defer func() {
		if err != nil {
			self.Cancel()
		}
	}()

	err = self.payments.pathGroup.connection.make_request("POST", "payments/payment", self, "send_"+self.uuid, self, false)

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
func (self *PaymentObject) Cancel() {
	// TODO: Should I have some sort of status check? Like for canceling a payment that is complete?
	delete(self.payments.pending, self.uuid)
}

func (self *Payments) Execute(req *http.Request) error {
	var payerid string
	var pathname string

	var pymt, err = self.get(req)

	if err != nil {
		return err
	}

	var query = pymt.url_values

	if query == nil {
		return fmt.Errorf("Attempt to execute a payment that has not been approved")
	}

	payerid = query.Get("PayerID")
	if payerid == "" {
		return fmt.Errorf("PayerID is missing")
	}

	pathname = path.Join("payments/payment", pymt.Id, "execute")

	// TODO Maybe make_request() should check the resulting unmarshaled body for a PayPal error object?
	// 		Every object I pass to take the body would need to implement an interface to allow this.

	err = self.pathGroup.connection.make_request("POST", pathname, `{"payer_id":"`+payerid+`"}`, "execute_"+pymt.uuid, pymt, false)
	if err != nil {
		return err
	}

	if pymt.State != Approved {
		return fmt.Errorf("Payment not approved")
	}
	// TODO I should remove the Payment object from Pending at this point?
	return nil
}

func (self *Payments) Cancel(req *http.Request) error {
	var pymt, err = self.get(req)
	if err != nil {
		return err
	}
	pymt.Cancel()
	return nil
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
func (self *PaymentObject) GetSale() ([]*SaleObject) {
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
	Payments []*PaymentObject	`json:"payments"`
	Count int					`json:"count"`
	Next_id string				`json:"next_id"`
	*identity_error
}

type PaymentObject struct {
	_times
	Intent        Intent         `json:"intent,omitempty"`
	Payer         payer          `json:"payer,omitempty"`
	Transactions  []*transaction `json:"transactions,omitempty"`
	Redirect_urls redirects      `json:"redirect_urls,omitempty"`
	Id            string         `json:"id,omitempty"`
	State         State          `json:"state,omitempty"`
	Links         links          `json:"links,omitempty"`

	*payment_error
	payments *Payments
	uuid       string
	url_values url.Values // Holds values from response from PayPal
}

// This is for simplified Transaction creation
// TODO: Should I keep this abstraction, or just require the PayPal model directly?
type Transaction struct {
	Total           float64
	Currency        CurrencyType
	Description     string
	Details         *Details
	ShippingAddress *ShippingAddress
	*item_list
}
type transaction struct {
	Amount            Amount            `json:"amount"` // Required object
	Description       string            `json:"description,omitempty"`
	Item_list         *item_list        `json:"item_list,omitempty"`
	Related_resources related_resources `json:"related_resources,omitempty`
}

// array of SaleObject, AuthorizationObject, CaptureObject, or RefundObject
type related_resources []related_resource

type related_resource struct {
	Sale			*SaleObject				`json:"sale",omitempty`
	Authorization	*AuthorizationObject	`json:"authorization",omitempty`
	Capture			*CaptureObject			`json:"capture",omitempty`
	Refund			*RefundObject			`json:"refund",omitempty`
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
