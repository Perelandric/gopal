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
	pymt.Intent = intent.Sale
	pymt.Transactions = make([]*CreditCardTransaction, 0)

	pymt.Payer.PaymentMethod = PaymentMethod.CreditCard
	pymt.Payer.PayerInfo = nil

	return &pymt
}

/*
// TODO: Add `billing_agreement_tokens`, `payment_instruction`
*/

/*
@struct
*/
type CreditCardPayment struct {
	*connection
	Intent              intentEnum             `json:"intent,omitempty"`
	State               StateEnum              `json:"state,omitempty"`
	Id                  string                 `json:"id,omitempty"`
	FailureReason       FailureReasonEnum      `json:"failure_reason,omitempty"`
	CreateTime          dateTime               `json:"create_time,omitempty"`
	UpdateTime          dateTime               `json:"update_time,omitempty"`
	Links               links                  `json:"links,omitempty"`
	Transactions        CreditCardTransactions `json:"transactions,omitempty"`
	ExperienceProfileId string                 `json:"experience_profile_id"`
	Payer               creditCardPayer        `json:"payer,omitempty"`
	RedirectUrls        Redirects              `json:"redirect_urls,omitempty"`
	*payment_error
}

func (cp *CreditCardPayment) AddTransaction(
	c CurrencyTypeEnum,
	shp *ShippingAddress,
) *CreditCardTransaction {

	var t CreditCardTransaction

	t.Amount = amount{}
	t.ItemList = &creditCardItemList{}

	t.Amount.Currency = c
	t.Amount.Total = 0

	t.ItemList.Items = make([]*CreditCardItem, 0, 1)
	t.ItemList.ShippingAddress = shp

	cp.Transactions = append(cp.Transactions, &t)

	return &t
}

func (self *CreditCardPayment) calculateToAuthorize() {
	for _, t := range self.Transactions {
		t.calculateToAuthorize()
	}
}

// TODO: Needs to validate some sub-properties that are valid only when
// Payer.PaymentMethod is "paypal"
func (self *CreditCardPayment) validate() (err error) {
	if len(self.Transactions) == 0 {
		return fmt.Errorf("A Payment needs at least one Transaction")
	}

	for _, t := range self.Transactions {
		if err = t.validate(); err != nil {
			return err
		}
	}

	// TODO: More validation

	return self.Payer.validate()
}

// TODO: Finish implementation
func (self *CreditCardPayment) AddFundingInstrument(instrs ...fundingInstrument) error {
	for _, instr := range instrs {
		//		err := instr.validate()
		//		if err != nil {
		//			return err
		//		}
		self.Payer.FundingInstruments = append(
			self.Payer.FundingInstruments, &instr,
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
		switch pymt.State {
		case State.Created:
			// Set url to redirect to PayPal site to begin approval process
			to, _ = pymt.Links.get(relType.ApprovalUrl)
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
	for _, trans := range self.Transactions {
		for _, related_resource := range trans.RelatedResources {
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
type CreditCardTransaction struct {
	ItemList         *creditCardItemList `json:"item_list,omitempty"`
	Amount           amount              `json:"amount"`
	RelatedResources relatedResources    `json:"related_resources,omitempty"`
	Description      string              `json:"description,omitempty"`
	PaymentOptions   paymentOptions      `json:"payment_options,omitempty"`
	InvoiceNumber    string              `json:"invoice_number,omitempty"`
	Custom           string              `json:"custom,omitempty"`
	SoftDescriptor   string              `json:"soft_descriptor,omitempty"`
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

	item.Currency = t.Amount.Currency

	t.ItemList.Items =
		append(t.ItemList.Items, item)
	return nil
}

func (self *CreditCardTransaction) validate() (err error) {
	if err = self.ItemList.validate(); err != nil {
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

	return self.Amount.validate()
}

func (self *CreditCardTransaction) calculateToAuthorize() {
	// Calculate totals from itemList
	for _, item := range self.ItemList.Items {
		self.Amount.Details.Subtotal = roundTwoDecimalPlaces(
			self.Amount.Details.Subtotal + (item.Price * float64(item.Quantity)))
	}

	// Set Total, which is sum of Details
	self.Amount.Total = roundTwoDecimalPlaces(
		self.Amount.Details.Subtotal +
			self.Amount.Details.Tax +
			self.Amount.Details.Shipping +
			self.Amount.Details.Insurance -
			self.Amount.Details.ShippingDiscount)
}

/*
@struct
*/
type creditCardItemList struct {
	Items           []*CreditCardItem `json:"items,omitempty"`
	ShippingAddress *ShippingAddress  `json:"shipping_address,omitempty"`
}

func (self *creditCardItemList) validate() (err error) {
	if self == nil {
		return nil
	}
	if len(self.Items) == 0 {
		return fmt.Errorf("Transaction item list must have at least one Item")
	}

	for _, item := range self.Items {
		if err = item.validate(); err != nil {
			return err
		}
	}
	return self.ShippingAddress.validate()
}

/*
@struct
*/
// Source of the funds for this payment represented by a credit card.
type creditCardPayer struct {
	// Must be PaymentMethod.CreditCard
	PaymentMethod PaymentMethodEnum `json:"payment_method,omitempty"`

	FundingInstruments fundingInstruments `json:"funding_instruments,omitempty"`

	PayerInfo *PayerInfo `json:"payer_info,omitempty"`
}

func (self *creditCardPayer) validate() error {
	err := self.PaymentMethod.validate()
	if err != nil {
		return err
	}
	return nil
}

/*
@struct
*/
type PayerInfo struct {
	// Email address representing the payer. 127 characters max.
	Email string `json:"email,omitempty"`

	// Salutation of the payer.
	Salutation string `json:"salutation,omitempty"`

	// Suffix of the payer.
	Suffix string `json:"suffix,omitempty"`

	// Two-letter registered country code of the payer to identify the buyer country.
	CountryCode CountryCodeEnum `json:"country_code,omitempty"`

	// Phone number representing the payer. 20 characters max.
	Phone string `json:"phone,omitempty"`

	// First name of the payer. Value assigned by PayPal.
	FirstName string `json:"first_name,omitempty"`

	// Middle name of the payer. Value assigned by PayPal.
	MiddleName string `json:"middle_name,omitempty"`

	// Last name of the payer. Value assigned by PayPal.
	LastName string `json:"last_name,omitempty"`

	// PayPal assigned Payer ID. Value assigned by PayPal.
	PayerId string `json:"payer_id,omitempty"`

	// Shipping address of payer PayPal account. Value assigned by PayPal.
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

func (self *PayerInfo) validate() error {
	// TODO: Implement
	if self == nil {
		return nil
	}

	return nil
}
