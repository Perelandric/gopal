package main

//go:generate Golific $GOFILE

/*
This file contains types that are specific to creating Paypal payments. Several
types are used for both Paypal and credit cards, yet have restrictions for one
or the other.

Instead of trying to manage the restrictions using a single object, I've broken
them up into separate objects. This helps guarantee the correct data is sent.
*/

/*
// Source of the funds for this payment represented by a PayPal account.
@struct paypalPayer
  // Must be PaymentMethod.Paypal
	PaymentMethod      PaymentMethodEnum `json:"payment_method,omitempty"`

  // Status of the payer’s PayPal account. Allowed values: VERIFIED or UNVERIFIED.
	Status 						 payerStatusEnum `json:"status,omitempty"` --read

	FundingInstruments fundingInstruments `json:"funding_instruments,omitempty"` --read

	PayerInfo          *PaypalPayerInfo `json:"payer_info,omitempty"` --read
*/

func (self *creditCardPayer) validate() error {
	err := self.private.PaymentMethod.validate()
	if err != nil {
		return err
	}
	return nil
}

/*
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
