package gopal

/*************************************************************

    CAPTURES:  https://api.paypal.com/v1/payments/payment/capture

	The /capture resource and sub-resources allow you to look up and
	refund captured payments.

**************************************************************/

type Captures struct {
	connection *Connection
}

type CaptureObject struct {
	_trans
	State            State `json:"state,omitempty"` // TODO: Limit to allowed values
	Is_final_capture bool  `json:"is_final_capture,omitempty"`
	Links            links `json:"links,omitempty"`

	RawData		[]byte `json:"-"`

	*identity_error
	captures *Captures
}

/*************************************************************

	LOOK UP A CAPTURED PAYMENT
	GET https://api.paypal.com/v1/payments/capture/{capture_id}

Use this call to get details about a captured payment.


REQUEST: Pass the capture id in the resource URI.

curl -v -X GET https://api.sandbox.paypal.com/v1/payments/capture/8F148933LY9388354 \
-H "Content-Type:application/json" \
-H "Authorization: Bearer EBKoE2M-mctKH3OnORRVzUlxiromWnU5Mgz2PUZntQmt"


RESPONSE: Returns a CAPTURE object with details about the capture.

{
  "id": "8F148933LY9388354",
  "amount": {
    "total": "110.54",
    "currency": "USD",
    "details": {
      "subtotal": "110.54"
    }
  },
  "is_final_capture": false,
  "state": "completed",
  "parent_payment": "PAY-8PT597110X687430LKGECATA",
  "links": [
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/capture/8F148933LY9388354",
      "rel": "self",
      "method": "GET"
    },
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/capture/8F148933LY9388354/refund",
      "rel": "refund",
      "method": "POST"
    },
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/payment/PAY-8PT597110X687430LKGECATA",
      "rel": "parent_payment",
      "method": "GET"
    }
  ]
}


**************************************************************/

func (self *Captures) Get(capt_id string) (*CaptureObject, error) {
	var capt = new(CaptureObject)
	var err = self.connection.make_request("GET",
		"payments/capture/"+capt_id,
		nil, "", capt, false)
	if err != nil {
		return nil, err
	}
	capt.captures = self
	return capt, nil
}

func (self *CaptureObject) GetParentPayment() (*PaymentObject, error) {
	return self.captures.connection.Payments.Get(self.Parent_payment)
}

/*************************************************************

	REFUND A CAPTURED PAYMENT
	POST https://api.paypal.com/v1/payments/capture/{capture_id}/refund

Use this call to refund a captured payment. Provide the capture_id in the URI and an amount object. For partial refunds, you can include a lower amount object.


REQUEST: Provide the capture_id in the URI and an AMOUNT object.

curl -v https://api.sandbox.paypal.com/v1/payments/capture/8F148933LY9388354/refund \
-H "Content-Type:application/json"  \
-H "Authorization: Bearer EBKoE2M-mctKH3OnORRVzUlxiromWnU5Mgz2PUZntQmt"  \
-d '{
  "amount" : {
    "currency" : "USD",
    "total" : "110.54"
  }
}'


RESPONSE: Returns a REFUND object with details about a refund and whether the refund was successful.

{
  "id": "0P209507D6694645N",
  "create_time": "2013-05-06T22:11:51Z",
  "update_time": "2013-05-06T22:11:51Z",
  "state": "completed",
  "amount": {
    "total": "110.54",
    "currency": "USD"
  },
  "capture_id": "8F148933LY9388354",
  "parent_payment": "PAY-8PT597110X687430LKGECATA",
  "links": [
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/refund/0P209507D6694645N",
      "rel": "self",
      "method": "GET"
    },
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/payment/PAY-8PT597110X687430LKGECATA",
      "rel": "parent_payment",
      "method": "GET"
    },
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/capture/8F148933LY9388354",
      "rel": "capture",
      "method": "GET"
    }
  ]
}

**************************************************************/

func (self *CaptureObject) do_refund(ref_req interface{}) (*RefundObject, error) {
	var ref_resp = new(RefundObject)
	var err = self.captures.connection.make_request("POST",
		"payments/capture/"+self.Id+"/refund",
		ref_req, "", ref_resp, false)
	if err != nil {
		return nil, err
	}
	return ref_resp, nil
}

// the Amount must include the PayPal fee paid by the Payee
func (self *CaptureObject) Refund(amt Amount) (*RefundObject, error) {
	return self.do_refund(&RefundObject{_trans: _trans{Amount: amt}})
}

func (self *CaptureObject) FullRefund() (*RefundObject, error) {
	return self.do_refund(&RefundObject{_trans: _trans{Amount: self.Amount}})
}
