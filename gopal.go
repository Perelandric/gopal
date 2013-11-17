package gopal

import "fmt"
import "net/http"
import "net/url"
import "time"

const (
	Sandbox = c_type(false)
	Live = c_type(true)
)

type c_type bool
func (self c_type) getType() bool {
	return bool(self)
}
type connection_type interface {
	getType() bool
}

func NewConnection(live connection_type, id, secret, host string) (*PayPal, error) {
	var hosturl, err = url.Parse(host)
	if err != nil {
		return nil, err
	}
	var pp = &PayPal{live:live, id:id, secret:secret, hosturl:hosturl, client:http.Client{}}
	err = pp.authenticate()
	if err != nil {
		return nil, err
	}
	return pp, nil
}

type PayPal struct {
	live connection_type
	id, secret string
	hosturl *url.URL
	client http.Client
	tokeninfo tokeninfo
}

func (pp *PayPal) PathGroup(request, valid, cancel string) (*PathGroup, error) {
	for _, p := range [3]*string{&request, &valid, &cancel} {
		var u, err = url.Parse(*p)
		if err != nil {
			return nil, err
		}
		*p = pp.hosturl.ResolveReference(u).String()
	}

	return &PathGroup{pp, request, valid, cancel, make(map[string]*Payment)}, nil
}

func (pp *PayPal) authenticate() error {
	// If an error is returned, zero the tokeninfo
	var err error
	var duration time.Duration

	// No need to authenticate if the previous has not yet expired
	if time.Now().Before(pp.tokeninfo.expiration) {
		return nil
	}

	defer func() {
		if err != nil {
			pp.tokeninfo = tokeninfo{}
		}
	}()

	// (re)authenticate
	err = pp.make_request("POST", "/oauth2/token", "grant_type=client_credentials", "", &pp.tokeninfo, true)
	if err != nil {
		return err
	}

	// Set the duration to expire 3 minutes early to avoid expiration during a request cycle
	duration = time.Duration(pp.tokeninfo.Expires_in) * time.Second - 3 * time.Minute
	pp.tokeninfo.expiration = time.Now().Add(duration)

	return nil
}




type PathGroup struct {
	paypal *PayPal
	request_url string
	return_url string
	cancel_url string
	pending map[string]*Payment
}

func (ppp *PathGroup) Send(pymt *Payment) (string, int, error) {
	if pymt == nil || ppp.pending[pymt.uuid] != pymt {
		return "", 0, fmt.Errorf("Unknown payment.")
	}
	return pymt.send()
}

func (ppp *PathGroup) Execute(req *http.Request) error {
	var query = req.URL.Query()
	return ppp.pending[query.Get("uuid")].execute(query)
}

func (ppp *PathGroup) Cancel(req *http.Request) {
	delete(ppp.pending, req.URL.Query().Get("uuid"))
}

// TODO: Do I need this method?
func (ppp *PathGroup) FullPath() string {
	return ppp.request_url
}



// Authorization response
type tokeninfo struct {
	Scope string			`json:"scope,omitempty"`
	Access_token string		`json:"access_token,omitempty"`
	Refresh_token string	`json:"refresh_token,omitempty"`
	Token_type string		`json:"token_type,omitempty"`
	Expires_in uint			`json:"expires_in,omitempty"`

	// Not sure about this field. Appears in response, but not in documentation.
	App_id string			`json:"app_id,omitempty"`

	// Derived fields
	expiration time.Time

	// Handles the case where an error is received instead
	*identity_error
}
func (ti *tokeninfo) auth_token() string {
	return ti.Token_type + " " + ti.Access_token
}


type Intent string
const Sale = Intent("sale")
const Authorize = Intent("authorize")


type AddressType string
const Residential = AddressType("residential")
const Business = AddressType("business")
const Mailbox = AddressType("mailbox")


type PaymentMethod string
const CreditCardMethod = PaymentMethod("credit_card")
const PayPalMethod = PaymentMethod("paypal")


type CreditCardType string
const Visa = CreditCardType("visa")
const MasterCard = CreditCardType("mastercard")
const Discover = CreditCardType("discover")
const Amex = CreditCardType("amex")


type State string

// payment
const Created = State("created")
const Approved = State("approved")
const Canceled = State("canceled")

// payment/refund
const Failed = State("failed")

// refund/sale/capture/authorization
const Pending = State("pending")

// refund/sale/capture
const Completed = State("completed")

// sale/capture
const Refunded = State("refunded")
const PartiallyRefunded = State("partially_refunded")

// payment/credit_card/authorization
const Expired = State("expired")

// credit_card
const Ok = State("ok")

// authorization
const Authorized = State("authorized")
const Captured = State("captured")
const PartiallyCaptured = State("partially_captured")
const Voided = State("voided")



