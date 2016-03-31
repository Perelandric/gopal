package gopal

import (
	"fmt"
	"log"
)

//go:generate Golific $GOFILE

/*
This file contains types that are specific to creating credit card payments.
Several types are used for both Paypal and credit cards, yet have restrictions
for one or the other.
*/

func (self *connection) NewCreditCardPayment() *CreditCardPayment {
	var pymt = CreditCardPayment{
		connection: self,
	}
	pymt.private.Intent = intent.Sale
	pymt.private.Transactions = make([]*CreditCardTransaction, 0)

	pymt.private.Payer.private.PaymentMethod = PaymentMethod.CreditCard
	pymt.private.Payer.private.PayerInfo = nil

	return &pymt
}

/*
// TODO: Add `billing_agreement_tokens`, `payment_instruction`
@struct CreditCardPayment
  *connection
  Intent				      intentEnum 		`json:"intent,omitempty"` --read
  State 							StateEnum `json:"state,omitempty"` --read
  Id 									string `json:"id,omitempty"` --read
  FailureReason				FailureReasonEnum `json:"failure_reason,omitempty"` --read
  CreateTime 					dateTime `json:"create_time,omitempty"` --read
  UpdateTime 					dateTime `json:"update_time,omitempty"` --read
  Links 							links `json:"links,omitempty"` --read
  Transactions	      CreditCardTransactions 	`json:"transactions,omitempty"` --read
	ExperienceProfileId string `json:"experience_profile_id"` --read --write
	Payer 							creditCardPayer `json:"payer,omitempty"` --read
	RedirectUrls				Redirects `json:"redirect_urls,omitempty"` --read
  *payment_error
*/

func (self *CreditCardPayment) AddTransaction(
	c CurrencyTypeEnum, shp *ShippingAddress) *CreditCardTransaction {

	var t CreditCardTransaction
	t.private.Amount = amount{}
	t.private.ItemList = &creditCardItemList{}

	t.private.Amount.private.Currency = c
	t.private.Amount.private.Total = 0

	t.private.ItemList.private.Items = make([]*CreditCardItem, 0, 1)
	t.private.ItemList.private.ShippingAddress = shp

	self.private.Transactions = append(self.private.Transactions, &t)

	return &t
}

func (self *CreditCardPayment) calculateToAuthorize() {
	for _, t := range self.private.Transactions {
		t.calculateToAuthorize()
	}
}

// TODO: Needs to validate some sub-properties that are valid only when
// Payer.PaymentMethod is "paypal"
func (self *CreditCardPayment) validate() (err error) {
	if len(self.private.Transactions) == 0 {
		return fmt.Errorf("A Payment needs at least one Transaction")
	}

	for _, t := range self.private.Transactions {
		if err = t.validate(); err != nil {
			return err
		}
	}

	// TODO: More validation

	return self.private.Payer.validate()
}

// TODO: Finish implementation
func (self *CreditCardPayment) AddFundingInstrument(instrs ...fundingInstrument) error {
	for _, instr := range instrs {
		//		err := instr.validate()
		//		if err != nil {
		//			return err
		//		}
		self.private.Payer.private.FundingInstruments = append(
			self.private.Payer.private.FundingInstruments, &instr,
		)
	}
	return nil
}

// TODO: Should send a query string parameter `token=[some token]`
// TODO: Right now this is set up for Paypal Payments. Need to change for CCards
func (self *CreditCardPayment) Authorize() (to string, code int, err error) {
	if err = self.validate(); err != nil {
		return "", 0, err
	}

	self.calculateToAuthorize()

	// Create Totals
	var pymt CreditCardPayment

	err = self.send(&request{
		method:   method.Post,
		path:     _paymentsPath,
		body:     self,
		response: &pymt,
	})

	if err == nil {
		switch pymt.private.State {
		case State.Created:
			// Set url to redirect to PayPal site to begin approval process
			to, _ = pymt.private.Links.get(relType.ApprovalUrl)
			code = 303
		default:
			// otherwise cancel the payment and return an error
			err = UnexpectedResponse
		}
	}

	return to, code, err
}

// TODO: I'm returning a list of Sale objects because every payment can have multiple transactions.
//		I need to find out why a payment can have multiple transactions, and see if I should eliminate that in the API
func (self *CreditCardPayment) FetchSale() []*Sale {
	var sales = []*Sale{}
	for _, trans := range self.private.Transactions {
		for _, related_resource := range trans.private.RelatedResources {
			if s, ok := related_resource.(*Sale); ok {
				sales = append(sales, s)
			}
		}
	}
	return sales
}

type CreditCardTransactions []*CreditCardTransaction

/*
@struct CreditCardTransaction
  ItemList 				*creditCardItemList `json:"item_list,omitempty"` --read
  Amount 					amount `json:"amount"` --read
  RelatedResources relatedResources `json:"related_resources,omitempty"` --read
  Description 		string `json:"description,omitempty"` --read --write
  PaymentOptions  paymentOptions `json:"payment_options,omitempty"` --read --write
	InvoiceNumber 	string `json:"invoice_number,omitempty"` --read --write
	Custom 					string `json:"custom,omitempty"` --read --write
	SoftDescriptor 	string `json:"soft_descriptor,omitempty"` --read --write
*/

