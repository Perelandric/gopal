package gopal

import "net/http"
import "net/url"
import "time"
import "strings"
import "encoding/json"

func New(sandbox bool, id, secret, host string) (*PayPal, error) {
	var hosturl, err = url.Parse(host)
	if err != nil {
		return nil, err
	}
	var pp = &PayPal{sandbox:sandbox, id:id, secret:secret, hosturl:hosturl, client:http.Client{}}
	err = pp.authenticate()
	if err != nil {
		return nil, err
	}
	return pp, nil
}

type PayPal struct {
	sandbox bool
	id, secret string
	hosturl *url.URL
	client http.Client
	tokeninfo tokeninfo
}

func (pp *PayPal) ForPath(pathname string) (*PayPalPath, error) {
	var pathurl, err = url.Parse(pathname)
	if err != nil {
		return nil, err
	}

	pathurl = pp.hosturl.ResolveReference(pathurl)

	var ppp = &PayPalPath{pp, pathurl.Path, "", "", ""}
	var q = pathurl.Query()

	q.Set("status", "request")
	pathurl.RawQuery = q.Encode()
	ppp.request_url = pathurl.String()

	q.Set("status", "valid")
	pathurl.RawQuery = q.Encode()
	ppp.return_url = pathurl.String()

	q.Set("status", "cancel")
	pathurl.RawQuery = q.Encode()
	ppp.cancel_url = pathurl.String()

	return ppp, nil
}

func (pp *PayPal) is_authenticated() bool {
	return time.Now().Before(pp.tokeninfo.expiration)
}

func (pp *PayPal) authenticate() error {
	// If an error is returned, zero the tokeninfo
	var err error
	var data []byte
	var duration time.Duration

	// No need to authenticate if the previous has not yet expired
	if pp.is_authenticated() {
		return nil
	}

	defer func() {
		if err != nil {
			pp.tokeninfo = tokeninfo{}
		}
	}()

	// (re)authenticate
	data, err = pp.make_request("POST", "/oauth2/token", strings.NewReader("grant_type=client_credentials"), true)

	if err != nil {
		return err
	}

	// Parse the JSON response
	err = json.Unmarshal(data, &pp.tokeninfo)
	if err != nil {
		return err
	}

	if pp.tokeninfo.identity_error != nil {
		err = pp.tokeninfo.identity_error.to_error()
		return err
	}

	// Set the duration to expire 3 minutes early to avoid expiration during a request cycle
	duration = time.Duration(pp.tokeninfo.Expires_in) * time.Second - 3 * time.Minute
	pp.tokeninfo.expiration = time.Now().Add(duration)

	return nil
}


type PayPalPath struct {
	paypal *PayPal
	pathname string
	request_url string
	return_url string
	cancel_url string
}

func (ppp *PayPalPath) Path() string {
	return ppp.pathname
}
func (ppp *PayPalPath) FullPath() string {
	return ppp.request_url
}

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



