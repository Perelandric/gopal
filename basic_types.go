package gopal

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

//go:generate Golific $GOFILE

type dateTime string // TODO: How should this be done? [Un]Marshalers?

// TODO: only needed until Golific acknowledges more types
type fundingInstruments []*fundingInstrument
type Transactions []*Transaction
type Items []*Item

type resource interface {
	errorable
	getPath() string
	Amount() amount
}

type relatedResources []resource

type refundable interface {
	resource
	getRefundPath() string
}

type request struct {
	method    methodEnum
	path      string
	body      interface{}
	response  errorable
	isAuthReq bool
}

type connection struct {
	server     ServerEnum
	id, secret string
	tokeninfo  tokeninfo
	client     http.Client
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

	// Handles the case where an error is received instead
	*identity_error
}

type Redirects struct {
	Return string `json:"return_url,omitempty"`
	Cancel string `json:"cancel_url,omitempty"`
}

func (self *Redirects) validate() error {
	for _, s := range [2]string{self.Return, self.Cancel} {
		u, err := url.Parse(s)
		if err != nil {
			return err
		}

		if len(u.Scheme) == 0 {
			return fmt.Errorf("URL Scheme is required. Found %q\n", s)
		}
		if len(u.Host) == 0 {
			return fmt.Errorf("URL Host is required. Found %q\n", s)
		}
	}
	return nil
}

/**********************

Payment Object

**********************/

/*
// TODO: Add `billing_agreement_tokens`, `payment_instruction`
@struct Payment
	*connection
	ExperienceProfileId string `json:"experience_profile_id"` --read --write
	Intent							intentEnum `json:"intent,omitempty"` --read
	Payer 							payer `json:"payer,omitempty"` --read
	Transactions				Transactions `json:"transactions,omitempty"` --read
	RedirectUrls				Redirects `json:"redirect_urls,omitempty"` --read
	State 							stateEnum `json:"state,omitempty"` --read
	Id 									string `json:"id,omitempty"` --read
	FailureReason				FailureReasonEnum `json:"failure_reason,omitempty"` --read
	CreateTime 					dateTime `json:"create_time,omitempty"` --read
	UpdateTime 					dateTime `json:"update_time,omitempty"` --read
	Links 							links `json:"links,omitempty"` --read
	*payment_error
*/

// TODO: Needs to validate some sub-properties that are valid only when
// Payer.PaymentMethod is "paypal"
func (self *Payment) validate() (err error) {
	if len(self.private.Transactions) == 0 {
		return fmt.Errorf("A Payment needs at least one Transaction")
	}

	for _, t := range self.private.Transactions {
		if err = t.validate(); err != nil {
			return err
		}
	}

	// TODO: More validation

	return self.private.Payer.validate()
}

func (self *Payment) calculateToAuthorize() {
	for _, t := range self.private.Transactions {
		t.calculateToAuthorize()
	}
}

/*
@struct Transaction
	ItemList 				*itemList `json:"item_list,omitempty"` --read
	Amount 					amount `json:"amount"` --read
	RelatedResources relatedResources `json:"related_resources,omitempty"` --read
	Description 		string `json:"description,omitempty"` --read --write
	InvoiceNumber 	string `json:"invoice_number,omitempty"` --read --write
	Custom 					string `json:"custom,omitempty"` --read --write
	SoftDescriptor 	string `json:"soft_descriptor,omitempty"` --read --write
	PaymentOptions  paymentOptions `json:"payment_options,omitempty"` --read --write
*/

func (self *Transaction) validate() (err error) {
	if err = self.private.ItemList.validate(); err != nil {
		return err
	}

	// These can be truncated with a warning if too long
	checkStr("Transaction.Description", &self.Description, 127, false, false)
	checkStr("Transaction.Custom", &self.Custom, 256, false, false)
	checkStr("Transaction.SoftDescriptor", &self.SoftDescriptor, 22, false, false)

	// TODO: More validation... check docs

	err = checkStr(
		"Transaction.InvoiceNumber", &self.InvoiceNumber, 256, false, true)
	if err != nil {
		return err
	}

	return self.private.Amount.validate()
}

func (self *Transaction) calculateToAuthorize() {
	// Calculate totals from itemList
	for _, item := range self.private.ItemList.private.Items {
		self.private.Amount.Details.private.Subtotal = roundTwoDecimalPlaces(
			self.private.Amount.Details.private.Subtotal + (item.Price * float64(item.Quantity)))

		self.private.Amount.Details.private.Tax = roundTwoDecimalPlaces(
			self.private.Amount.Details.private.Tax + (item.Tax * float64(item.Quantity)))
	}

	// Set Total, which is sum of Details
	self.private.Amount.private.Total = roundTwoDecimalPlaces(
		self.private.Amount.Details.private.Subtotal +
			self.private.Amount.Details.private.Tax +
			self.private.Amount.Details.private.Shipping +
			self.private.Amount.Details.private.Insurance -
			self.private.Amount.Details.private.ShippingDiscount)
}

/*
@struct itemList
	Items           Items          	 `json:"items,omitempty"` --read
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"` --read
*/

func (self *itemList) validate() (err error) {
	if self == nil {
		return nil
	}
	if len(self.private.Items) == 0 {
		return fmt.Errorf("Transaction item list must have at least one Item")
	}

	for _, item := range self.private.Items {
		if err = item.validate(); err != nil {
			return err
		}
	}
	return self.private.ShippingAddress.validate()
}

