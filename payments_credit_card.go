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

func (c *connection) NewCreditCardPayment() *CreditCardPayment {
	var pymt = CreditCardPayment{
		connection: c,
	}
	pymt.private.Intent = intent.Sale
	pymt.private.Transactions = make([]*CreditCardTransaction, 0)

	pymt.private.Payer.private.PaymentMethod = PaymentMethod.CreditCard
	pymt.private.Payer.private.PayerInfo = nil

	return &pymt
}

/*
// TODO: Add `billing_agreement_tokens`, `payment_instruction`
*/

/*
@struct
*/
type __CreditCardPayment struct {
	*connection
	Intent              intentEnum             `gRead json:"intent,omitempty"`
	State               StateEnum              `gRead json:"state,omitempty"`
	Id                  string                 `gRead json:"id,omitempty"`
	FailureReason       FailureReasonEnum      `gRead json:"failure_reason,omitempty"`
	CreateTime          dateTime               `gRead json:"create_time,omitempty"`
	UpdateTime          dateTime               `gRead json:"update_time,omitempty"`
	Links               links                  `gRead json:"links,omitempty"`
	Transactions        CreditCardTransactions `gRead json:"transactions,omitempty"`
	ExperienceProfileId string                 `gRead gWrite json:"experience_profile_id"`
	Payer               creditCardPayer        `gRead json:"payer,omitempty"`
	RedirectUrls        Redirects              `gRead json:"redirect_urls,omitempty"`
	*payment_error
}

func (cp *CreditCardPayment) AddTransaction(
	c CurrencyTypeEnum,
	shp *ShippingAddress,
) *CreditCardTransaction {

	var t CreditCardTransaction

	t.private.Amount = amount{}
	t.private.ItemList = &creditCardItemList{}

	t.private.Amount.private.Currency = c
	t.private.Amount.private.Total = 0

	t.private.ItemList.private.Items = make([]*CreditCardItem, 0, 1)
	t.private.ItemList.private.ShippingAddress = shp

	cp.private.Transactions = append(cp.private.Transactions, &t)

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
@struct
*/
type __CreditCardTransaction struct {
	ItemList         *creditCardItemList `gRead json:"item_list,omitempty"`
	Amount           amount              `gRead json:"amount"`
	RelatedResources relatedResources    `gRead json:"related_resources,omitempty"`
	Description      string              `gRead gWrite json:"description,omitempty"`
	PaymentOptions   paymentOptions      `gRead gWrite json:"payment_options,omitempty"`
	InvoiceNumber    string              `gRead gWrite json:"invoice_number,omitempty"`
	Custom           string              `gRead gWrite json:"custom,omitempty"`
	SoftDescriptor   string              `gRead gWrite json:"soft_descriptor,omitempty"`
}

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

/*
@struct
*/
type __creditCardItemList struct {
	Items           []*CreditCardItem `gRead json:"items,omitempty"`
	ShippingAddress *ShippingAddress  `gRead json:"shipping_address,omitempty"`
}

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
@struct
*/
// Source of the funds for this payment represented by a credit card.
type __creditCardPayer struct {
	// Must be PaymentMethod.CreditCard
	PaymentMethod PaymentMethodEnum `gRead json:"payment_method,omitempty"`

	FundingInstruments fundingInstruments `gRead json:"funding_instruments,omitempty"`

	PayerInfo *PayerInfo `gRead json:"payer_info,omitempty"`
}

func (self *creditCardPayer) validate() error {
	err := self.private.PaymentMethod.validate()
	if err != nil {
		return err
	}
	return nil
}

/*
@struct
*/
type __PayerInfo struct {
	// Email address representing the payer. 127 characters max.
	Email string `gRead gWrite json:"email,omitempty"`

	// Salutation of the payer.
	Salutation string `gRead gWrite json:"salutation,omitempty"`

	// Suffix of the payer.
	Suffix string `gRead gWrite json:"suffix,omitempty"`

	// Two-letter registered country code of the payer to identify the buyer country.
	CountryCode CountryCodeEnum `gRead gWrite json:"country_code,omitempty"`

	// Phone number representing the payer. 20 characters max.
	Phone string `gRead gWrite json:"phone,omitempty"`

	// First name of the payer. Value assigned by PayPal.
	FirstName string `gRead gWrite json:"first_name,omitempty"`

	// Middle name of the payer. Value assigned by PayPal.
	MiddleName string `gRead gWrite json:"middle_name,omitempty"`

	// Last name of the payer. Value assigned by PayPal.
	LastName string `gRead gWrite json:"last_name,omitempty"`

	// PayPal assigned Payer ID. Value assigned by PayPal.
	PayerId string `gRead gWrite json:"payer_id,omitempty"`

	// Shipping address of payer PayPal account. Value assigned by PayPal.
	ShippingAddress *ShippingAddress `gRead gWrite json:"shipping_address,omitempty"`
}

func (self *PayerInfo) validate() error {
	// TODO: Implement
	if self == nil {
		return nil
	}

	return nil
}
