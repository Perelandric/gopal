package gopal

import "fmt"

//go:generate Golific $GOFILE

/*
This file contains types that are specific to creating credit card payments.
Several types are used for both Paypal and credit cards, yet have restrictions
for one or the other.
*/

func (self *connection) NewCreditCardPayment() *CreditCardPayment {
	var pymt = CreditCardPayment{
		connection: self,
	}
	pymt.private.Intent = intent.Sale
	pymt.PaymentBase.private.Transactions = make([]*Transaction, 0)

	pymt.private.Payer.private.PaymentMethod = PaymentMethod.PayPal
	pymt.private.Payer.private.PayerInfo = nil

	return &pymt
}

/*
// TODO: Add `billing_agreement_tokens`, `payment_instruction`
@struct CreditCardPayment
	*connection
  PaymentBase
	ExperienceProfileId string `json:"experience_profile_id"` --read --write
	Payer 							creditCardPayer `json:"payer,omitempty"` --read
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
func (self *CreditCardPayment) validate() (err error) {
	if len(self.PaymentBase.private.Transactions) == 0 {
		return fmt.Errorf("A Payment needs at least one Transaction")
	}

	for _, t := range self.PaymentBase.private.Transactions {
		if err = t.validate(); err != nil {
			return err
		}
	}

	// TODO: More validation

	return self.private.Payer.validate()
}

// TODO: Finish implementation
func (self *CreditCardPayment) AddFundingInstrument(instrs ...fundingInstrument) error {
	for _, instr := range instrs {
		//		err := instr.validate()
		//		if err != nil {
		//			return err
		//		}
		self.private.Payer.private.FundingInstruments = append(
			self.private.Payer.private.FundingInstruments, &instr,
		)
	}
	return nil
}

// TODO: Should send a query string parameter `token=[some token]`
// TODO: Right now this is set up for Paypal Payments. Need to change for CCards
func (self *CreditCardPayment) Authorize() (to string, code int, err error) {
	if err = self.validate(); err != nil {
		return "", 0, err
	}

	self.calculateToAuthorize()

	// Create Totals
	var pymt CreditCardPayment

	err = self.send(&request{
		method:   method.Post,
		path:     _paymentsPath,
		body:     self,
		response: &pymt,
	})

	if err == nil {
		switch pymt.private.State {
		case state.Created:
			// Set url to redirect to PayPal site to begin approval process
			to, _ = pymt.private.Links.get(relType.ApprovalUrl)
			code = 303
		default:
			// otherwise cancel the payment and return an error
			err = UnexpectedResponse
		}
	}

	return to, code, err
}

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

func (self *PayerInfo) validate() error {
	// TODO: Implement
	if self == nil {
		return nil
	}

	return nil
}
