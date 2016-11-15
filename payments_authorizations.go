package gopal

import "path"

//go:generate Golific $GOFILE

/*************************************************************

	AUTHORIZATIONS:  https://api.paypal.com/v1/payments/payment/authorization

	  Use the /authorization resource and related sub-resources to act up upon
	a previously created authorization. Options include retrieving, capturing,
	voiding, and reauthorizing authorizations.

**************************************************************/

/*
@struct drop_ctor
*/
// State items are:
// pending, authorized, captured, partially_captured, expired, voided
type __Authorization struct {
	_shared
	Amount             amount                 `gRead json:"amount"`
	BillingAgreementId string                 `gRead json:"billing_agreement_id"`
	PaymentMode        paymentModeEnum        `gRead json:"payment_mode"`
	ReasonCode         reasonCodeEnum         `gRead json:"reason_code"`
	ValidUntil         dateTime               `gRead json:"valid_until"`
	ClearingTime       string                 `gRead json:"clearing_time"`
	ProtectionElig     protectionEligEnum     `gRead json:"protection_eligibility"`
	ProtectionEligType protectionEligTypeEnum `gRead json:"protection_eligibility_type"`
	FmfDetails         fmfDetails             `gRead json:"fmf_details"`
}

// Implement the resource interface

func (self *Authorization) getPath() string {
	return path.Join(_authorizationPath, self._shared.private.Id)
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
		method:   method.Get,
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
		_shared:        _shared{},
		IsFinalCapture: is_final,
	}
	capt_req.private.Amount = *amt

	var capt_resp = new(Capture)

	if err := self.send(&request{
		method:   method.Post,
		path:     path.Join(_authorizationPath, self._shared.private.Id, _capture),
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
		method:   method.Post,
		path:     path.Join(_authorizationPath, self._shared.private.Id, _void),
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
		auth_req.private.Amount = self.private.Amount
	} else {
		auth_req.private.Amount = *amt
	}
	var auth_resp = new(Authorization)

	if err := self.send(&request{
		method:   method.Post,
		path:     path.Join(_authorizationPath, self._shared.private.Id, _reauthorize),
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
