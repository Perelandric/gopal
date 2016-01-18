package gopal

import (
	"fmt"
	"path"
)

/*************************************************************

	REFUNDS:  https://api.paypal.com/v1/refund

	  Use the /refund resource to look up details of a specific
	refund on direct and captured payments.

	See "Refund a sale" in the API reference and "Refund a completed
	payment (sale)" for more information about refunding payments.

	https://developer.paypal.com/webapps/developer/docs/api/#refund-a-sale
	https://developer.paypal.com/webapps/developer/docs/integration/direct/refund-payment/

**************************************************************/

// State items are: pending; completed; failed
type Refund struct {
	_trans

	// ID of the Sale transaction being refunded. One among sale_id or capture_id
	// will be returned based on the resource used to initiate refund.
	SaleId string `json:"sale_id,omitempty"`

	// ID of the sale transaction being refunded.
	CaptureId string `json:"capture_id,omitempty"`
}

// Implement the transactionable interface

func (self *Refund) getPath() string {
	return path.Join("payments/refund", self.Id)
}

// General purpose function for performing a refund.

func (self *connection) doRefund(obj refundable, amt float64) (*Refund, error) {
	if amt <= 0 {
		return nil, fmt.Errorf("Refund must be greater than 0. Found: %f\n", amt)
	}

	ref := &Refund{}

	ref.Amount.Currency = obj.getCurrencyType()
	if err := ref.Amount.setTotal(amt); err != nil {
		return nil, err
	}

	if err := self.send(&request{
		method:   post,
		path:     obj.getRefundPath(),
		body:     ref,
		response: ref,
	}); err != nil {
		return nil, err
	}
	return ref, nil
}

/*************************************************************

	LOOK UP A REFUND
	GET https://api.paypal.com/v1/refund/{id}

Use this call to get details about a specific refund.

To get a list of your refunds, you can first get a "list of payments".
Within the list, you can see the state of the sale object as refunded
and a refund object with the state of completed.

https://developer.paypal.com/webapps/developer/docs/api/#list-payment-resources

**************************************************************/

func (self *connection) FetchRefund(refund_id string) (*Refund, error) {
	var refund = &Refund{}
	refund.connection = self

	if err := self.send(&request{
		method:   get,
		path:     path.Join("payments/refund", refund_id),
		body:     nil,
		response: refund,
	}); err != nil {
		return nil, err
	}

	return refund, nil
}