/*
@struct Item
	Currency 		CurrencyTypeEnum 	`json:"currency"` --read
	Quantity 		int64 			`json:"quantity,string"` --read --write
	Name 				string 			`json:"name"` --read --write
	Price 			float64 		`json:"price,string"` --read --write
	Tax 				float64 		`json:"tax,omitempty"` --read --write
	Sku 				string 			`json:"sku,omitempty"` --read --write
	Description string 			`json:"description,omitempty"` --read --write
*/

func (self *Item) validate() (err error) {
	if self.Tax < 0 { // TODO: No other validation here???
		return fmt.Errorf("%q must not be a negative number", "Item.Tax")
	}
	if err = checkStr("Item.Name", &self.Name, 127, true, true); err != nil {
		return err
	}
	if err = checkStr("Item.Sku", &self.Sku, 50, false, true); err != nil {
		return err
	}
	_ = checkStr("Item.Description", &self.Description, 127, false, false)

	if err = checkFloat7_10("Item.Price", &self.Price, true); err != nil {
		return err
	}

	return checkInt10("Item.Quantity", self.Quantity, true)
}

/*
@struct _shared --drop_ctor
	*connection
	Id 						string `json:"id,omitempty"` 						--read
	CreateTime 		dateTime `json:"create_time,omitempty"` --read
	UpdateTime 		dateTime `json:"update_time,omitempty"` --read
	State 				stateEnum `json:"state,omitempty"` 			--read
	ParentPayment string `json:"parent_payment,omitempty"` --read
	Links 				links `json:"links,omitempty"`
	*identity_error
*/

func (self *_shared) FetchParentPayment() (*Payment, error) {
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
	Total 	 float64 					`json:"total"` --read
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
	// Amount charged for shipping. 10 chars max, with support for 2 decimal places
	Shipping float64 `json:"shipping,omitempty"` --read --write

	// Amount of the subtotal of the items. REQUIRED if line items are specified.
	// 10 chars max, with support for 2 decimal places
	Subtotal float64 `json:"subtotal,omitempty"` --read --write

	// Amount charged for tax. 10 chars max, with support for 2 decimal places
	Tax float64 `json:"tax,omitempty"` --read --write

	// Amount being charged for handling fee. When `payment_method` is `paypal`
	HandlingFee float64 `json:"handling_fee,omitempty"` --read --write

	// Amount being charged for insurance fee. When `payment_method` is `paypal`
	Insurance float64 `json:"insurance,omitempty"` --read --write

	// Amount being discounted for shipping fee. When `payment_method` is `paypal`
	ShippingDiscount float64 `json:"shipping_discount,omitempty"` --read --write
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


// This object represents Fraud Management Filter (FMF) details for a payment.
@struct fmfDetails
	FilterType 	fmfFilterEnum `json:"filter_type"` --read
	FilterID 		filterIdEnum `json:"filter_id"` --read
	Name 				string `json:"name"` --read
	Description string `json:"description"` --read
*/

/*
// TODO: This is used both for Paypal and Credit cards. In the case of Paypal,
// many (most) of the fields may not be included in a requrest. Should we have
// a separate `PaypalPayerInfo` object? It could be embedded into the main
// `PayerInfo` object.

// This object is pre-filled by PayPal when the payment_method is paypal.
@struct PayerInfo
	// Email address representing the payer. 127 characters max.
	Email string `json:"email,omitempty"` --read --write

	// Salutation of the payer.
	Salutation string `json:"salutation,omitempty"` --read --write

	// Suffix of the payer.
	Suffix string `json:"suffix,omitempty"` --read --write

	// Two-letter registered country code of the payer to identify the buyer country.
	CountryCode CountryCodeEnum `json:"country_code,omitempty"` --read --write

	// Phone number representing the payer. 20 characters max.
	Phone string `json:"phone,omitempty"` --read --write

	// Payer’s tax ID type. Allowed values: BR_CPF or BR_CNPJ. Only supported when
	// the payment_method is set to paypal.
	TaxIdType TaxIdTypeEnum `json:"tax_id_type,omitempty"` --read --write

	// Payer’s tax ID. Only supported when the payment_method is set to paypal.
	TaxId string `json:"tax_id,omitempty"` --read --write

	// First name of the payer. Value assigned by PayPal.
	FirstName string `json:"first_name,omitempty"` --read

	// Middle name of the payer. Value assigned by PayPal.
	MiddleName string `json:"middle_name,omitempty"` --read

	// Last name of the payer. Value assigned by PayPal.
	LastName string `json:"last_name,omitempty"` --read

	// PayPal assigned Payer ID. Value assigned by PayPal.
	PayerId string `json:"payer_id,omitempty"` --read

	// Shipping address of payer PayPal account. Value assigned by PayPal.
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"` --read
*/

func (self *PayerInfo) validate() error {
	// TODO: Implement
	if self == nil {
		return nil
	}

	return nil
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
	Address

	// Name of the recipient at this address. 50 characters max. Required
	RecipientName string `json:"recipient_name,omitempty"`

	// Address type. Must be one of the following: `residential`, `business`, or
	// `mailbox`.
	Type AddressTypeEnum `json:"type,omitempty"`
}

func (self *ShippingAddress) validate() (err error) {
	err = checkStr(
		"ShippingAddress.RecipientName", &self.RecipientName, 50, true, true)
	if err != nil {
		return err
	}

	if err = self.Type.validate(); err != nil {
		return err
	}
	return self.Address.validate()
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
	State          stateEnum `json:"state,omitempty"`
	ValidUntil     string    `json:"valid_until,omitempty"`
}
type credit_card_token struct {
	CreditCardId string `json:"credit_card_id,omitempty"`
	Last4        string `json:"last4,omitempty"`
	_cc_details
}
