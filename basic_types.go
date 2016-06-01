package gopal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
	"sync"
	"time"
)

//go:generate Golific $GOFILE

type dateTime string // TODO: How should this be done? [Un]Marshalers?

// TODO: only needed until Golific acknowledges more types
type fundingInstruments []*fundingInstrument

type resource interface {
	errorable
	getPath() string
	Amount() amount
}

type Payment interface {
	calculateToAuthorize()
}

type Transaction interface {
	validate() error
	calculateToAuthorize()
}

type refundable interface {
	resource
	getRefundPath() string
}

type request struct {
	method       methodEnum
	path         string
	body         interface{}
	response     errorable
	responseData []byte
	isAuthReq    bool
}

type connection struct {
	server     ServerEnum
	id, secret string
	tokeninfo  tokeninfo
	client     http.Client
}

type Connection interface {
	FetchPayment(string) (Payment, error)
	GetAllPayments(int, sortByEnum, sortOrderEnum, ...time.Time) *PaymentBatcher
	FetchAuthorization(string) (*Authorization, error)
	FetchCapture(string) (*Capture, error)
	NewCreditCardPayment() *CreditCardPayment
	NewPaypalPayment(Redirects, *PaypalPayerInfo) (*PaypalPayment, error)
	Execute(*url.URL) error
	FetchRefund(string) (*Refund, error)
	FetchSale(string) (*Sale, error)

	authenticate() error
	send(*request) error
}

// Authorization response
type tokeninfo struct {
	// The access token issued by the authorization server.
	AccessToken string `json:"access_token,omitempty"`

	// The refresh token, which can be used to obtain new access tokens using the
	// same authorization grant as described in OAuth2.0 RFC6749 - Section 6.
	RefreshToken string `json:"refresh_token,omitempty"`

	// The type of the token issued as described in OAuth2.0 RFC6749 - Section 7.1
	// Value is case insensitive.
	TokenType string `json:"token_type,omitempty"`

	// The lifetime of the access token in seconds. After the access token
	// expires, use the refresh_token to refresh the access token.
	ExpiresIn uint `json:"expires_in,omitempty"`

	// Not sure what this is for. Appears once in a response example, but no
	// explanation is given anywhere.
	AppId string `json:"app_id,omitempty"`

	// THERE SEEM TO BE DIFFERENT DOCS IN DIFFERENT PLACES FOR THIS!!!
	//
	// Required if different from the scope requested by the client. For a list of
	// possible values, see the attributes table.
	// https://developer.paypal.com/docs/integration/direct/identity/attributes/
	//
	// Scopes expressed in the form of resource URL endpoints. The value of the
	// scope parameter is expressed as a list of space-delimited, case-sensitive
	// strings.
	// Value assigned by PayPal.
	Scope string `json:"scope,omitempty"`

	////////////// Derived fields
	expiration time.Time

	reauthAttempts uint

	mux sync.RWMutex

	// Handles the case where an error is received instead
	*identity_error
}

type relatedResources []resource

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

func (self *connection) FetchPayment(payment_id string) (Payment, error) {
	// For determining Paypal or CreditCard payment
	var f struct {
		Payer struct {
			PaymentMethod PaymentMethodEnum
		}
		*payment_error
	}

	var req = &request{
		method:   method.Get,
		path:     path.Join(_paymentsPath, payment_id),
		body:     nil,
		response: &f,
	}

	if err := self.send(req); err != nil {
		return nil, err
	}

	var pymt Payment

	switch f.Payer.PaymentMethod {
	case PaymentMethod.PayPal:
		pymt = &PaypalPayment{}

	case PaymentMethod.CreditCard:
		pymt = &CreditCardPayment{}

	default:
		return nil, fmt.Errorf("Unknown payment method %q while unmarshaling",
			f.Payer.PaymentMethod)
	}

	if err := json.Unmarshal(req.responseData, pymt); err != nil {
		return nil, err
	}

	return pymt, nil
}

/*
@struct _shared --drop_ctor
	*connection
	Id 						string `json:"id,omitempty"` 						--read
	CreateTime 		dateTime `json:"create_time,omitempty"` --read
	UpdateTime 		dateTime `json:"update_time,omitempty"` --read
	State 				StateEnum `json:"state,omitempty"` 			--read
	ParentPayment string `json:"parent_payment,omitempty"` --read
	Links 				links `json:"links,omitempty"`
	*identity_error
*/

