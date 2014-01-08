package gopal

/*************************************************************

	AUTHORIZATIONS:  https://api.paypal.com/v1/payments/payment/authorization

	  Use the /authorization resource and related sub-resources to act up upon
	a previously created authorization. Options include retrieving, capturing,
	voiding, and reauthorizing authorizations.

**************************************************************/

type Authorizations struct {
	connection *Connection
}

type AuthorizationObject struct {
    _trans
    State       State  `json:"state,omitempty"` // TODO: Limit to allowed values
    Valid_until string `json:"valid_until,omitempty"`
    Links       links  `json:"links,omitempty"`

	*identity_error
	authorizations *Authorizations
}


/*************************************************************

	LOOK UP AN AUTHORIZATION
	GET https://api.paypal.com/v1/payments/authorization/{authorization_id}

Use this call to get details about authorizations.


REQUEST: Pass the authorization ID. (No object to send)

curl -v -X GET https://api.sandbox.paypal.com/v1/payments/authorization/2DC87612EK520411B \
-H "Content-Type:application/json" \
-H "Authorization:Bearer ENxom5Fof1KqAffEsXtxwEDa6E1HTEK__KVdIsaCYF8C"


RESPONSE: Returns a AUTHORIZATION object.

{
  "id": "2DC87612EK520411B",
  "create_time": "2013-06-25T21:39:15Z",
  "update_time": "2013-06-25T21:39:17Z",
  "state": "authorized",
  "amount": {
    "total": "7.47",
    "currency": "USD",
    "details": {
      "subtotal": "7.47"
    }
  },
  "parent_payment": "PAY-36246664YD343335CKHFA4AY",
  "valid_until": "2013-07-24T21:39:15Z",
  "links": [
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/authorization/2DC87612EK520411B",
      "rel": "self",
      "method": "GET"
    },
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/authorization/2DC87612EK520411B/capture",
      "rel": "capture",
      "method": "POST"
    },
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/authorization/2DC87612EK520411B/void",
      "rel": "void",
      "method": "POST"
    },
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/payment/PAY-36246664YD343335CKHFA4AY",
      "rel": "parent_payment",
      "method": "GET"
    }
  ]
}

**************************************************************/

func (self *Authorizations) Get(auth_id string) (*AuthorizationObject, error) {
    var auth = new(AuthorizationObject)
    var err = self.connection.make_request("GET",
													"payments/authorization/" + auth_id,
													nil, "", auth, false)
    if err != nil {
        return nil, err
    }
    auth.authorizations = self
    return auth, nil
}

func (self *AuthorizationObject) GetParentPayment() (*PaymentObject, error) {
    return self.authorizations.connection.Payments.Get(self.Parent_payment)
}


/*************************************************************

	CAPTURE AN AUTHORIZATION
	POST https://api.paypal.com/v1/payments/authorization/{authorization_id}/capture

Use this resource to capture and process a previously created authorization.
To use this resource, the original payment call must have the intent set to authorize.


REQUEST: Provide an authorization_id along with an amount object.
		For a partial capture, you can provide a lower amount.
		Additionally, you can explicitly indicate a final capture (prevent future captures)
		by setting the is_final_capture value to true.

curl -v https://api.sandbox.paypal.com/v1/payments/authorization/5RA45624N3531924N/capture \
-H "Content-Type:application/json" \
-H "Authorization:Bearer EMxItHE7Zl4cMdkvMg-f7c63GQgYZU8FjyPWKQlpsqQP" \
-d '{
  "amount":{
    "currency":"USD",
    "total":"4.54"
  },
  "is_final_capture":true
}'

RESPONSE:  Returns a CAPTURE object along with the STATE of the capture.

{
  "id": "6BA17599X0950293U",
  "create_time": "2013-05-06T22:32:24Z",
  "update_time": "2013-05-06T22:32:25Z",
  "amount": {
    "total": "4.54",
    "currency": "USD"
  },
  "is_final_capture": true,
  "state": "completed",
  "parent_payment": "PAY-44664305570317015KGEC5DI",
  "links": [
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/capture/6BA17599X0950293U",
      "rel": "self",
      "method": "GET"
    },
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/capture/6BA17599X0950293U/refund",
      "rel": "refund",
      "method": "POST"
    },
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/authorization/5RA45624N3531924N",
      "rel": "authorization",
      "method": "GET"
    },
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/payment/PAY-44664305570317015KGEC5DI",
      "rel": "parent_payment",
      "method": "GET"
    }
  ]
}

**************************************************************/

