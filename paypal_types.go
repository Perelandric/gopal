package gopal

import (
	"fmt"
	"net/url"
	"path"
)

//go:generate Golific $GOFILE

/*
This file contains types that are specific to creating Paypal payments. Several
types are used for both Paypal and credit cards, yet have restrictions for one
or the other.
*/

func (self *connection) NewPaypalPayment(urls Redirects) (*PaypalPayment, error) {

	var pymt = PaypalPayment{
		connection: self,
	}
	pymt.private.Intent = intent.Sale
	pymt.PaymentBase.private.Transactions = make([]*Transaction, 0)
	pymt.private.RedirectUrls = urls

	pymt.private.Payer.private.PaymentMethod = PaymentMethod.PayPal
	pymt.private.Payer.private.PaypalPayerInfo = nil

	if err := urls.validate(); err != nil {
		return nil, err
	}
	return &pymt, nil
}

/*
// TODO: Add `billing_agreement_tokens`, `payment_instruction`
@struct PaypalPayment
	*connection
  PaymentBase
	ExperienceProfileId string `json:"experience_profile_id"` --read --write
	Payer 							paypalPayer `json:"payer,omitempty"` --read
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
func (self *PaypalPayment) validate() (err error) {
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

// TODO: Should send a query string parameter `token=[some token]`
func (self *PaypalPayment) Authorize() (to string, code int, err error) {
	if err = self.validate(); err != nil {
		return "", 0, err
	}

	self.calculateToAuthorize()

	// Create Totals
	var pymt PaypalPayment

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

func (self *connection) Execute(u url.URL) error {
	var query = u.Query()

	var payerid = query.Get("PayerID")
	if payerid == "" {
		return fmt.Errorf("PayerID is missing\n")
	}

	var pymtid = query.Get("paymentId")
	if pymtid == "" {
		return fmt.Errorf("paymentId is missing\n")
	}

	var pymt PaypalPayment

	if err := self.send(&request{
		method:   method.Post,
		path:     path.Join(_paymentsPath, pymtid, _execute),
		body:     &paymentExecution{PayerID: payerid},
		response: &pymt,
	}); err != nil {
		return err
	}

	if pymt.private.State != state.Approved {
		return fmt.Errorf(
			"Payment with ID %q for payer %q was not approved\n", pymtid, payerid)
	}

	return nil
}

/*
// Source of the funds for this payment represented by a PayPal account.
@struct paypalPayer
  // Must be PaymentMethod.Paypal
	PaymentMethod      PaymentMethodEnum `json:"payment_method,omitempty"`

  // Status of the payer’s PayPal account. Allowed values: VERIFIED or UNVERIFIED.
	Status 						 payerStatusEnum `json:"status,omitempty"` --read

	PaypalPayerInfo    *PaypalPayerInfo `json:"payer_info,omitempty"` --read
*/

func (self *paypalPayer) validate() error {
	err := self.private.PaymentMethod.validate()
	if err != nil {
		return err
	}
	return nil
}

/*
// This object is pre-filled by PayPal when the payment_method is paypal.
@struct PaypalPayerInfo
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

func (self *PaypalPayerInfo) validate() error {
	// TODO: Implement
	if self == nil {
		return nil
	}

	return nil
}
