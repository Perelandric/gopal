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

func NewConnection(live connection_type, id, secret, host string) (*Connection, error) {
	var hosturl, err = url.Parse(host)
	if err != nil {
		return nil, err
	}
	var conn = &Connection{live:live, id:id, secret:secret, hosturl:hosturl, client:http.Client{}}
	err = conn.authenticate()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

type Connection struct {
	live connection_type
	id, secret string
	hosturl *url.URL
	client http.Client
	tokeninfo tokeninfo
}

func (self *Connection) PathGroup(valid, cancel string) (*PathGroup, error) {
	for _, p := range [...]*string{&valid, &cancel} {
		var u, err = url.Parse(*p)
		if err != nil {
			return nil, err
		}
		*p = self.hosturl.ResolveReference(u).String()
	}

	return &PathGroup{self, valid, cancel, make(map[string]*Payment)}, nil
}

func (self *Connection) authenticate() error {
	// If an error is returned, zero the tokeninfo
	var err error
	var duration time.Duration

	// No need to authenticate if the previous has not yet expired
	if time.Now().Before(self.tokeninfo.expiration) {
		return nil
	}

	defer func() {
		if err != nil {
			self.tokeninfo = tokeninfo{}
		}
	}()

	// (re)authenticate
	err = self.make_request("POST", "/oauth2/token", "grant_type=client_credentials", "", &self.tokeninfo, true)
	if err != nil {
		return err
	}

	// Set the duration to expire 3 minutes early to avoid expiration during a request cycle
	duration = time.Duration(self.tokeninfo.Expires_in) * time.Second - 3 * time.Minute
	self.tokeninfo.expiration = time.Now().Add(duration)

	return nil
}




type PathGroup struct {
	connection *Connection
	return_url string
	cancel_url string
	pending map[string]*Payment
}

func (self *PathGroup) GetPayment(req *http.Request) (*Payment, error) {
	var query = req.URL.Query()
	var uuid = query.Get("uuid")
	var pymt, _ = self.pending[uuid]

	if pymt == nil || pymt.uuid != uuid {
		return nil, fmt.Errorf("Unknown payment")
	}
	pymt.url_values = query
	return pymt, nil
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
func (self PaymentMethod) payment_method() PaymentMethod {
	return self
}
type payment_method_i interface {
	payment_method() PaymentMethod
}
const CreditCard = PaymentMethod("credit_card")
const PayPal = PaymentMethod("paypal")




type CreditCardType string
func (self CreditCardType) credit_card_type() CreditCardType {
	return self
}
type credit_card_type_i interface {
	credit_card_type() CreditCardType
}
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




type CurrencyType string
func (self CurrencyType) currency_type() CurrencyType {
	return self
}
type currency_type_i interface {
	currency_type() CurrencyType
}
const AUD = CurrencyType("AUD") // Australian dollar
const BRL = CurrencyType("BRL") // Brazilian real**
const CAD = CurrencyType("CAD") // Canadian dollar
const CZK = CurrencyType("CZK") // Czech koruna
const DKK = CurrencyType("DKK") // Danish krone
const EUR = CurrencyType("EUR") // Euro
const HKD = CurrencyType("HKD") // Hong Kong dollar
const HUF = CurrencyType("HUF") // Hungarian forint
const ILS = CurrencyType("ILS") // Israeli new shekel
const JPY = CurrencyType("JPY") // Japanese yen*
const MYR = CurrencyType("MYR") // Malaysian ringgit**
const MXN = CurrencyType("MXN") // Mexican peso
const TWD = CurrencyType("TWD") // New Taiwan dollar*
const NZD = CurrencyType("NZD") // New Zealand dollar
const NOK = CurrencyType("NOK") // Norwegian krone
const PHP = CurrencyType("PHP") // Philippine peso
const PLN = CurrencyType("PLN") // Polish złoty
const GBP = CurrencyType("GBP") // Pound sterling
const SGD = CurrencyType("SGD") // Singapore dollar
const SEK = CurrencyType("SEK") // Swedish krona
const CHF = CurrencyType("CHF") // Swiss franc
const THB = CurrencyType("THB") // Thai baht
const TRY = CurrencyType("TRY") // Turkish lira**
const USD = CurrencyType("USD") // United States dollar



