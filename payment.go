package gopal

import "fmt"
import "net/url"
import "net/http"
import "path"
import "strconv"
import "time"

func (self *Payments) Create(method payment_method_i) (*Payment, error) {

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

		self.pending[uuid] = &Payment{
			payment: payment{
				Intent: Sale,
				Redirect_urls: redirects{
					Return_url: return_url,
					Cancel_url: cancel_url,
				},
				Payer: payer{
					Payment_method: method.payment_method(), // PayPal
				},
				Transactions: make([]*transaction, 0),
			},
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

func (self *Payment) AddTransaction(trans Transaction) {
	var t = transaction{
		Amount: amount{
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
	self.payment.Transactions = append(self.payment.Transactions, &t)
}

func (self *Payments) Authorize(req *http.Request) (to string, code int, err error) {
	var pymt *Payment

	pymt, err = self.get(req)
	if err != nil {
		return to, code, err
	}

	defer func() {
		if err != nil {
			self.Cancel(req)
		}
	}()

	err = self.pathGroup.connection.make_request("POST", "payments/payment", &pymt.payment, "send_"+pymt.uuid, &pymt.payment, false)

	if err == nil {
		switch pymt.payment.State {
		case Created:
			// Set url to redirect to PayPal site to begin approval process
			to, _ = pymt.payment.Links.get("approval_url")
			code = 303
		default:
			// otherwise cancel the payment and return an error
			err = UnexpectedResponse
		}
	}

	return to, code, err
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
	// TODO: Should I have some sort of status check? Like for canceling a payment that is complete?
	delete(self.pending, pymt.uuid)
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



// The _times are assigned by PayPal in responses
type _times struct {
	Create_time string `json:"create_time,omitempty"`
	Update_time string `json:"update_time,omitempty"`
}

type Payment struct {
	payment
	payments *Payments
	uuid       string
	url_values url.Values // Holds values from response from PayPal
}
type payment struct {
	_times
	Intent        Intent         `json:"intent,omitempty"`
	Payer         payer          `json:"payer,omitempty"`
	Transactions  []*transaction `json:"transactions,omitempty"`
	Redirect_urls redirects      `json:"redirect_urls,omitempty"`
	Id            string         `json:"id,omitempty"`
	State         State          `json:"state,omitempty"`
	Links         links          `json:"links,omitempty"`

	*payment_error
}

type payment_list struct {
	Payments []payment `json:"payments,omitempty"`
	Count    int       `json:"count,omitempty"`
}

type Transaction struct {
	Total           float64
	Currency        CurrencyType
	Description     string
	Details         *Details
	ShippingAddress *ShippingAddress
	*item_list
}
type transaction struct {
	Amount            amount            `json:"amount"` // Required object
	Description       string            `json:"description,omitempty"`
	Item_list         *item_list        `json:"item_list,omitempty"`
	Related_resources related_resources `json:"related_resources,omitempty`
}

// TODO: I need to define this somehow
// array of sale, authorization, capture, or refund, objects
type related_resources []interface{}

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
type amount struct {
	Currency CurrencyType `json:"currency,omitempty"`     // Required
	Total    float64      `json:"total,omitempty,string"` // Required
	Details  *Details     `json:"details,omitempty"`
}

type Details struct {
	Shipping float64 `json:"shipping,omitempty,string"`
	Subtotal float64 `json:"subtotal,omitempty,string"` // Apparently must be greater than 0
	Tax      float64 `json:"tax,omitempty,string"`
	Fee      float64 `json:"fee,omitempty,string"` // Response only
}

type _trans struct {
	_times
	Id     string `json:"id,omitempty"`
	Amount amount `json:"amount,omitempty"`

	// TODO `state` can hold different values for different types. How to deal?
	State          State  `json:"state,omitempty"`
	Parent_payment string `json:"parent_payment,omitempty"`
}
type refund struct {
	_trans
	Sale_id string `json:"sale_id,omitempty"`
}
type sale struct {
	_trans
	Sale_id     string `json:"sale_id,omitempty"`
	Description string `json:"description,omitempty"`
}
type authorization struct {
	_trans
	Valid_until string `json:"valid_until,omitempty"`
	Links       links  `json:"links,omitempty"`
}
type capture struct {
	_trans
	Is_final_capture bool  `json:"is_final_capture,omitempty"`
	Links            links `json:"links,omitempty"`
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