func (self *_shared) FetchParentPayment() (Payment, error) {
	return self.FetchPayment(self.private.ParentPayment)
}

/*
// State items are: pending, authorized, captured, partially_captured, expired,
// 									voided
@struct Authorization --drop_ctor
	_shared
	Amount 						amount `json:"amount"` --read
	BillingAgreementId string `json:"billing_agreement_id"` --read
	PaymentMode 			paymentModeEnum `json:"payment_mode"` --read
	ReasonCode 				reasonCodeEnum `json:"reason_code"` --read
	ValidUntil 				dateTime `json:"valid_until"` --read
	ClearingTime 			string `json:"clearing_time"` --read
	ProtectionElig 		protectionEligEnum `json:"protection_eligibility"` --read
	ProtectionEligType protectionEligTypeEnum `json:"protection_eligibility_type"` --read
	FmfDetails 				fmfDetails `json:"fmf_details"` --read


// State values are: pending, completed, refunded, partially_refunded
@struct Capture
	_shared
	Amount 				 amount `json:"amount"` --read
	TransactionFee currency `json:"transaction_fee"` --read --write
	IsFinalCapture bool `json:"is_final_capture,omitempty"` --read --write


// State values are: pending; completed; refunded; partially_refunded
// TODO: PendingReason appears in the old docs under the general Sale object description
// but not under the lower "sale object" definition. The new docs have it
// marked as [DEPRECATED] in one area, but not another.
@struct Sale
	_shared
	Amount 										amount `json:"amount"` --read
	Description 							string `json:"description,omitempty"` --read --write
	TransactionFee 						currency `json:"transaction_fee"` --read --write
	ReceivableAmount 					currency `json:"receivable_amount"` --read --write
	PendingReason 						pendingReasonEnum `json:"pending_reason"` --read
	PaymentMode 							paymentModeEnum `json:"payment_mode"` --read
	ExchangeRate 							string `json:"exchange_rate"` --read
	FmfDetails 								fmfDetails `json:"fmf_details"` --read
	ReceiptId 								string `json:"receipt_id"` --read
	ReasonCode 								reasonCodeEnum `json:"reason_code"` --read
	ProtectionEligibility 		protectionEligEnum `json:"protection_eligibility"` --read
	ProtectionEligibilityType protectionEligTypeEnum `json:"protection_eligibility_type"` --read
	ClearingTime 							string `json:"clearing_time"` --read
	BillingAgreementId 				string `json:"billing_agreement_id"` --read


// State items are: pending; completed; failed
@struct Refund
	_shared
	Amount 			amount `json:"amount"` --read
	Description string `json:"description,omitempty"` --read --write
	Reason 			string `json:"reason,omitempty"` --read --write
	SaleId 			string `json:"sale_id,omitempty"` --read
	CaptureId 	string `json:"capture_id,omitempty"` --read


// Amount Object
//  A`Transaction` object also may have an `ItemList`, which has dollar amounts.
//  These amounts are used to calculate the `Total` field of the `Amount` object
//
//	All other uses of `Amount` do have `shipping`, `shipping_discount` and
// `subtotal` to calculate the `Total`.
@struct amount
	Currency CurrencyTypeEnum `json:"currency"` --read
	Total 	 float64 					`json:"total,string"` --read
	Details *details 					`json:"details,omitempty"` --read --write
*/

func (self amount) validate() (err error) {
	if err = self.private.Currency.validate(); err != nil {
		return err
	}

	err = checkFloat7_10("Amount.Total", &self.private.Total, false)
	if err != nil {
		return err
	}
	return nil
}

// The Details can all be read/write because the Amount object is read only,
// so it gets a copy anyway.
// No need to validate because its values are calculated or validated when set.
/*
@struct details
	// Amount of the subtotal of the items. REQUIRED if line items are specified.
	// 10 chars max, with support for 2 decimal places
	Subtotal float64 `json:"subtotal,string,omitempty"` --read

	// Amount charged for tax. 10 chars max, with support for 2 decimal places
	Tax float64 `json:"tax,string,omitempty"` --read

	// Amount charged for shipping. 10 chars max, with support for 2 decimal places
	Shipping float64 `json:"shipping,string,omitempty"` --read --write

	// Amount being charged for handling fee. When `payment_method` is `paypal`
	HandlingFee float64 `json:"handling_fee,string,omitempty"` --read --write

	// Amount being charged for insurance fee. When `payment_method` is `paypal`
	Insurance float64 `json:"insurance,string,omitempty"` --read --write

	// Amount being discounted for shipping fee. When `payment_method` is `paypal`
	ShippingDiscount float64 `json:"shipping_discount,string,omitempty"` --read --write
*/