func (self *AuthorizationObject) Capture(amt *Amount, is_final bool) (*CaptureObject, error) {
    var capt_req = &CaptureObject{
		_trans: _trans{
			Amount: *amt,
		},
		Is_final_capture: is_final,
	}
	var capt_resp = new(CaptureObject)

    var err = self.authorizations.connection.make_request("POST",
															"payments/authorization/" + self.Id + "/capture",
															capt_req, "", capt_resp, false)
    if err != nil {
        return nil, err
    }
    return capt_resp, nil
}


/*************************************************************

	VOID AN AUTHORIZATION
	POST https://api.paypal.com/v1/payments/authorization/{authorization_id}/void

Use this call to void a previously authorized payment. Note a fully captured authorization
cannot be voided.


REQUEST: Pass the authorization ID in the resource URI.

curl -v -X POST https://api.sandbox.paypal.com/v1/payments/authorization/6CR34526N64144512/void \
-H "Content-Type:application/json" \
-H "Authorization:Bearer ENxom5Fof1KqAffEsXtxwEDa6E1HTEK__KVdIsaCYF8C"


RESPONSE:  Returns an AUTHORIZATION object.

{
  "id": "6CR34526N64144512",
  "create_time": "2013-05-06T21:56:50Z",
  "update_time": "2013-05-06T21:57:51Z",
  "state": "voided",
  "amount": {
    "total": "110.54",
    "currency": "USD",
    "details": {
      "subtotal": "110.54"
    }
  },
  "parent_payment": "PAY-0PL82432AD7432233KGECOIQ",
  "links": [
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/authorization/6CR34526N64144512",
      "rel": "self",
      "method": "GET"
    },
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/payment/PAY-0PL82432AD7432233KGECOIQ",
      "rel": "parent_payment",
      "method": "GET"
    }
  ]
}

**************************************************************/

func (self *AuthorizationObject) Void() (*AuthorizationObject, error) {
	var void_resp = new(AuthorizationObject)

    var err = self.authorizations.connection.make_request("POST",
															"payments/authorization/" + self.Id + "/void",
															nil, "", void_resp, false)
    if err != nil {
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



REQUEST: Pass the authorization id in the resource URI along with a new AUTHORIZATION object if you need to authorize a different amount.

curl -v https://api.sandbox.paypal.com/v1/payments/authorization/7YK97137TM7104916/reauthorize \
-H "Content-Type:application/json"  \
-H "Authorization: Bearer EBKoE2M-mctKH3OnORRVzUlxiromWnU5Mgz2PUZntQmt"  \
-d '{
  "amount":{
    "total":"7.00",
    "currency":"USD"
  }
}


RESPONSE:  Returns an AUTHORIZATION object with details of the authorization.

{
  "id": "8AA831015G517922L",
  "create_time": "2013-06-25T21:39:15Z",
  "update_time": "2013-06-25T21:39:17Z",
  "state": "authorized",
  "amount": {
    "total": "7.00",
    "currency": "USD"
  },
  "parent_payment": "PAY-7LD317540C810384EKHFAGYA",
  "valid_until": "2013-07-24T21:39:15Z",
  "links": [
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/authorization/8AA831015G517922L",
      "rel": "self",
      "method": "GET"
    },
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/payment/PAY-7LD317540C810384EKHFAGYA",
      "rel": "parent_payment",
      "method": "GET"
    },
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/authorization/8AA831015G517922L/capture",
      "rel": "capture",
      "method": "POST"
    }
  ]
}

**************************************************************/

// TODO: Since only PayPal authorizations can be re-authorized, should I add a check
//			for this? Or should I just let the error come back from the server?
func (self *AuthorizationObject) ReauthorizeAmount(amt *Amount) (*AuthorizationObject, error) {
	var auth_req = new(AuthorizationObject)
	if amt == nil {
		auth_req.Amount = self.Amount
	} else {
		auth_req.Amount = *amt
	}
	var auth_resp = new(AuthorizationObject)

    var err = self.authorizations.connection.make_request("POST",
															"payments/authorization/" + self.Id + "/reauthorize",
															auth_req, "", auth_resp, false)
    if err != nil {
        return nil, err
    }
    return auth_resp, nil
}

func (self *AuthorizationObject) Reauthorize() (*AuthorizationObject, error) {
	return self.ReauthorizeAmount(nil)
}

