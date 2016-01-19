package gopal

import "path"

/*************************************************************

	SALE TRANSACTIONS:  https://api.paypal.com/v1/payments/sale

	Get and refund completed payments (sale transactions).

	  To get details about completed payments (sale transaction)
	created by a payment request or to refund a direct sale transaction,
	PayPal provides the /sale resource and related sub-resources. You can
	find the sale transactions in the payment resource within related_resources.

**************************************************************/

// State values are: pending; completed; refunded; partially_refunded
type Sale struct {
	_trans

	// TODO: Verify that `sale_id` shouldn't be there and can be removed
	//	Sale_id     string `json:"sale_id,omitempty"`

	Description string `json:"description,omitempty"`

	// Specifies payment mode of the transaction. Only supported when the
	// `payment_method` is set to `paypal`. Assigned by PayPal.
	PaymentMode paymentModeEnum `json:"payment_mode"`

	// Reason the transaction is in pending state. Only supported when the
	// `payment_method` is set to `paypal`
	PendingReason pendingReasonEnum `json:"pending_reason"`

	// Expected clearing time for eCheck transactions. Only supported when the
	// payment_method is set to paypal. Assigned by PayPal.
	ClearingTime string `json:"clearing_time"`

	// Transaction fee charged by PayPal for this transaction.
	TransactionFee currency `json:"transaction_fee"`

	// Net amount the merchant receives for this transaction in their receivable
	// currency. Returned only in cross-currency use cases where a merchant bills
	// a buyer in a non-primary currency for that buyer.
	ReceivableAmount currency `json:"receivable_amount"`

	// Exchange rate applied for this transaction. Returned only in cross-currency
	// use cases where a merchant bills a buyer in a non-primary currency for that
	// buyer.
	ExchangeRate string `json:"exchange_rate"`

	// Fraud Management Filter (FMF) details applied for the payment that could
	// result in accept, deny, or pending action. Returned in a payment response
	// only if the merchant has enabled FMF in the profile settings and one of the
	// fraud filters was triggered based on those settings. See "Fraud Management
	// Filters Summary" for more information.
	FmfDetails fmfDetails `json:"fmf_details"`

	// Receipt ID is a 16-digit payment identification number returned for guest
	// users to identify the payment.
	ReceiptId string `json:"receipt_id"`

	// Reason code for the transaction state being Pending or Reversed. Only
	// supported when the `payment_method` is set to `paypal`.
	ReasonCode reasonCodeEnum `json:"reason_code"`

	// The level of seller protection in force for the transaction. Only supported
	// when the `payment_method` is set to `paypal`.
	ProtectionEligibility protectionEligEnum `json:"protection_eligibility"`

	// The kind of seller protection in force for the transaction. This property
	// is returned only when the protection_eligibility property is set to
	// `ELIGIBLE` or `PARTIALLY_ELIGIBLE`. Only supported when the `payment_method`
	// is set to paypal. One or both of the allowed values can be returned.
	ProtectionEligibilityType protectionEligTypeEnum `json:"protection_eligibility_type"`
}

// Implement the Resource interface

func (self *Sale) getPath() string {
	return path.Join(_salePath, self.Id)
}

// Implement `refundable` interface

func (self *Sale) getRefundPath() string {
	return path.Join(_salePath, self.Id, _refund)
}

/*************************************************************

	LOOK UP A SALE
	GET https://api.paypal.com/v1/payments/sale/{id}

Use this call to get details about a sale transaction.

**************************************************************/

func (self *connection) FetchSale(sale_id string) (*Sale, error) {
	var sale = &Sale{}
	sale.connection = self

	if err := self.send(&request{
		method:   method.get,
		path:     path.Join(_salePath, sale_id),
		body:     nil,
		response: sale,
	}); err != nil {
		return nil, err
	}

	return sale, nil
}

/*************************************************************

	REFUND A SALE
	POST https://api.paypal.com/v1/payments/sale/{sale_id}/refund

Use this call to refund a completed payment.
Provide the sale_id in the URI and an empty JSON payload for a full refund.
For partial refunds, you can include an amount.

**************************************************************/

func (self *Sale) Refund(amt float64) (*Refund, error) {
	return self.doRefund(self, amt)
}

func (self *Sale) FullRefund() (*Refund, error) {
	return self.doRefund(self, self.totalAsFloat())
}
