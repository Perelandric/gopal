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

// Implement the resource interface

func (self *Sale) getPath() string {
	return path.Join(_salePath, self._shared.private.Id)
}

// Implement `refundable` interface

func (self *Sale) getRefundPath() string {
	return path.Join(_salePath, self._shared.private.Id, _refund)
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
	return self.doRefund(self, self.private.Amount.private.Total)
}
