package gopal

/*************************************************************

	SALE TRANSACTIONS:  https://api.paypal.com/v1/payments/sale

	Get and refund completed payments (sale transactions).

	  To get details about completed payments (sale transaction)
	created by a payment request or to refund a direct sale transaction,
	PayPal provides the /sale resource and related sub-resources. You can
	find the sale transactions in the payment resource within related_resources.

**************************************************************/

type Sales struct {
    pathGroup *PathGroup
}


type SaleObject struct {
    _trans
    State       State  `json:"state,omitempty"` // TODO: Limit to allowed values

// TODO: Verify that `sale_id` shouldn't be there and can be removed
//	Sale_id     string `json:"sale_id,omitempty"`

    Description string `json:"description,omitempty"`

	*identity_error // TODO: Is this right, or is there a special error object like `payments` has?
	sales *Sales
}


/*************************************************************

	LOOK UP A SALE
	GET https://api.paypal.com/v1/payments/sale/{id}

Use this call to get details about a sale transaction.


REQUEST: (No object to send)

curl -v -X GET https://api.sandbox.paypal.com/v1/payments/sale/36C38912MN9658832 \
-H "Content-Type:application/json" \
-H "Authorization:Bearer EIpHzmJ8BGVEBw5KxBQZfibqL-3xfc3ECveRDjHX0_Ur"


RESPONSE: Returns a SALE object.

{
  "id": "36C38912MN9658832",
  "create_time": "2013-02-19T22:01:53Z",
  "update_time": "2013-02-19T22:01:55Z",
  "state": "completed",
  "amount": {
    "total": "7.47",
    "currency": "USD"
  },
  "parent_payment": "PAY-5YK922393D847794YKER7MUI",
  "links": [
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/sale/36C38912MN9658832",
      "rel": "self",
      "method": "GET"
    },
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/sale/36C38912MN9658832/refund",
      "rel": "refund",
      "method": "POST"
    },
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/payment/PAY-5YK922393D847794YKER7MUI",
      "rel": "parent_payment",
      "method": "GET"
    }
  ]
}

**************************************************************/


func (self *Sales) Get(sale_id string) (*SaleObject, error) {
	var sale = new(SaleObject)
	var err = self.pathGroup.connection.make_request("GET", "payments/sale/" + sale_id, nil, "", sale, false)
	if err != nil {
		return nil, err
	}
	sale.sales = self
	return sale, nil
}

func (self *SaleObject) GetParentPayment() (*PaymentObject, error) {
	return self.sales.pathGroup.Payments.Get(self.Parent_payment)
}



/*************************************************************

	REFUND A SALE
	POST https://api.paypal.com/v1/payments/sale/{sale_id}/refund

Use this call to refund a completed payment.
Provide the sale_id in the URI and an empty JSON payload for a full refund.
For partial refunds, you can include an amount.


REQUEST: Sends an AMOUNT object

curl -v https://api.sandbox.paypal.com/v1/payments/sale/32G235960X3106808/refund \
-H "Content-Type:application/json"  \
-H "Authorization: Bearer EBKoE2M-mctKH3OnORRVzUlxiromWnU5Mgz2PUZntQmt"  \
-d '{"amount":{"total":"2.34","currency":"USD"}}'


RESPONSE: Returns a REFUND object with details about a refund and whether the refund was successful.

{
  "id": "4CF18861HF410323U",
  "create_time": "2013-01-31T04:13:34Z",
  "update_time": "2013-01-31T04:13:36Z",
  "state": "completed",
  "amount": {
    "total": "2.34",
    "currency": "USD"
  },
  "sale_id": "2MU78835H4515710F",
  "parent_payment": "PAY-46E69296BH2194803KEE662Y",
  "links": [
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/refund/4CF18861HF410323U",
      "rel": "self",
      "method": "GET"
    },
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/payment/PAY-46E69296BH2194803KEE662Y",
      "rel": "parent_payment",
      "method": "GET"
    },
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/sale/2MU78835H4515710F",
      "rel": "sale",
      "method": "GET"
    }
  ]
}

**************************************************************/

func (self *SaleObject) do_refund(ref_req interface{}) (*RefundObject, error) {
	var ref_resp = new(RefundObject)
	var err = self.sales.pathGroup.connection.make_request("POST",
															"payments/sale/" + self.Id + "/refund",
															ref_req, "", ref_resp, false)
	if err != nil {
		return nil, err
	}
	return ref_resp, nil
}

func (self *SaleObject) Refund(amt Amount) (*RefundObject, error) {
	return self.do_refund(&RefundObject{_trans:_trans{Amount:amt}})
}

func (self *SaleObject) FullRefund() (*RefundObject, error) {
	return self.do_refund(`{}`)
}



