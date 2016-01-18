package gopal

// Source of the funds for this payment represented by a PayPal account or a
// credit card.
type payer struct {
	// Payment method used. Must be either credit_card or paypal. Required.
	PaymentMethod PaymentMethod `json:"payment_method,omitempty"`

	// A list of funding instruments for the current payment
	FundingInstruments []*funding_instrument `json:"funding_instruments,omitempty"`

	// Information related to the payer.
	PayerInfo *payer_info `json:"payer_info,omitempty"`

	// Status of the payer’s PayPal account. Only supported when the
	// payment_method is set to paypal. Allowed values: VERIFIED or UNVERIFIED.
	Status payerStatus `json:"status,omitempty"`
}

// This object is pre-filled by PayPal when the payment_method is paypal.
type payer_info struct {
	// Email address representing the payer. 127 characters max.
	Email string `json:"email,omitempty"`

	// Salutation of the payer.
	Salutation string `json:"salutation,omitempty"`

	// First name of the payer. Value assigned by PayPal.
	First_name string `json:"first_name,omitempty"`

	// Middle name of the payer. Value assigned by PayPal.
	MiddleName string `json:"middle_name,omitempty"`

	// Last name of the payer. Value assigned by PayPal.
	Last_name string `json:"last_name,omitempty"`

	// Suffix of the payer.
	Suffix string `json:"suffix,omitempty"`

	// PayPal assigned Payer ID. Value assigned by PayPal.
	Payer_id string `json:"payer_id,omitempty"`

	// Phone number representing the payer. 20 characters max.
	Phone string `json:"phone,omitempty"`

	// Two-letter registered country code of the payer to identify the buyer country.
	CountryCode country_code `json:"country_code,omitempty"`

	// Shipping address of payer PayPal account. Value assigned by PayPal.
	Shipping_address *ShippingAddress `json:"shipping_address,omitempty"`

	// Payer’s tax ID type. Allowed values: BR_CPF or BR_CNPJ. Only supported when
	// the payment_method is set to paypal.
	TaxIdType TaxIdType `json:"tax_id_type,omitempty"`

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
	CountryCode country_code `json:"country_code,omitempty"`

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
	NormalizationStatus normalizationStatus `json:"normalization_status,omitempty"`

	// Address status. Allowed values: CONFIRMED, UNCONFIRMED
	Status addressStatus `json:"status,omitempty"`
}

type ShippingAddress struct {
	Address

	// Name of the recipient at this address. 50 characters max. Required
	RecipientName string `json:"recipient_name,omitempty"`

	// Address type. Must be one of the following: `residential`, `business`, or
	// `mailbox`.
	Type addressType `json:"type,omitempty"`
}

type identity_address struct {
	Street_address string `json:"street_address,omitempty"`
	Locality       string `json:"locality,omitempty"`
	Region         string `json:"region,omitempty"`
	Postal_code    string `json:"postal_code,omitempty"`
	Country        string `json:"country,omitempty"`
}

type userinfo struct {
	User_id     string `json:"user_id,omitempty"`
	Sub         string `json:"sub,omitempty"`
	Name        string `json:"name,omitempty"`
	Given_name  string `json:"given_name,omitempty"`
	Family_name string `json:"family_name,omitempty"`
	Middle_name string `json:"middle_name,omitempty"`

	Email          string `json:"email,omitempty"`
	Email_verified string `json:"email_verified,omitempty"`
	Phone_number   string `json:"phone_number,omitempty"`

	Address  identity_address `json:"address,omitempty"`
	Zoneinfo string           `json:"zoneinfo,omitempty"`
	Locale   string           `json:"locale,omitempty"`

	Gender    string `json:"gender,omitempty"`
	Birthdate string `json:"birthdate,omitempty"`
	Age_range string `json:"age_range,omitempty"`
	Picture   string `json:"picture,omitempty"`

	Verified_account bool   `json:"verified_account,omitempty"`
	Account_type     string `json:"account_type,omitempty"`
	Payer_id         string `json:"payer_id,omitempty"`
}

type funding_instrument struct {
	CreditCard      *credit_card       `json:"credit_card,omitempty"`
	CreditCardToken *credit_card_token `json:"credit_card_token,omitempty"`
}
type _cc_details struct {
	Payer_id     string         `json:"payer_id,omitempty"`
	Type         CreditCardType `json:"type,omitempty"`
	Expire_month string         `json:"expire_month,omitempty"`
	Expire_year  string         `json:"expire_year,omitempty"`
}
type credit_card struct {
	Id     string `json:"id,omitempty"`
	Number string `json:"number,omitempty"`
	_cc_details
	Cvv2            string   `json:"cvv2,omitempty"`
	First_name      string   `json:"first_name,omitempty"`
	Last_name       string   `json:"last_name,omitempty"`
	Billing_address *Address `json:"billing_address,omitempty"`
	State           state    `json:"state,omitempty"`
	Valid_until     string   `json:"valid_until,omitempty"`
}
type credit_card_token struct {
	Credit_card_id string `json:"credit_card_id,omitempty"`
	Last4          string `json:"last4,omitempty"`
	_cc_details
}
