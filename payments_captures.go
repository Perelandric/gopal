package gopal

import "path"

//go:generate Golific $GOFILE

/*************************************************************

    CAPTURES:  https://api.paypal.com/v1/payments/payment/capture

	The /capture resource and sub-resources allow you to look up and
	refund captured payments.

**************************************************************/

/*
@struct
*/
// State values are: pending, completed, refunded, partially_refunded
type Capture struct {
	_shared
	Amount         amount   `json:"amount"`
	TransactionFee currency `json:"transaction_fee"`
	IsFinalCapture bool     `json:"is_final_capture,omitempty"`
}

// Implement the resource interface

func (self *Capture) getPath() string {
	return path.Join(_capturePath, self._shared.Id)
}
func (c *Capture) GetAmount() amount {
	return c.Amount
}

// Implement `refundable` interface

func (self *Capture) getRefundPath() string {
	return path.Join(_capturePath, self._shared.Id, _refund)
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
		method:   method.Get,
		path:     path.Join(_capturePath, capt_id),
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
	return self.doRefund(self, self.Amount.Total)
}
