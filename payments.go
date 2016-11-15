package gopal

import (
	"encoding/json"
	"fmt"
	"log"
	"path"
	"time"
)

//go:generate Golific $GOFILE

type dateTime string // TODO: How should this be done? [Un]Marshalers?

// TODO: only needed until Golific acknowledges more types
type fundingInstruments []*fundingInstrument

/*
@struct
*/
type __CreditCardItem struct {
	Currency CurrencyTypeEnum `gRead json:"currency"`
	Quantity int64            `gRead gWrite json:"quantity,string"`
	Name     string           `gRead gWrite json:"name"`
	Price    float64          `gRead gWrite json:"price,string"`
	Sku      string           `gRead gWrite json:"sku,omitempty"`
	Url      string           `gRead gWrite json:"url,omitempty"`
}

func (ti *CreditCardItem) validate() (err error) {
	if err = checkStr("Item.Name", &ti.Name, 127, true, true); err != nil {
		return err
	}
	if err = checkStr("Item.Sku", &ti.Sku, 127, false, true); err != nil {
		return err
	}

	if err = checkFloat7_10("Item.Price", &ti.Price, true); err != nil {
		return err
	}

	return checkInt10("Item.Quantity", ti.Quantity, true)
}

/*
@struct
*/
type __PaypalItem struct {
	Currency    CurrencyTypeEnum `gRead json:"currency"`
	Quantity    int64            `gRead gWrite json:"quantity,string"`
	Name        string           `gRead gWrite json:"name"`
	Price       float64          `gRead gWrite json:"price,string"`
	Sku         string           `gRead gWrite json:"sku,omitempty"`
	Url         string           `gRead gWrite json:"url,omitempty"`
	Tax         float64          `gRead gWrite json:"tax,string,omitempty"`
	Description string           `gRead gWrite json:"description,omitempty"`
}

func (pi *PaypalItem) validate() (err error) {
	if pi.Tax < 0 { // TODO: No other validation here???
		return fmt.Errorf("%q must not be a negative number", "Item.Tax")
	}
	_ = checkStr("Item.Description", &pi.Description, 127, false, false)

	if err = checkStr("Item.Name", &pi.Name, 127, true, true); err != nil {
		return err
	}
	if err = checkStr("Item.Sku", &pi.Sku, 127, false, true); err != nil {
		return err
	}

	if err = checkFloat7_10("Item.Price", &pi.Price, true); err != nil {
		return err
	}

	return checkInt10("Item.Quantity", pi.Quantity, true)
}

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
@struct drop_ctor
*/
type ___shared struct {
	*connection
	Id            string    `gRead json:"id,omitempty"`
	CreateTime    dateTime  `gRead json:"create_time,omitempty"`
	UpdateTime    dateTime  `gRead json:"update_time,omitempty"`
	State         StateEnum `gRead json:"state,omitempty"`
	ParentPayment string    `gRead json:"parent_payment,omitempty"`
	Links         links     `json:"links,omitempty"`
	*identity_error
}

func (self *_shared) FetchParentPayment() (Payment, error) {
	return self.FetchPayment(self.private.ParentPayment)
}

/*
@struct
*/
// Amount Object
//  A `Transaction` object also may have an `ItemList`, which has dollar amounts.
//  These amounts are used to calculate the `Total` field of the `Amount` object
//
//	All other uses of `Amount` do have `shipping`, `shipping_discount` and
// `subtotal` to calculate the `Total`.
type __amount struct {
	Currency CurrencyTypeEnum `gRead json:"currency"`
	Total    float64          `gRead json:"total,string"`
	Details  *details         `gRead gWrite json:"details,omitempty"`
}

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

/*
@struct
*/
// The Details can all be read/write because the Amount object is read only,
// so it gets a copy anyway.
// No need to validate because its values are calculated or validated when set.
type __details struct {
	// Amount of the subtotal of the items. REQUIRED if line items are specified.
	// 10 chars max, with support for 2 decimal places
	Subtotal float64 `gRead json:"subtotal,string,omitempty"`

	// Amount charged for tax. 10 chars max, with support for 2 decimal places
	Tax float64 `gRead json:"tax,string,omitempty"`

	// Amount charged for shipping. 10 chars max, with support for 2 decimal places
	Shipping float64 `gRead gWrite json:"shipping,string,omitempty"`

	// Amount being charged for handling fee. When `payment_method` is `paypal`
	HandlingFee float64 `gRead gWrite json:"handling_fee,string,omitempty"`

	// Amount being charged for insurance fee. When `payment_method` is `paypal`
	Insurance float64 `gRead gWrite json:"insurance,string,omitempty"`

	// Amount being discounted for shipping fee. When `payment_method` is `paypal`
	ShippingDiscount float64 `gRead gWrite json:"shipping_discount,string,omitempty"`
}

/*
@struct
*/
type __link struct {
	Href   string      `json:"href,omitempty"`
	Rel    relTypeEnum `json:"rel,omitempty"`
	Method string      `json:"method,omitempty"`
}

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
@struct
*/
// Base object for all financial value related fields (balance, payment due, etc.)
type __currency struct {
	Currency string `gRead gWrite json:"currency"`
	Value    string `gRead gWrite json:"value"`
}

/*
@struct
*/
// This object represents Fraud Management Filter (FMF) details for a payment.
type __fmfDetails struct {
	FilterType  fmfFilterEnum `gRead json:"filter_type"`
	FilterID    filterIdEnum  `gRead json:"filter_id"`
	Name        string        `gRead json:"name"`
	Description string        `gRead json:"description"`
}

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
