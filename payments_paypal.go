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

	pymt.Intent = intent.Sale
	pymt.Transactions = make([]*PaypalTransaction, 0)
	pymt.RedirectUrls = urls

	pymt.Payer.PaymentMethod = PaymentMethod.PayPal
	pymt.Payer.PaypalPayerInfo = info

	if err := urls.validate(); err != nil {
		return nil, err
	}
	if err := info.validate(); err != nil {
		return nil, err
	}

	return &pymt, nil
}

/*
@struct
*/
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

// TODO: Add `billing_agreement_tokens`, `payment_instruction`

/*
@struct
*/
type PaypalPayment struct {
	*connection
	Intent              intentEnum           `json:"intent,omitempty"`
	State               StateEnum            `json:"state,omitempty"`
	Id                  string               `json:"id,omitempty"`
	FailureReason       FailureReasonEnum    `json:"failure_reason,omitempty"`
	CreateTime          dateTime             `json:"create_time,omitempty"`
	UpdateTime          dateTime             `json:"update_time,omitempty"`
	Links               links                `json:"links,omitempty"`
	Transactions        []*PaypalTransaction `json:"transactions,omitempty"`
	Payer               paypalPayer          `json:"payer"`
	RedirectUrls        Redirects            `json:"redirect_urls,omitempty"`
	ExperienceProfileId string               `json:"experience_profile_id,omitempty"`
	*payment_error
}

func (pp *PaypalPayment) AddTransaction(
	c CurrencyTypeEnum,
	shp *ShippingAddress,
) *PaypalTransaction {

	var t PaypalTransaction
	t.Amount = amount{Details: &details{}}
	t.ItemList = &paypalItemList{}

	t.Amount.Currency = c
	t.Amount.Total = 0

	t.ItemList.Items = make([]*PaypalItem, 0, 1)
	t.ItemList.ShippingAddress = shp

	pp.Transactions = append(pp.Transactions, &t)

	return &t
}

func (pp *PaypalPayment) calculateToAuthorize() {
	for _, t := range pp.Transactions {
		t.calculateToAuthorize()
	}
}

// TODO: Needs to validate some sub-properties that are valid only when
// Payer.PaymentMethod is "paypal"
func (pp *PaypalPayment) validate() (err error) {
	if len(pp.Transactions) == 0 {
		return fmt.Errorf("A Payment needs at least one Transaction")
	}

	for _, t := range pp.Transactions {
		if err = t.validate(); err != nil {
			return err
		}
	}

	// TODO: More validation

	return pp.Payer.validate()
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
		switch pymt.State {
		case State.Created:
			// Set url to redirect to PayPal site to begin approval process
			to, _ = pymt.Links.get(relType.ApprovalUrl)
			code = 303
		default:
			// otherwise cancel the payment and return an error
			err = fmt.Errorf("Unexpected state: %s", pymt.State.String())
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
	Transactions []*PaypalTransaction `json:"transactions,omitempty"`
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

	if pymt.State != State.Approved {
		return fmt.Errorf(
			"Payment with ID %q for payer %q was not approved", pymtid, payerid)
	}

	return nil
}

// TODO: I'm returning a list of Sale objects because every payment can have multiple transactions.
//		I need to find out why a payment can have multiple transactions, and see if I should eliminate that in the API
func (self *PaypalPayment) FetchSale() []*Sale {
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

/*
@struct
*/
type PaypalTransaction struct {
	ItemList         *paypalItemList  `json:"item_list,omitempty"`
	Amount           amount           `json:"amount"`
	RelatedResources relatedResources `json:"related_resources,omitempty"`
	Description      string           `json:"description,omitempty"`
	PaymentOptions   *paymentOptions  `json:"payment_options,omitempty"`
	InvoiceNumber    string           `json:"invoice_number,omitempty"`
	Custom           string           `json:"custom,omitempty"`
	SoftDescriptor   string           `json:"soft_descriptor,omitempty"`
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

	item.Currency = t.Amount.Currency

	t.ItemList.Items =
		append(t.ItemList.Items, item)
	return nil
}

func (self *PaypalTransaction) validate() (err error) {
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

func (self *PaypalTransaction) calculateToAuthorize() {
	// Calculate totals from itemList
	for _, item := range self.ItemList.Items {
		self.Amount.Details.Subtotal = roundTwoDecimalPlaces(
			self.Amount.Details.Subtotal + (item.Price * float64(item.Quantity)))

		self.Amount.Details.Tax = roundTwoDecimalPlaces(
			self.Amount.Details.Tax + (item.Tax * float64(item.Quantity)))
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
type paypalItemList struct {
	Items           []*PaypalItem    `json:"items,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

func (self *paypalItemList) validate() (err error) {
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
// Source of the funds for this payment represented by a PayPal account.
type paypalPayer struct {
	// Must be PaymentMethod.Paypal
	PaymentMethod PaymentMethodEnum `json:"payment_method"`

	// Status of the payer’s PayPal account. Allowed values: VERIFIED or UNVERIFIED.
	Status payerStatusEnum `json:"status,omitempty"`

	PaypalPayerInfo *PaypalPayerInfo `json:"payer_info,omitempty"`
}

func (pp *paypalPayer) validate() error {
	err := pp.PaymentMethod.validate()
	if err != nil {
		return err
	}
	return nil
}

/*
@struct
*/
// This object is pre-filled by PayPal when the payment_method is paypal.
type PaypalPayerInfo struct {
	// Payer’s tax ID type. Allowed values: BR_CPF or BR_C`NPJ. Only supported when
	// the payment_method is set to paypal.
	TaxIdType TaxIdTypeEnum `json:"tax_id_type,omitempty"`

	// Payer’s tax ID. Only supported when the payment_method is set to paypal.
	TaxId string `json:"tax_id,omitempty"`

	// Email address representing the payer. 127 characters max.
	Email string `json:"email,omitempty"`

	// Salutation of the payer.
	Salutation string `json:"salutation,omitempty"`

	// Suffix of the payer.
	Suffix string `json:"suffix,omitempty"`

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
