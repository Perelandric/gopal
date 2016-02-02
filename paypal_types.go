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
