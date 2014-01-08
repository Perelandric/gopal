package gopal

/*************************************************************

	REFUNDS:  https://api.paypal.com/v1/refund

	  Use the /refund resource to look up details of a specific
	refund on direct and captured payments.

	See "Refund a sale" in the API reference and "Refund a completed
	payment (sale)" for more information about refunding payments.

	https://developer.paypal.com/webapps/developer/docs/api/#refund-a-sale
	https://developer.paypal.com/webapps/developer/docs/integration/direct/refund-payment/

**************************************************************/

type Refunds struct {
	connection *Connection
}

type RefundObject struct {
	_trans
	State   State  `json:"state,omitempty"` // TODO: Limit to allowed values
	Sale_id string `json:"sale_id,omitempty"`

	*identity_error
	refunds *Refunds
}

/*************************************************************

	LOOK UP A REFUND
	GET https://api.paypal.com/v1/refund/{id}

Use this call to get details about a specific refund.

To get a list of your refunds, you can first get a "list of payments".
Within the list, you can see the state of the sale object as refunded
and a refund object with the state of completed.

https://developer.paypal.com/webapps/developer/docs/api/#list-payment-resources


REQUEST: Pass the refund ID. (No object to send)

curl -v -X GET https://api.sandbox.paypal.com/v1/refund/4GU360220B627614A \
-H "Content-Type:application/json" \
-H "Authorization:Bearer EMxItHE7Zl4cMdkvMg-f7c63GQgYZU8FjyPWKQlpsqQP"


RESPONSE: Returns a REFUND object with details about a refund and whether the refund
			was successful.

{
  "id": "4GU360220B627614A",
  "create_time": "2013-01-01T02:00:00Z",
  "update_time": "2013-01-01T03:00:02Z",
  "state": "completed",
  "amount": {
    "total": "2.34",
    "currency": "USD"
  },
  "sale_id": "36C38912MN9658832",
  "parent_payment": "PAY-5YK922393D847794YKER7MUI",
  "links": [
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/refund/4GU360220B627614A",
      "rel": "self",
      "method": "GET"
    },
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/payment/PAY-5YK922393D847794YKER7MUI",
      "rel": "parent_payment",
      "method": "GET"
    },
    {
      "href": "https://api.sandbox.paypal.com/v1/payments/sale/36C38912MN9658832",
      "rel": "sale",
      "method": "GET"
    }
  ]
}

**************************************************************/

func (self *Refunds) Get(refund_id string) (*RefundObject, error) {
	var refund = new(RefundObject)
	var err = self.connection.make_request("GET", "payments/refund/"+refund_id, nil, "", refund, false)
	if err != nil {
		return nil, err
	}
	refund.refunds = self
	return refund, nil
}

func (self *RefundObject) GetParentPayment() (*PaymentObject, error) {
	return self.refunds.connection.Payments.Get(self.Parent_payment)
}