/*
@struct link
	Href 	 string `json:"href,omitempty"`
	Rel 	 relTypeEnum `json:"rel,omitempty"`
	Method string `json:"method,omitempty"`
*/

type links []link

func (l links) get(r relTypeEnum) (string, string) {
	for i, _ := range l {
		if l[i].private.Rel == r {
			return l[i].private.Href, l[i].private.Method
		}
	}
	return "", ""
}

/*
// Base object for all financial value related fields (balance, payment due, etc.)
@struct currency
	Currency string `json:"currency"` --read --write
	Value 	 string `json:"value"` --read --write
*/

/*
// This object represents Fraud Management Filter (FMF) details for a payment.
@struct fmfDetails
	FilterType 	fmfFilterEnum `json:"filter_type"` --read
	FilterID 		filterIdEnum `json:"filter_id"` --read
	Name 				string `json:"name"` --read
	Description string `json:"description"` --read
*/

// Address: This object is used for billing or shipping addresses.
type Address struct {
	// Line 1 of the address (e.g., Number, street, etc). 100 characters max.
	// Required.
	Line1 string `json:"line1,omitempty"`

	// Line 2 of the address (e.g., Suite, apt #, etc). 100 characters max.
	Line2 string `json:"line2,omitempty"`

	// City name. 50 characters max. Required.
	City string `json:"city,omitempty"`

	// 2-letter country code. 2 characters max. Required.
	CountryCode CountryCodeEnum `json:"country_code,omitempty"`

	// Zip code or equivalent is usually required for countries that have them.
	// 20 characters max. Required in certain countries.
	PostalCode string `json:"postal_code,omitempty"`

	// 2-letter code for US states, and the equivalent for other countries.
	// 100 characters max.
	State string `json:"state,omitempty"`

	// Phone number in E.123 format. 50 characters max.
	Phone string `json:"phone,omitempty"`

	// Address normalization status, returned only for payers from Brazil. Allowed
	// values: UNKNOWN, UNNORMALIZED_USER_PREFERRED, NORMALIZED, UNNORMALIZED
	NormalizationStatus normStatusEnum `json:"normalization_status,omitempty"`

	// Address status. Allowed values: CONFIRMED, UNCONFIRMED
	Status addressStatusEnum `json:"status,omitempty"`
}

func (self *Address) validate() (err error) {
	if err = checkStr("Address.Line1", &self.Line1, 100, true, true); err != nil {
		return err
	}
	if err = checkStr("Address.Line2", &self.Line2, 100, false, true); err != nil {
		return err
	}
	if err = checkStr("Address.City", &self.City, 50, true, true); err != nil {
		return err
	}
	if err = self.CountryCode.validate(); err != nil {
		return err
	}
	// TODO: If the country is the united states, should verify that a State was given
	if err = checkStr("Address.State", &self.State, 100, false, true); err != nil {
		return err
	}
	if err = checkStr("Address.Phone", &self.Phone, 50, false, true); err != nil {
		return err
	}
	// PostalCode "0" is used for countries that do not have postal codes.
	// TODO: Provide an marshaler/unmarshaler that handles "0"
	return checkStr("Address.PostalCode", &self.PostalCode, 20, true, true)

	// TODO: Anything should be done with NormalizationStatus?
	// TODO: Anything should be done with AddressStatus?
}

