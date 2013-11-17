package gopal

import "fmt"
import "net/url"
import "path"
import "strconv"
import "time"

func (ppp *PathGroup) PayPalPayment() (*Payment, error) {

	// Make sure we're still authenticated. Will refresh if not.
	var err = ppp.paypal.authenticate()
	if err != nil {
		return nil, err
	}

	var uuid string

    for {
        // TODO Clearly this needs to be improved
        uuid = strconv.FormatInt(time.Now().UnixNano(), 36)
        var _, ok = ppp.pending[uuid]
		if ok {
            continue
        }
        break
    }

	u, err := url.Parse(ppp.return_url)
	var q = u.Query()
	q.Set("uuid", uuid)
	u.RawQuery = q.Encode()
	var return_url = u.String()

	u, err = url.Parse(ppp.cancel_url)
	q.Set("uuid", uuid)
	u.RawQuery = q.Encode()
	var cancel_url = u.String()

	ppp.pending[uuid] = &Payment {
		payment: payment {
			Intent: Sale,
			Redirect_urls: redirects {
				Return_url: return_url,
				Cancel_url: cancel_url,
			},
			Payer: payer {
				Payment_method: PayPalMethod,
			},
			Transactions: make([]*Transaction, 0),
		},
		path: ppp,
		uuid: uuid,
	}

	return ppp.pending[uuid], nil
}

func (ppp *PathGroup) CreditCardPayment() error {
	return nil
}


type Payment struct {
	payment
	path *PathGroup
	uuid string
}

func (pymt *Payment) AddTransaction(amt float64, curr, desc string) (*Transaction, error) {
	var t = &Transaction{
		transaction{
			Amount: amount {
				Currency: curr,
				Total: amt,
			},
			Description: desc,
		},
	}
	pymt.payment.Transactions = append(pymt.payment.Transactions, t)
	return t, nil
}

func (pymt *Payment) send() (string, int, error) {
	var err error
	var to = ""
	var code = 500

	err = pymt.path.paypal.make_request("POST", "payments/payment", &pymt.payment, "send_" + pymt.uuid, &pymt.payment, false)
	if err != nil {
		return to, code, err
	}

	to = pymt.payment.Redirect_urls.Cancel_url

	if pymt.payment.State == Created {
		to, _ = pymt.payment.Links.get("approval_url")
		code = 303
	}

	return to, code, err
}

func (pymt *Payment) execute(query url.Values) error {
	var err error
	var payerid string
	var pathname string

	if pymt == nil {
		return fmt.Errorf("No payment found")
	}

	payerid = query.Get("PayerID")
	if payerid == "" {
		return fmt.Errorf("PayerID is missing")
	}

	pathname = path.Join("payments/payment", pymt.Id, "execute")

// TODO Maybe make_request() should check the resulting unmarshaled body for a PayPal error object?
// 		Every object I pass to take the body would need to implement an interface to allow this.

	err = pymt.path.paypal.make_request("POST", pathname, `{"payer_id":"`+payerid+`"}`, "execute_" + pymt.uuid, pymt, false)
	if err != nil {
		return err
	}

	if pymt.State != Approved {
		return fmt.Errorf("Payment not approved")
	}
// TODO I should remove the Payment object from Pending at this point?
	return nil
}




type Transaction struct {
	transaction
}
func (t *Transaction) SetDetails(shipping, subtotal, tax, fee float64) {
	t.Amount.Details = &details {
		Shipping: shipping,
		Subtotal: subtotal,
		Tax: tax,
		Fee: fee,
	}
}
func (t *Transaction) SetShippingAddress(recip_name string, typ AddressType, addrss Address) {
	if t.Item_list == nil {
		t.Item_list = new(item_list)
	}
	t.Item_list.Shipping_address = &shipping_address {
		Recipient_name: recip_name,
		Type: typ,
		Address: addrss,
	}
}
func (t *Transaction) AddItem(qty uint, price float64, curr, name, sku string) {
	if t.Item_list == nil {
		t.Item_list = new(item_list)
	}
	t.Item_list.Items = append(t.Item_list.Items, &item {
		Quantity: qty,
		Name: name,
		Price: price,
		Currency: curr,
		Sku: sku,
	})
}



// The _times are assigned by PayPal in responses
type _times struct {
    Create_time string  `json:"create_time,omitempty"`
    Update_time string  `json:"update_time,omitempty"`
}

type payment struct {
	_times
	Intent Intent					`json:"intent,omitempty"`
	Payer payer						`json:"payer,omitempty"`
	Transactions []*Transaction		`json:"transactions,omitempty"`
	Redirect_urls redirects			`json:"redirect_urls,omitempty"`
	Id string						`json:"id,omitempty"`
	State State						`json:"state,omitempty"`
	Links links					`json:"links,omitempty"`

	// 
	*payment_error	`json:"payment_error,omitempty"`
}

type payment_list struct {
	Payments []payment	`json:"payments,omitempty"`
	Count int			`json:"count,omitempty"`
}

type transaction struct {
	Amount amount			`json:"amount,omitempty"`
	Description string		`json:"description,omitempty"`
	Item_list *item_list	`json:"item_list,omitempty"`
}

type payment_execution struct {
	Payer_id string				`json:"payer_id,omitempty"`
	Transactions []*Transaction	`json:"transactions,omitempty"`
}

type item_list struct {
	Items []*item						`json:"items,omitempty"`
	Shipping_address *shipping_address	`json:"shipping_address,omitempty"`
}
type item struct {
	Quantity uint	`json:"quantity,omitempty,string"`
	Name string		`json:"name,omitempty"`
	Price float64	`json:"price,omitempty,string"`
	Currency string	`json:"currency,omitempty"`
	Sku string		`json:"sku,omitempty"`
}
type amount struct {
	Currency string		`json:"currency,omitempty"`
	Total float64		`json:"total,omitempty,string"`
	Details *details	`json:"details,omitempty"`
}
type details struct {
	Shipping float64	`json:"shipping,omitempty"`
	Subtotal float64	`json:"subtotal,omitempty,string"`
	Tax float64			`json:"tax,omitempty,string"`
	Fee float64			`json:"fee,omitempty,string"`
}

type _trans struct {
    _times
    Id string               `json:"id,omitempty"`
    Amount amount           `json:"amount,omitempty"`

    // TODO `state` can hold different values for different types. How to deal?
    State State				`json:"state,omitempty"`
    Parent_payment string   `json:"parent_payment,omitempty"`
}
type refund struct {
	_trans
	Sale_id string	`json:"sale_id,omitempty"`
}
type sale struct {
	_trans
	Sale_id string		`json:"sale_id,omitempty"`
	Description string	`json:"description,omitempty"`
}
type authorization struct {
	_trans
	Valid_until string	`json:"valid_until,omitempty"`
	Links links			`json:"links,omitempty"`
}
type capture struct {
	_trans
	Is_final_capture bool	`json:"is_final_capture,omitempty"`
	Links links				`json:"links,omitempty"`
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
	Href string		`json:"href,omitempty"`
	Rel string		`json:"rel,omitempty"`
	Method string	`json:"method,omitempty"`
}

type redirects struct {
    Return_url string	`json:"return_url,omitempty"`
    Cancel_url string	`json:"cancel_url,omitempty"`
}

