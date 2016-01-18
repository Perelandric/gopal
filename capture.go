package gopal

import "path"

/*************************************************************

    CAPTURES:  https://api.paypal.com/v1/payments/payment/capture

	The /capture resource and sub-resources allow you to look up and
	refund captured payments.

**************************************************************/

// State values are: pending, completed, refunded, partially_refunded
type Capture struct {
	_trans

	// Transaction fee applicable for this payment.
	TransactionFee currency `json:"transaction_fee"`

	// If set to `true`, all remaining funds held by the authorization will be
	// released in the funding instrument. Default is `false`.
	IsFinalCapture bool `json:"is_final_capture,omitempty"`
}

// Implement the transactionable interface

func (self *Capture) getPath() string {
	return path.Join("payments/capture", self.Id)
}

// Implement `refundable` interface

func (self *Capture) getRefundPath() string {
	return path.Join("payments/capture", self.Id, "refund")
}

/*************************************************************

	LOOK UP A CAPTURED PAYMENT
	GET https://api.paypal.com/v1/payments/capture/{capture_id}

	Use this call to get details about a captured payment.

**************************************************************/

func (self *connection) FetchCapture(capt_id string) (*Capture, error) {
	var capt = &Capture{}
	capt.connection = self

	if err := self.send(&request{
		method:   get,
		path:     path.Join("payments/capture", capt_id),
		body:     nil,
		response: capt,
	}); err != nil {
		return nil, err
	}

	return capt, nil
}

/*************************************************************

	REFUND A CAPTURED PAYMENT
	POST https://api.paypal.com/v1/payments/capture/{capture_id}/refund

**************************************************************/

// the Amount must include the PayPal fee paid by the Payee
func (self *Capture) Refund(amt float64) (*Refund, error) {
	return self.doRefund(self, amt)
}

func (self *Capture) FullRefund() (*Refund, error) {
	return self.doRefund(self, self.totalAsFloat())
}
