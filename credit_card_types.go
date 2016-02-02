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
