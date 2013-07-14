package gopal

import "encoding/json"
import "bytes"
import "fmt"

func (ppp *PayPalPath) PayPalPayment() (*Payment, error) {
	var err = ppp.paypal.authenticate()
	if err != nil {
		return nil, err
	}

	return &Payment {
		payment: payment {
			Intent: Sale,
			Redirect_urls: redirects {
				Return_url: ppp.return_url,
				Cancel_url: ppp.cancel_url,
			},
			Payer: payer {
				Payment_method: PayPalMethod,
			},
			Transactions: make([]*Transaction, 0),
//			Payment_error: nil,
		},
		pp_auth: ppp.paypal,
	}, nil
}

func (ppp *PayPalPath) CreditCardPayment() error {
	return nil
}


type Payment struct {
	payment
	pp_auth *PayPal
}

func (pay *Payment) AddTransaction(amt float64, curr, desc string) (*Transaction, error) {
	var t = &Transaction{
		transaction{
			Amount: amount {
				Currency: curr,
				Total: amt,
				Details: details {
					Subtotal: amt,
				},
			},
			Description: desc,
		},
	}
	pay.payment.Transactions = append(pay.payment.Transactions, t)
	return t, nil
}

func (pay *Payment) Execute() (string, int, error) {
	var err error
	var to = ""
	var code = 500
	var body []byte
	var resp []byte

	body, err = json.Marshal(&pay.payment)
	if err != nil {
		return to, code, err
	}

fmt.Println(string(body))

	resp, err = pay.pp_auth.make_request("POST", "payments/payment", bytes.NewReader(body), false)
	if err != nil {
		return to, code, err
	}

fmt.Printf("\n%s\n", resp)

	err = json.Unmarshal(resp, &pay.payment)
	if err != nil {
		return to, code, err
	}

	if pay.payment.Payment_error != nil {
		return to, code, pay.payment.Payment_error.to_error()
	}

	to = pay.payment.Redirect_urls.Cancel_url

	if pay.payment.State == Created {
		for _, link := range pay.payment.Links {
			if link.Rel == "approval_url" {
				to = link.Href
				break
			}
		}
		code = 303
	}

	return to, code, err
}


type Transaction struct {
	transaction
}
func (t *Transaction) SetDetails(shipping, subtotal, tax, fee float64) {
	t.Amount.Details.Shipping = shipping
	t.Amount.Details.Subtotal = subtotal
	t.Amount.Details.Tax = tax
	t.Amount.Details.Fee = fee
}
func (t *Transaction) SetShippingAddress(recip_name string, typ AddressType, addrss Address) {
	if t.Item_list == nil {
		t.Item_list = new(item_list)
	}
	t.Item_list.Shipping_address = &shipping_address{
		Recipient_name: recip_name,
		Type: typ,
		Address: addrss,
	}
}
func (t *Transaction) AddItem(qty uint, price float64, curr, name, sku string) error {
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
	return nil
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
	Links []links					`json:"links,omitempty"`

	// 
	Payment_error *payment_error	`json:"payment_error,omitempty"`
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
	Currency string	`json:"currency,omitempty"`
	Total float64	`json:"total,omitempty,string"`
	Details details	`json:"details,omitempty"`
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
	Links []links		`json:"links,omitempty"`
}
type capture struct {
	_trans
	Is_final_capture bool	`json:"is_final_capture,omitempty"`
	Links []links			`json:"links,omitempty"`
}

type links struct {
	Href string		`json:"href,omitempty"`
	Rel string		`json:"rel,omitempty"`
	Method string	`json:"method,omitempty"`
}

type redirects struct {
    Return_url string	`json:"return_url,omitempty"`
    Cancel_url string	`json:"cancel_url,omitempty"`
}

