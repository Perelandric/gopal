package main

//go:generate Golific $GOFILE

/*
  This file contains the counterparts to the types in `paypal_types.go`.
*/

/*
// Source of the funds for this payment represented by a credit card.
@struct creditCardPayer
  // Must be PaymentMethod.CreditCard
	PaymentMethod      PaymentMethodEnum `json:"payment_method,omitempty"` --read
	FundingInstruments fundingInstruments `json:"funding_instruments,omitempty"` --read
	PayerInfo          *PayerInfo `json:"payer_info,omitempty"` --read
*/

func (self *creditCardPayer) validate() error {
	err := self.private.PaymentMethod.validate()
	if err != nil {
		return err
	}
	return nil
}

/*
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