// Prices are assumed to use the CurrencyType passed to NewTransaction.
func (t *CreditCardTransaction) AddItem(item *CreditCardItem) (err error) {
	if item == nil {
		log.Println("Transaction received unexpected nil Item")
		return nil
	}
	if err = item.validate(); err != nil {
		return err
	}

	item.private.Currency = t.private.Amount.private.Currency

	t.private.ItemList.private.Items =
		append(t.private.ItemList.private.Items, item)
	return nil
}

func (self *CreditCardTransaction) validate() (err error) {
	if err = self.private.ItemList.validate(); err != nil {
		return err
	}

	// TODO: More validation... check docs

	// These can be truncated with a warning if too long
	checkStr("Transaction.Description", &self.Description, 127, false, false)
	checkStr("Transaction.Custom", &self.Custom, 256, false, false)
	checkStr("Transaction.SoftDescriptor", &self.SoftDescriptor, 22, false, false)

	err = checkStr(
		"Transaction.InvoiceNumber", &self.InvoiceNumber, 256, false, true)
	if err != nil {
		return err
	}

	return self.private.Amount.validate()
}

func (self *CreditCardTransaction) calculateToAuthorize() {
	// Calculate totals from itemList
	for _, item := range self.private.ItemList.private.Items {
		self.private.Amount.Details.private.Subtotal = roundTwoDecimalPlaces(
			self.private.Amount.Details.private.Subtotal + (item.Price * float64(item.Quantity)))
	}

	// Set Total, which is sum of Details
	self.private.Amount.private.Total = roundTwoDecimalPlaces(
		self.private.Amount.Details.private.Subtotal +
			self.private.Amount.Details.private.Tax +
			self.private.Amount.Details.Shipping +
			self.private.Amount.Details.Insurance -
			self.private.Amount.Details.ShippingDiscount)
}

type CreditCardItems []*CreditCardItem

/*
@struct creditCardItemList
	Items           CreditCardItems          	 `json:"items,omitempty"` --read
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"` --read
*/

func (self *creditCardItemList) validate() (err error) {
	if self == nil {
		return nil
	}
	if len(self.private.Items) == 0 {
		return fmt.Errorf("Transaction item list must have at least one Item")
	}

	for _, item := range self.private.Items {
		if err = item.validate(); err != nil {
			return err
		}
	}
	return self.private.ShippingAddress.validate()
}

/*
@struct CreditCardItem
	Currency 		CurrencyTypeEnum 	`json:"currency"` --read
	Quantity 		int64 			`json:"quantity,string"` --read --write
	Name 				string 			`json:"name"` --read --write
	Price 			float64 		`json:"price,string"` --read --write
	Sku 				string 			`json:"sku,omitempty"` --read --write
*/

func (self *CreditCardItem) validate() (err error) {
	if err = checkStr("Item.Name", &self.Name, 127, true, true); err != nil {
		return err
	}
	if err = checkStr("Item.Sku", &self.Sku, 50, false, true); err != nil {
		return err
	}

	if err = checkFloat7_10("Item.Price", &self.Price, true); err != nil {
		return err
	}

	return checkInt10("Item.Quantity", self.Quantity, true)
}

/*
// Source of the funds for this payment represented by a credit card.
@struct creditCardPayer
  // Must be PaymentMethod.CreditCard
	PaymentMethod      PaymentMethodEnum `json:"payment_method,omitempty"` --read

	FundingInstruments fundingInstruments `json:"funding_instruments,omitempty"` --read

	PayerInfo          *PayerInfo `json:"payer_info,omitempty"` --read
*/

func (self *creditCardPayer) validate() error {
	err := self.private.PaymentMethod.validate()
	if err != nil {
		return err
	}
	return nil
}

/*
@struct PayerInfo
	// Email address representing the payer. 127 characters max.
	Email string `json:"email,omitempty"` --read --write

	// Salutation of the payer.
	Salutation string `json:"salutation,omitempty"` --read --write

	// Suffix of the payer.
	Suffix string `json:"suffix,omitempty"` --read --write

	// Two-letter registered country code of the payer to identify the buyer country.
	CountryCode CountryCodeEnum `json:"country_code,omitempty"` --read --write

	// Phone number representing the payer. 20 characters max.
	Phone string `json:"phone,omitempty"` --read --write

	// First name of the payer. Value assigned by PayPal.
	FirstName string `json:"first_name,omitempty"` --read --write

	// Middle name of the payer. Value assigned by PayPal.
	MiddleName string `json:"middle_name,omitempty"` --read --write

	// Last name of the payer. Value assigned by PayPal.
	LastName string `json:"last_name,omitempty"` --read --write

	// PayPal assigned Payer ID. Value assigned by PayPal.
	PayerId string `json:"payer_id,omitempty"` --read --write

	// Shipping address of payer PayPal account. Value assigned by PayPal.
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"` --read --write
*/

func (self *PayerInfo) validate() error {
	// TODO: Implement
	if self == nil {
		return nil
	}

	return nil
}
