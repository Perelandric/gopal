package gopal

import (
	"fmt"
	"log"
	"net/url"
	"path"
)

//go:generate Golific $GOFILE

/*
This file contains types that are specific to creating Paypal payments. Several
types are used for both Paypal and credit cards, yet have restrictions for one
or the other.
*/

func (self *connection) NewPaypalPayment(
	urls Redirects, info *PaypalPayerInfo) (*PaypalPayment, error) {

	var pymt = PaypalPayment{
		connection: self,
	}
	pymt.private.Intent = intent.Sale
	pymt.private.Transactions = make([]*PaypalTransaction, 0)
	pymt.private.RedirectUrls = urls

	pymt.private.Payer.private.PaymentMethod = PaymentMethod.PayPal
	pymt.private.Payer.private.PaypalPayerInfo = info

	if err := urls.validate(); err != nil {
		return nil, err
	}
	if err := info.validate(); err != nil {
		return nil, err
	}

	return &pymt, nil
}

type Redirects struct {
	Return string `json:"return_url,omitempty"`
	Cancel string `json:"cancel_url,omitempty"`
}

func (self *Redirects) validate() error {
	for _, s := range [2]string{self.Return, self.Cancel} {
		u, err := url.Parse(s)
		if err != nil {
			return err
		}

		if len(u.Scheme) == 0 {
			return fmt.Errorf("URL Scheme is required. Found %q\n", s)
		}
		if len(u.Host) == 0 {
			return fmt.Errorf("URL Host is required. Found %q\n", s)
		}
	}
	return nil
}

type PaypalTransactions []*PaypalTransaction

/*
// TODO: Add `billing_agreement_tokens`, `payment_instruction`
@struct PaypalPayment
  *connection
  Intent							intentEnum 		`json:"intent,omitempty"` --read
  State 							*StateEnum `json:"state,omitempty"` --read
  Id 									string `json:"id,omitempty"` --read
  FailureReason				*FailureReasonEnum `json:"failure_reason,omitempty"` --read
  CreateTime 					dateTime `json:"create_time,omitempty"` --read
  UpdateTime 					dateTime `json:"update_time,omitempty"` --read
  Links 							links `json:"links,omitempty"` --read
  Transactions				PaypalTransactions 	`json:"transactions,omitempty"` --read
	Payer 							paypalPayer `json:"payer,omitempty"` --read
	RedirectUrls				Redirects `json:"redirect_urls,omitempty"` --read
	ExperienceProfileId string `json:"experience_profile_id,omitempty"` --read --write
  *payment_error
*/

func (self *PaypalPayment) AddTransaction(
	c CurrencyTypeEnum, shp *ShippingAddress) *PaypalTransaction {

	var t PaypalTransaction
	t.private.Amount = amount{Details: &details{}}
	t.private.ItemList = &paypalItemList{}

	t.private.Amount.private.Currency = c
	t.private.Amount.private.Total = 0

	t.private.ItemList.private.Items = make([]*PaypalItem, 0, 1)
	t.private.ItemList.private.ShippingAddress = shp

	self.private.Transactions = append(self.private.Transactions, &t)

	return &t
}

func (self *PaypalPayment) calculateToAuthorize() {
	for _, t := range self.private.Transactions {
		t.calculateToAuthorize()
	}
}

