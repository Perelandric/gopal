package gopal

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

type resource interface {
	errorable
	getPath() string
	getAmount() amount
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

// State items are: pending, authorized, captured, partially_captured, expired,
// 									voided
type Authorization struct {
	_shared

	// Details about the amount
	Amount amount `json:"amount"`

	// ID of the billing agreement used as reference to execute this transaction.
	// Read only.
	BillingAgreementId string `json:"billing_agreement_id"`

	// Specifies the payment mode of the transaction.
	// Read only.
	PaymentMode paymentModeEnum `json:"payment_mode"`

	// Reason code, AUTHORIZATION, for a transaction state of `pending`.
	// Read only.
	ReasonCode reasonCodeEnum `json:"reason_code"`

	// Authorization expiration time and date as defined in RFC 3339 Section 5.6.
	// Read only.
	ValidUntil dateTime `json:"valid_until"`

	// Expected clearing time for eCheck transactions. Only supported when the
	// payment_method is set to paypal.
	// Read only.
	ClearingTime string `json:"clearing_time"`

	// The level of seller protection in force for the transaction. Only supported
	// when the payment_method is set to paypal.
	// Read only.
	ProtectionElig protectionEligEnum `json:"protection_eligibility"`

	// The kind of seller protection in force for the transaction. This property
	// is returned only when the protection_eligibility property is set to
	// ELIGIBLE or PARTIALLY_ELIGIBLE. Only supported when the `payment_method` is
	// set to `paypal`.
	// Read only.
	ProtectionEligType protectionEligTypeEnum `json:"protection_eligibility_type"`

	// Fraud Management Filter (FMF) details applied for the payment that could
	// result in accept, deny, or pending action. Returned in a payment response
	// only if the merchant has enabled FMF in the profile settings and one of the
	// fraud filters was triggered based on those settings. See "Fraud Management
	// Filters Summary" for more information.
	// Read only.
	FmfDetails fmfDetails `json:"fmf_details"`
}

// State values are: pending, completed, refunded, partially_refunded
type Capture struct {
	_shared

	// Details about the amount
	Amount amount `json:"amount"`

	// Transaction fee applicable for this payment.
	TransactionFee currency `json:"transaction_fee"`

	// If set to `true`, all remaining funds held by the authorization will be
	// released in the funding instrument. Default is `false`.
	IsFinalCapture bool `json:"is_final_capture,omitempty"`
}

// State values are: pending; completed; refunded; partially_refunded
type Sale struct {
	_shared

	// Details about the amount
	Amount amount `json:"amount"`

	// Description of sale.
	Description string `json:"description,omitempty"`

	// Transaction fee charged by PayPal for this transaction.
	TransactionFee currency `json:"transaction_fee"`

	// Net amount the merchant receives for this transaction in their receivable
	// currency. Returned only in cross-currency use cases where a merchant bills
	// a buyer in a non-primary currency for that buyer.
	ReceivableAmount currency `json:"receivable_amount"`

	// Reason the transaction is in pending state. Only supported when the
	// `payment_method` is set to `paypal`
	// Read only.
	// TODO: This appears in the old docs under the general Sale object description
	// but not under the lower "sale object" definition. The new docs have it
	// marked as [DEPRECATED] in one area, but not another.
	PendingReason pendingReasonEnum `json:"pending_reason"`

	// Specifies payment mode of the transaction. Only supported when the
	// `payment_method` is set to `paypal`.
	// Read only.
	PaymentMode paymentModeEnum `json:"payment_mode"`

	// Exchange rate applied for this transaction. Returned only in cross-currency
	// use cases where a merchant bills a buyer in a non-primary currency for that
	// buyer.
	// Read only.
	ExchangeRate string `json:"exchange_rate"`

	// Fraud Management Filter (FMF) details applied for the payment that could
	// result in accept, deny, or pending action. Returned in a payment response
	// only if the merchant has enabled FMF in the profile settings and one of the
	// fraud filters was triggered based on those settings. See "Fraud Management
	// Filters Summary" for more information.
	// Read only.
	FmfDetails fmfDetails `json:"fmf_details"`

	// Receipt ID is a 16-digit payment identification number returned for guest
	// users to identify the payment.
	// Read only.
	ReceiptId string `json:"receipt_id"`

	// Reason code for the transaction state being Pending or Reversed. Only
	// supported when the `payment_method` is set to `paypal`.
	// Read only.
	ReasonCode reasonCodeEnum `json:"reason_code"`

	// The level of seller protection in force for the transaction. Only supported
	// when the `payment_method` is set to `paypal`.
	// Read only.
	ProtectionEligibility protectionEligEnum `json:"protection_eligibility"`

	// The kind of seller protection in force for the transaction. This property
	// is returned only when the protection_eligibility property is set to
	// `ELIGIBLE` or `PARTIALLY_ELIGIBLE`. Only supported when the `payment_method`
	// is set to paypal. One or both of the allowed values can be returned.
	// Read only.
	ProtectionEligibilityType protectionEligTypeEnum `json:"protection_eligibility_type"`

	// Expected clearing time for eCheck transactions. Only supported when the
	// payment_method is set to paypal.
	// Read only.
	ClearingTime string `json:"clearing_time"`

	// ID of the billing agreement used as reference to execute this transaction.
	// Read only.
	BillingAgreementId string `json:"billing_agreement_id"`
}

// State items are: pending; completed; failed
type Refund struct {
	_shared

	// Details about the amount
	Amount amount `json:"amount"`

	// Description of what is being refunded
	Description string `json:"description,omitempty"`

	// Reason description for the Sale transaction being refunded.
	Reason string `json:"reason,omitempty"`

	// ID of the Sale transaction being refunded. One among sale_id or capture_id
	// will be returned based on the resource used to initiate refund.
	// Read Only.
	SaleId string `json:"sale_id,omitempty"`

	// ID of the sale transaction being refunded.
	// Read Only.
	CaptureId string `json:"capture_id,omitempty"`
}

type dateTime string // TODO: How should this be done? [Un]Marshalers?

type _shared struct {
	*connection

	// ID of the Sale/Refund/Capture/Authorization.
	// Read Only.
	Id string `json:"id,omitempty"`

	// Time of sale as defined in RFC 3339 Section 5.6
	// Read Only.
	CreateTime dateTime `json:"create_time,omitempty"`

	// Time that the resource was last updated.
	// Read Only.
	UpdateTime dateTime `json:"update_time,omitempty"`

	// State of the Sale/Refund/Capture/Auth
	// Read Only.
	State stateEnum `json:"state,omitempty"`

	// ID of the payment resource on which this transaction is based.
	// Read Only.
	ParentPayment string `json:"parent_payment,omitempty"`

	// HATEOAS links related to this call. Value generated by PayPal.
	// Read Only.
	Links links `json:"links,omitempty"`

	*identity_error
}

func (self *_shared) FetchParentPayment() (*Payment, error) {
	return self.FetchPayment(self.ParentPayment)
}

// Amount Object
//  A`Transaction` object also may have an `ItemList`, which has dollar amounts.
//  These amounts are used to calculate the `Total` field of the `Amount` object
//
//	All other uses of `Amount` do have `shipping`, `shipping_discount` and
// `subtotal` to calculate the `Total`.
type amount struct {
	// 3 letter currency code. PayPal does not support all currencies. REQUIRED.
	Currency CurrencyTypeEnum `json:"currency"`

	// Total amount charged from the payer to the payee. In case of a refund, this
	// is the refunded amount to the original payer from the payee. 10 characters
	// max with support for 2 decimal places. REQUIRED.
	Total string `json:"total"`

	Details *details `json:"details,omitempty"`
}

func (self *amount) setTotal(amt float64) error {
	if s, err := make10CharAmount(amt); err != nil {
		return err
	} else {
		self.Total = s
	}
	return nil
}

func (self amount) getCurrencyType() CurrencyTypeEnum {
	return self.Currency
}

func (self amount) totalAsFloat() float64 {
	f, err := strconv.ParseFloat(self.Total, 64)
	if err != nil {
		log.Printf("Invalid Amount.Total: %s\n", self.Total)
		return 0.0
	}
	return f
}

type links []link

func (l links) get(r relTypeEnum) (string, string) {
	for i, _ := range l {
		if l[i].Rel == r {
			return l[i].Href, l[i].Method
		}
	}
	return "", ""
}

type link struct {
	// URL of the related HATEOAS link you can use for subsequent calls
	Href string `json:"href,omitempty"`

	// Link relation that describes how this link relates to the previous call.
	// Examples include `self` (get details of the current call), `parent_payment`
	// (get details of the parent payment), or a related call such as `execute` or
	// `refund`
	Rel relTypeEnum `json:"rel,omitempty"`

	// The HTTP method required for the related call
	Method string `json:"method,omitempty"`
}

// Base object for all financial value related fields (balance, payment due, etc.)
type currency struct {
	// 3 letter currency code as defined by ISO 4217. Required.
	Currency string `json:"currency"`

	// amount up to N digit after the decimals separator as defined in ISO 4217
	// for the appropriate currency code. Required.
	Value string `json:"value"`
}

// This object represents Fraud Management Filter (FMF) details for a payment.
type fmfDetails struct {
	// Type of filter
	FilterType fmfFilterEnum `json:"filter_type"`

	// Name of the fraud management filter. For more information about filters,
	// see Fraud Management Filters Summary.
	FilterID filterIdEnum `json:"filter_id"`

	// Name of the filter. This property is reserved for future use.
	Name string `json:"name"`

	// Description of the filter. This property is reserved for future use.
	Description string `json:"description"`
}

// Source of the funds for this payment represented by a PayPal account or a
// credit card.
type payer struct {
	// Payment method used. Must be either credit_card or paypal. Required.
	PaymentMethod PaymentMethodEnum `json:"payment_method,omitempty"`

	// A list of funding instruments for the current payment
	FundingInstruments []*fundingInstrument `json:"funding_instruments,omitempty"`

	// Information related to the payer.
	PayerInfo *payerInfo `json:"payer_info,omitempty"`

	// Status of the payer’s PayPal account. Only supported when the
	// payment_method is set to paypal. Allowed values: VERIFIED or UNVERIFIED.
	Status payerStatusEnum `json:"status,omitempty"`
}

// This object is pre-filled by PayPal when the payment_method is paypal.
type payerInfo struct {
	// Email address representing the payer. 127 characters max.
	Email string `json:"email,omitempty"`

	// Salutation of the payer.
	Salutation string `json:"salutation,omitempty"`

	// First name of the payer. Value assigned by PayPal.
	FirstName string `json:"first_name,omitempty"`

	// Middle name of the payer. Value assigned by PayPal.
	MiddleName string `json:"middle_name,omitempty"`

	// Last name of the payer. Value assigned by PayPal.
	LastName string `json:"last_name,omitempty"`

	// Suffix of the payer.
	Suffix string `json:"suffix,omitempty"`

	// PayPal assigned Payer ID. Value assigned by PayPal.
	PayerId string `json:"payer_id,omitempty"`

	// Phone number representing the payer. 20 characters max.
	Phone string `json:"phone,omitempty"`

	// Two-letter registered country code of the payer to identify the buyer country.
	CountryCode CountryCodeEnum `json:"country_code,omitempty"`

	// Shipping address of payer PayPal account. Value assigned by PayPal.
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`

	// Payer’s tax ID type. Allowed values: BR_CPF or BR_CNPJ. Only supported when
	// the payment_method is set to paypal.
	TaxIdType TaxIdTypeEnum `json:"tax_id_type,omitempty"`

	// Payer’s tax ID. Only supported when the payment_method is set to paypal.
	TaxId string `json:"tax_id,omitempty"`
}

/// Address: This object is used for billing or shipping addresses.
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

type ShippingAddress struct {
	Address

	// Name of the recipient at this address. 50 characters max. Required
	RecipientName string `json:"recipient_name,omitempty"`

	// Address type. Must be one of the following: `residential`, `business`, or
	// `mailbox`.
	Type AddressTypeEnum `json:"type,omitempty"`
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
