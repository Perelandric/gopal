package gopal

import "path"

//go:generate Golific $GOFILE

/*************************************************************

	SALE TRANSACTIONS:  https://api.paypal.com/v1/payments/sale

	Get and refund completed payments (sale transactions).

	  To get details about completed payments (sale transaction)
	created by a payment request or to refund a direct sale transaction,
	PayPal provides the /sale resource and related sub-resources. You can
	find the sale transactions in the payment resource within related_resources.

**************************************************************/

/*
@struct
*/
// State values are: pending; completed; refunded; partially_refunded
// TODO: PendingReason appears in the old docs under the general Sale object description
// but not under the lower "sale object" definition. The new docs have it
// marked as [DEPRECATED] in one area, but not another.
type Sale struct {
	_shared
	Amount                    amount                 `json:"amount"`
	Description               string                 `json:"description,omitempty"`
	TransactionFee            currency               `json:"transaction_fee"`
	ReceivableAmount          currency               `json:"receivable_amount"`
	PendingReason             pendingReasonEnum      `json:"pending_reason"`
	PaymentMode               paymentModeEnum        `json:"payment_mode"`
	ExchangeRate              string                 `json:"exchange_rate"`
	FmfDetails                fmfDetails             `json:"fmf_details"`
	ReceiptId                 string                 `json:"receipt_id"`
	ReasonCode                reasonCodeEnum         `json:"reason_code"`
	ProtectionEligibility     protectionEligEnum     `json:"protection_eligibility"`
	ProtectionEligibilityType protectionEligTypeEnum `json:"protection_eligibility_type"`
	ClearingTime              string                 `json:"clearing_time"`
	BillingAgreementId        string                 `json:"billing_agreement_id"`
}

// Implement the resource interface

func (self *Sale) getPath() string {
	return path.Join(_salePath, self._shared.Id)
}
func (s *Sale) GetAmount() amount {
	return s.Amount
}

// Implement `refundable` interface

func (self *Sale) getRefundPath() string {
	return path.Join(_salePath, self._shared.Id, _refund)
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
		method:   method.Get,
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
	return self.doRefund(self, self.Amount.Total)
}