// TODO: Needs to validate some sub-properties that are valid only when
// Payer.PaymentMethod is "paypal"
func (self *PaypalPayment) validate() (err error) {
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

// TODO: Should send a query string parameter `token=[some token]`
func (self *PaypalPayment) Authorize() (to string, code int, err error) {
	if err = self.validate(); err != nil {
		return "", 0, err
	}

	self.calculateToAuthorize()

	// Create Totals
	var pymt PaypalPayment

	err = self.send(&request{
		method:   method.Post,
		path:     _paymentsPath,
		body:     self,
		response: &pymt,
	})

	if err == nil {
		if pymt.private.State == nil {
			err = UnexpectedResponse

		} else {
			switch *pymt.private.State {
			case State.Created:
				// Set url to redirect to PayPal site to begin approval process
				to, _ = pymt.private.Links.get(relType.ApprovalUrl)
				code = 303
			default:
				// otherwise cancel the payment and return an error
				err = UnexpectedResponse
			}
		}
	}

	return to, code, err
}

// Used to execute the approved payment.
type paymentExecution struct {
	// The ID of the Payer, passed in the return_url by PayPal.
	PayerID string `json:"payer_id,omitempty"`

	// Transactional details if updating a payment. Note that this instance of
	// the transactions object accepts only the Amount object.
	Transactions PaypalTransactions `json:"transactions,omitempty"`
}

func (self *connection) Execute(u *url.URL) error {
	var query = u.Query()

	var payerid = query.Get("PayerID")
	if payerid == "" {
		return fmt.Errorf("PayerID is missing")
	}

	var pymtid = query.Get("paymentId")
	if pymtid == "" {
		return fmt.Errorf("paymentId is missing")
	}

	var pymt PaypalPayment

	if err := self.send(&request{
		method:   method.Post,
		path:     path.Join(_paymentsPath, pymtid, _execute),
		body:     &paymentExecution{PayerID: payerid},
		response: &pymt,
	}); err != nil {
		return err
	}

	if pymt.private.State == nil || *pymt.private.State != State.Approved {
		return fmt.Errorf(
			"Payment with ID %q for payer %q was not approved", pymtid, payerid)
	}

	return nil
}

// TODO: I'm returning a list of Sale objects because every payment can have multiple transactions.
//		I need to find out why a payment can have multiple transactions, and see if I should eliminate that in the API
func (self *PaypalPayment) FetchSale() []*Sale {
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

/*
@struct PaypalTransaction
  ItemList 				*paypalItemList `json:"item_list,omitempty"` --read
  Amount 					amount `json:"amount"` --read
  RelatedResources relatedResources `json:"related_resources,omitempty"` --read
  Description 		string `json:"description,omitempty"` --read --write
  PaymentOptions  *paymentOptions `json:"payment_options,omitempty"` --read --write
	InvoiceNumber 	string `json:"invoice_number,omitempty"` --read --write
	Custom 					string `json:"custom,omitempty"` --read --write
	SoftDescriptor 	string `json:"soft_descriptor,omitempty"` --read --write
*/

// Prices are assumed to use the CurrencyType passed to NewTransaction.
func (t *PaypalTransaction) AddItem(item *PaypalItem) (err error) {
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

func (self *PaypalTransaction) validate() (err error) {
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

func (self *PaypalTransaction) calculateToAuthorize() {
	// Calculate totals from itemList
	for _, item := range self.private.ItemList.private.Items {
		self.private.Amount.Details.private.Subtotal = roundTwoDecimalPlaces(
			self.private.Amount.Details.private.Subtotal + (item.Price * float64(item.Quantity)))

		self.private.Amount.Details.private.Tax = roundTwoDecimalPlaces(
			self.private.Amount.Details.private.Tax + (item.Tax * float64(item.Quantity)))
	}

	// Set Total, which is sum of Details
	self.private.Amount.private.Total = roundTwoDecimalPlaces(
		self.private.Amount.Details.private.Subtotal +
			self.private.Amount.Details.private.Tax +
			self.private.Amount.Details.Shipping +
			self.private.Amount.Details.Insurance -
			self.private.Amount.Details.ShippingDiscount)
}

type PaypalItems []*PaypalItem

/*
@struct paypalItemList
	Items           PaypalItems          	 `json:"items,omitempty"` --read
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"` --read
*/

func (self *paypalItemList) validate() (err error) {
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
@struct PaypalItem
	Currency 		CurrencyTypeEnum 	`json:"currency"` --read
	Quantity 		int64 			`json:"quantity"` --read --write
	Name 				string 			`json:"name"` --read --write
	Price 			float64 		`json:"price,string"` --read --write
	Tax 				float64 		`json:"tax,string,omitempty"` --read --write
	Sku 				string 			`json:"sku,omitempty"` --read --write
	Description string 			`json:"description,omitempty"` --read --write
*/

func (self *PaypalItem) validate() (err error) {
	if self.Tax < 0 { // TODO: No other validation here???
		return fmt.Errorf("%q must not be a negative number", "Item.Tax")
	}
	if err = checkStr("Item.Name", &self.Name, 127, true, true); err != nil {
		return err
	}
	if err = checkStr("Item.Sku", &self.Sku, 50, false, true); err != nil {
		return err
	}
	_ = checkStr("Item.Description", &self.Description, 127, false, false)

	if err = checkFloat7_10("Item.Price", &self.Price, true); err != nil {
		return err
	}

	return checkInt10("Item.Quantity", self.Quantity, true)
}

/*
// Source of the funds for this payment represented by a PayPal account.
@struct paypalPayer
  // Must be PaymentMethod.Paypal
	PaymentMethod      PaymentMethodEnum `json:"payment_method,omitempty"`

  // Status of the payer’s PayPal account. Allowed values: VERIFIED or UNVERIFIED.
	Status 						 *payerStatusEnum `json:"status,omitempty"` --read

	PaypalPayerInfo    *PaypalPayerInfo `json:"payer_info,omitempty"` --read
*/

func (self *paypalPayer) validate() error {
	err := self.private.PaymentMethod.validate()
	if err != nil {
		return err
	}
	return nil
}

/*
// This object is pre-filled by PayPal when the payment_method is paypal.
@struct PaypalPayerInfo
  // Payer’s tax ID type. Allowed values: BR_CPF or BR_C`NPJ. Only supported when
  // the payment_method is set to paypal.
  TaxIdType TaxIdTypeEnum `json:"tax_id_type,omitempty"` --read --write

  // Payer’s tax ID. Only supported when the payment_method is set to paypal.
  TaxId string `json:"tax_id,omitempty"` --read --write

	// Email address representing the payer. 127 characters max.
	Email string `json:"email,omitempty"` --read

	// Salutation of the payer.
	Salutation string `json:"salutation,omitempty"` --read

	// Suffix of the payer.
	Suffix string `json:"suffix,omitempty"` --read

	// Two-letter registered country code of the payer to identify the buyer country.
	CountryCode CountryCodeEnum `json:"country_code,omitempty"` --read

	// Phone number representing the payer. 20 characters max.
	Phone string `json:"phone,omitempty"` --read

	// First name of the payer. Value assigned by PayPal.
	FirstName string `json:"first_name,omitempty"` --read

	// Middle name of the payer. Value assigned by PayPal.
	MiddleName string `json:"middle_name,omitempty"` --read

	// Last name of the payer. Value assigned by PayPal.
	LastName string `json:"last_name,omitempty"` --read

	// PayPal assigned Payer ID. Value assigned by PayPal.
	PayerId string `json:"payer_id,omitempty"` --read

	// Shipping address of payer PayPal account. Value assigned by PayPal.
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"` --read
*/

func (self *PaypalPayerInfo) validate() error {
	// TODO: Implement
	if self == nil {
		return nil
	}

	return nil
}
