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

func (c *connection) NewPaypalPayment(
	urls Redirects,
	info *PaypalPayerInfo,
) (*PaypalPayment, error) {

	var pymt = PaypalPayment{
		connection: c,
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

func (r *Redirects) validate() error {
	for _, s := range [2]string{r.Return, r.Cancel} {
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

// TODO: Add `billing_agreement_tokens`, `payment_instruction`

/*
@struct
*/
type __PaypalPayment struct {
	*connection
	Intent              intentEnum         `gRead json:"intent,omitempty"`
	State               StateEnum          `gRead json:"state,omitempty"`
	Id                  string             `gRead json:"id,omitempty"`
	FailureReason       FailureReasonEnum  `gRead json:"failure_reason,omitempty"`
	CreateTime          dateTime           `gRead json:"create_time,omitempty"`
	UpdateTime          dateTime           `gRead json:"update_time,omitempty"`
	Links               links              `gRead json:"links,omitempty"`
	Transactions        PaypalTransactions `gRead json:"transactions,omitempty"`
	Payer               paypalPayer        `gRead json:"payer"`
	RedirectUrls        Redirects          `gRead json:"redirect_urls,omitempty"`
	ExperienceProfileId string             `gRead gWrite json:"experience_profile_id,omitempty"`
	*payment_error
}

func (pp *PaypalPayment) AddTransaction(
	c CurrencyTypeEnum,
	shp *ShippingAddress,
) *PaypalTransaction {

	var t PaypalTransaction
	t.private.Amount = amount{Details: &details{}}
	t.private.ItemList = &paypalItemList{}

	t.private.Amount.private.Currency = c
	t.private.Amount.private.Total = 0

	t.private.ItemList.private.Items = make([]*PaypalItem, 0, 1)
	t.private.ItemList.private.ShippingAddress = shp

	pp.private.Transactions = append(pp.private.Transactions, &t)

	return &t
}

func (pp *PaypalPayment) calculateToAuthorize() {
	for _, t := range pp.private.Transactions {
		t.calculateToAuthorize()
	}
}

// TODO: Needs to validate some sub-properties that are valid only when
// Payer.PaymentMethod is "paypal"
func (pp *PaypalPayment) validate() (err error) {
	if len(pp.private.Transactions) == 0 {
		return fmt.Errorf("A Payment needs at least one Transaction")
	}

	for _, t := range pp.private.Transactions {
		if err = t.validate(); err != nil {
			return err
		}
	}

	// TODO: More validation

	return pp.private.Payer.validate()
}

// TODO: Should send a query string parameter `token=[some token]`
func (pp *PaypalPayment) Authorize() (to string, code int, err error) {
	if err = pp.validate(); err != nil {
		return "", 0, err
	}

	pp.calculateToAuthorize()

	// Create Totals
	var pymt PaypalPayment

	err = pp.send(&request{
		method:   method.Post,
		path:     _paymentsPath,
		body:     pp,
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
			err = fmt.Errorf("Unexpected state: %s", pymt.private.State.String())
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

	if pymt.private.State != State.Approved {
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
@struct
*/
type __PaypalTransaction struct {
	ItemList         *paypalItemList  `gRead json:"item_list,omitempty"`
	Amount           amount           `gRead json:"amount"`
	RelatedResources relatedResources `gRead json:"related_resources,omitempty"`
	Description      string           `gRead gWrite json:"description,omitempty"`
	PaymentOptions   *paymentOptions  `gRead gWrite json:"payment_options,omitempty"`
	InvoiceNumber    string           `gRead gWrite json:"invoice_number,omitempty"`
	Custom           string           `gRead gWrite json:"custom,omitempty"`
	SoftDescriptor   string           `gRead gWrite json:"soft_descriptor,omitempty"`
}

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
@struct
*/
type __paypalItemList struct {
	Items           PaypalItems      `gRead json:"items,omitempty"`
	ShippingAddress *ShippingAddress `gRead json:"shipping_address,omitempty"`
}

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
@struct
*/
// Source of the funds for this payment represented by a PayPal account.
type __paypalPayer struct {
	// Must be PaymentMethod.Paypal
	PaymentMethod PaymentMethodEnum `json:"payment_method"`

	// Status of the payer’s PayPal account. Allowed values: VERIFIED or UNVERIFIED.
	Status payerStatusEnum `gRead json:"status,omitempty"`

	PaypalPayerInfo *PaypalPayerInfo `gRead json:"payer_info,omitempty"`
}

func (pp *paypalPayer) validate() error {
	err := pp.private.PaymentMethod.validate()
	if err != nil {
		return err
	}
	return nil
}

/*
@struct
*/
// This object is pre-filled by PayPal when the payment_method is paypal.
type __PaypalPayerInfo struct {
	// Payer’s tax ID type. Allowed values: BR_CPF or BR_C`NPJ. Only supported when
	// the payment_method is set to paypal.
	TaxIdType TaxIdTypeEnum `gRead gWrite json:"tax_id_type,omitempty"`

	// Payer’s tax ID. Only supported when the payment_method is set to paypal.
	TaxId string `gRead gWrite json:"tax_id,omitempty"`

	// Email address representing the payer. 127 characters max.
	Email string `gRead json:"email,omitempty"`

	// Salutation of the payer.
	Salutation string `gRead json:"salutation,omitempty"`

	// Suffix of the payer.
	Suffix string `gRead json:"suffix,omitempty"`

	// Two-letter registered country code of the payer to identify the buyer country.
	CountryCode CountryCodeEnum `gRead json:"country_code,omitempty"`

	// Phone number representing the payer. 20 characters max.
	Phone string `gRead json:"phone,omitempty"`

	// First name of the payer. Value assigned by PayPal.
	FirstName string `gRead json:"first_name,omitempty"`

	// Middle name of the payer. Value assigned by PayPal.
	MiddleName string `gRead json:"middle_name,omitempty"`

	// Last name of the payer. Value assigned by PayPal.
	LastName string `gRead json:"last_name,omitempty"`

	// PayPal assigned Payer ID. Value assigned by PayPal.
	PayerId string `gRead json:"payer_id,omitempty"`

	// Shipping address of payer PayPal account. Value assigned by PayPal.
	ShippingAddress *ShippingAddress `gRead json:"shipping_address,omitempty"`
}

func (self *PaypalPayerInfo) validate() error {
	// TODO: Implement
	if self == nil {
		return nil
	}

	return nil
}
