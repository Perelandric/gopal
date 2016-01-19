package gopal

import "path"

/*************************************************************

	AUTHORIZATIONS:  https://api.paypal.com/v1/payments/payment/authorization

	  Use the /authorization resource and related sub-resources to act up upon
	a previously created authorization. Options include retrieving, capturing,
	voiding, and reauthorizing authorizations.

**************************************************************/

// State items are: pending, authorized, captured, partially_captured, expired,
// 									voided
type Authorization struct {
	_trans

	// ID of the billing agreement used as reference to execute this transaction.
	BillingAgreementId string `json:"billing_agreement_id"`

	// Specifies the payment mode of the transaction.
	PaymentMode paymentModeEnum `json:"payment_mode"`

	// Reason code, AUTHORIZATION, for a transaction state of `pending`. Value
	// assigned by PayPal.
	ReasonCode reasonCodeEnum `json:"reason_code"`

	// Authorization expiration time and date as defined in RFC 3339 Section 5.6.
	// Value assigned by PayPal.
	ValidUntil dateTime `json:"valid_until"`

	// Expected clearing time for eCheck transactions. Only supported when the
	// payment_method is set to paypal. Value assigned by PayPal.
	ClearingTime string `json:"clearing_time"`

	// The level of seller protection in force for the transaction. Only supported
	// when the payment_method is set to paypal.
	ProtectionElig protectionEligEnum `json:"protection_eligibility"`

	// The kind of seller protection in force for the transaction. This property
	// is returned only when the protection_eligibility property is set to
	// ELIGIBLE or PARTIALLY_ELIGIBLE. Only supported when the `payment_method` is
	// set to `paypal`.
	ProtectionEligType protectionEligTypeEnum `json:"protection_eligibility_type"`

	// Fraud Management Filter (FMF) details applied for the payment that could
	// result in accept, deny, or pending action. Returned in a payment response
	// only if the merchant has enabled FMF in the profile settings and one of the
	// fraud filters was triggered based on those settings. See "Fraud Management
	// Filters Summary" for more information.
	FmfDetails fmfDetails `json:"fmf_details"`
}

// Implement the Resource interface

func (self *Authorization) getPath() string {
	return path.Join(_authorizationPath, self.Id)
}

/*************************************************************

	LOOK UP AN AUTHORIZATION
	GET https://api.paypal.com/v1/payments/authorization/{authorization_id}

Use this call to get details about authorizations.

**************************************************************/

func (self *connection) FetchAuthorization(
	auth_id string) (*Authorization, error) {

	var auth = &Authorization{}
	auth.connection = self

	if err := self.send(&request{
		method:   method.get,
		path:     path.Join(_authorizationPath, auth_id),
		body:     nil,
		response: auth,
	}); err != nil {
		return nil, err
	}

	return auth, nil
}

/*************************************************************

	CAPTURE AN AUTHORIZATION
	POST https://api.paypal.com/v1/payments/authorization/{authorization_id}/capture

Use this resource to capture and process a previously created authorization.
To use this resource, the original payment call must have the intent set to authorize.

**************************************************************/

func (self *Authorization) Capture(amt *amount, is_final bool) (*Capture, error) {
	var capt_req = &Capture{
		_trans: _trans{
			Amount: *amt,
		},
		IsFinalCapture: is_final,
	}
	var capt_resp = new(Capture)

	if err := self.send(&request{
		method:   method.post,
		path:     path.Join(_authorizationPath, self.Id, _capture),
		body:     capt_req,
		response: capt_resp,
	}); err != nil {
		return nil, err
	}
	return capt_resp, nil
}

/*************************************************************

	VOID AN AUTHORIZATION
	POST https://api.paypal.com/v1/payments/authorization/{authorization_id}/void

Use this call to void a previously authorized payment. Note a fully captured authorization
cannot be voided.

**************************************************************/

func (self *Authorization) Void() (*Authorization, error) {
	var void_resp = new(Authorization)

	if err := self.send(&request{
		method:   method.post,
		path:     path.Join(_authorizationPath, self.Id, _void),
		body:     nil,
		response: void_resp,
	}); err != nil {
		return nil, err
	}
	return void_resp, nil
}

/*************************************************************

	REAUTHORIZE A PAYMENT
	POST https://api.paypal.com/v1/payments/authorization/{authorization_id}/reauthorize

  Use this call to reauthorize a PayPal account payment. We recommend that you reauthorize a
payment after the initial 3-day honor period to ensure that funds are still available.

  You can reauthorize a payment only once 4 to 29 days after 3-day honor period for the
original authorization expires. If 30 days have passed from the original authorization,
you must create a new authorization instead. A reauthorized payment itself has a new 3-day
honor period. You can reauthorize a transaction once for up to 115% of the originally
authorized amount, not to exceed an increase of $75 USD.

  Note: You can only reauthorize PayPal account payments.

**************************************************************/

// TODO: Since only PayPal authorizations can be re-authorized, should I add a check
//			for this? Or should I just let the error come back from the server?
func (self *Authorization) ReauthorizeAmount(amt *amount) (*Authorization, error) {
	var auth_req = new(Authorization)
	if amt == nil {
		auth_req.Amount = self.Amount
	} else {
		auth_req.Amount = *amt
	}
	var auth_resp = new(Authorization)

	if err := self.send(&request{
		method:   method.post,
		path:     path.Join(_authorizationPath, self.Id, _reauthorize),
		body:     auth_req,
		response: auth_resp,
	}); err != nil {
		return nil, err
	}
	return auth_resp, nil
}

func (self *Authorization) Reauthorize() (*Authorization, error) {
	return self.ReauthorizeAmount(nil)
}