type ShippingAddress struct {
	// Line 1 of the address (e.g., Number, street, etc). 100 characters max.
	// Required.
	Line1 string `json:"line1,omitempty"`

	// Line 2 of the address (e.g., Suite, apt #, etc). 100 characters max.
	Line2 string `json:"line2,omitempty"`

	// City name. 50 characters max. Required.
	City string `json:"city,omitempty"`

	// 2-letter country code. 2 characters max. Required.
	CountryCode CountryCodeEnum `json:"country_code,omitempty"`

	// Zip code or equivalent is usually required for countries that have them.
	// 20 characters max. Required in certain countries.
	PostalCode string `json:"postal_code,omitempty"`

	// 2-letter code for US states, and the equivalent for other countries.
	// 100 characters max.
	State string `json:"state,omitempty"`

	// Phone number in E.123 format. 50 characters max.
	Phone string `json:"phone,omitempty"`

	// Address normalization status, returned only for payers from Brazil. Allowed
	// values: UNKNOWN, UNNORMALIZED_USER_PREFERRED, NORMALIZED, UNNORMALIZED
	NormalizationStatus normStatusEnum `json:"normalization_status,omitempty"`

	// Address status. Allowed values: CONFIRMED, UNCONFIRMED
	Status addressStatusEnum `json:"status,omitempty"`

	// Name of the recipient at this address. 50 characters max. Required
	RecipientName string `json:"recipient_name,omitempty"`

	// Address type. Must be one of the following: `residential`, `business`, or
	// `mailbox`.
	Type AddressTypeEnum `json:"type,omitempty"`
}

func (self *ShippingAddress) validate() (err error) {
	if self == nil {
		return nil
	}

	err = checkStr(
		"ShippingAddress.RecipientName", &self.RecipientName, 50, true, true)
	if err != nil {
		return err
	}

	if err = self.Type.validate(); err != nil {
		return err
	}

	// Borrow the `.vaildate()` method of Address for now.
	var a = Address{
		self.Line1, self.Line2, self.City, self.CountryCode, self.PostalCode,
		self.State, self.Phone, self.NormalizationStatus, self.Status,
	}
	return a.validate()
}

type identityAddress struct {
	StreetAddress string `json:"street_address,omitempty"`
	Locality      string `json:"locality,omitempty"`
	Region        string `json:"region,omitempty"`
	PostalCode    string `json:"postal_code,omitempty"`
	Country       string `json:"country,omitempty"`
}

type userinfo struct {
	UserId     string `json:"user_id,omitempty"`
	Sub        string `json:"sub,omitempty"`
	Name       string `json:"name,omitempty"`
	GivenName  string `json:"given_name,omitempty"`
	FamilyName string `json:"family_name,omitempty"`
	MiddleName string `json:"middle_name,omitempty"`

	Email         string `json:"email,omitempty"`
	EmailVerified string `json:"email_verified,omitempty"`
	PhoneNumber   string `json:"phone_number,omitempty"`

	Address  identityAddress `json:"address,omitempty"`
	Zoneinfo string          `json:"zoneinfo,omitempty"`
	Locale   string          `json:"locale,omitempty"`

	Gender    string `json:"gender,omitempty"`
	Birthdate string `json:"birthdate,omitempty"`
	AgeRange  string `json:"age_range,omitempty"`
	Picture   string `json:"picture,omitempty"`

	VerifiedAccount bool   `json:"verified_account,omitempty"`
	AccountType     string `json:"account_type,omitempty"`
	PayerId         string `json:"payer_id,omitempty"`
}

type fundingInstrument struct {
	CreditCard      *credit_card       `json:"credit_card,omitempty"`
	CreditCardToken *credit_card_token `json:"credit_card_token,omitempty"`
}
type _cc_details struct {
	PayerId     string             `json:"payer_id,omitempty"`
	Type        CreditCardTypeEnum `json:"type,omitempty"`
	ExpireMonth string             `json:"expire_month,omitempty"`
	ExpireYear  string             `json:"expire_year,omitempty"`
}
type credit_card struct {
	Id     string `json:"id,omitempty"`
	Number string `json:"number,omitempty"`
	_cc_details
	Cvv2           string    `json:"cvv2,omitempty"`
	FirstName      string    `json:"first_name,omitempty"`
	LastName       string    `json:"last_name,omitempty"`
	BillingAddress *Address  `json:"billing_address,omitempty"`
	State          StateEnum `json:"state,omitempty"`
	ValidUntil     string    `json:"valid_until,omitempty"`
}
type credit_card_token struct {
	CreditCardId string `json:"credit_card_id,omitempty"`
	Last4        string `json:"last4,omitempty"`
	_cc_details
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
		method:   method.Get,
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
