package gopal

import (
	"fmt"
	"log"
	"strconv"
)

func make10CharAmount(amt float64) (string, error) {
	if amt < 0 {
		return "", fmt.Errorf("Amount may not be a negative number. Found: %f\n",
			amt)
	}

	var t string
	if (amt - float64(int(amt))) < .005 { // If no fraction, drop the decimal
		t = strconv.FormatInt(int64(amt), 10)
	} else {
		t = fmt.Sprintf("%.2f", amt)
	}

	if len(t) > 10 {
		return "", fmt.Errorf("Amount Total allows 10 chars max. Found: %q\n", t)
	}
	return t, nil
}

type Resource interface {
	errorable
	getPath() string
	getCurrencyType() CurrencyTypeEnum
	totalAsFloat() float64
}

type refundable interface {
	Resource
	getRefundPath() string
}

// The _times are assigned by PayPal in responses
type _times struct {
	CreateTime string `json:"create_time,omitempty"`
	UpdateTime string `json:"update_time,omitempty"`
}

type dateTime string // TODO: How should this be done? [Un]Marshalers?

type _trans struct {
	*connection

	_times

	Links links `json:"links,omitempty"`

	// ID of the Sale/Refund/Capture/Authorization. Read only
	Id string `json:"id,omitempty"`

	Amount amount `json:"amount"` // Details about the amount

	State stateEnum `json:"state,omitempty"` // State of the Sale/Refund/Capture/Auth

	ParentPayment string `json:"parent_payment,omitempty"`

	*identity_error
}

func (self *_trans) FetchParentPayment() (*Payment, error) {
	return self.FetchPayment(self.ParentPayment)
}

func (self _trans) getCurrencyType() CurrencyTypeEnum {
	return self.Amount.Currency
}

func (self _trans) totalAsFloat() float64 {
	f, err := strconv.ParseFloat(self.Amount.Total, 64)
	if err != nil {
		log.Printf("Invalid Amount.Total: %s\n", self.Amount.Total)
		return 0.0
	}
	return f
}

type links []link

func (l links) get(r relTypeEnum) (string, string) {
	for i, _ := range l {
		if l[i].Rel == r {
			return l[i].Href, l[i].Method
		}
	}
	return "", ""
}

type link struct {
	// URL of the related HATEOAS link you can use for subsequent calls
	Href string `json:"href,omitempty"`

	// Link relation that describes how this link relates to the previous call.
	// Examples include `self` (get details of the current call), `parent_payment`
	// (get details of the parent payment), or a related call such as `execute` or
	// `refund`
	Rel relTypeEnum `json:"rel,omitempty"`

	// The HTTP method required for the related call
	Method string `json:"method,omitempty"`
}

// Base object for all financial value related fields (balance, payment due, etc.)
type currency struct {
	// 3 letter currency code as defined by ISO 4217. Required.
	Currency string `json:"currency"`

	// amount up to N digit after the decimals separator as defined in ISO 4217
	// for the appropriate currency code. Required.
	Value string `json:"value"`
}

// This object represents Fraud Management Filter (FMF) details for a payment.
type fmfDetails struct {
	// Type of filter
	FilterType fmfFilter `json:"filter_type"`

	// Name of the fraud management filter. For more information about filters,
	// see Fraud Management Filters Summary.
	FilterID filterId `json:"filter_id"`

	// Name of the filter. This property is reserved for future use.
	Name string `json:"name"`

	// Description of the filter. This property is reserved for future use.
	Description string `json:"description"`
}

type fmfFilter string

const (
	accept  = fmfFilter("ACCEPT")  // An ACCEPT filter is triggered only for the TOTAL_PURCHASE_PRICE_MINIMUM filter setting and is returned only in direct credit card payments where payment is accepted.
	pending = fmfFilter("PENDING") // Triggers a PENDING filter action where you need to explicitly accept or deny the transaction.
	deny    = fmfFilter("DENY")    // Triggers a DENY action where payment is denied automatically.
	report  = fmfFilter("REPORT")  // Triggers the Flag testing mode where payment is accepted.
)

type filterId string

const (
	maximumTransactionAmount           = filterId("MAXIMUM_TRANSACTION_AMOUNT")             // basic filter
	unconfirmedAddress                 = filterId("UNCONFIRMED_ADDRESS")                    // basic filter
	countryMonitor                     = filterId("COUNTRY_MONITOR")                        // basic filter
	avsNoMatch                         = filterId("AVS_NO_MATCH")                           // Address Verification Service No Match (advanced filter)
	avsPartialMatch                    = filterId("AVS_PARTIAL_MATCH")                      // Address Verification Service Partial Match (advanced filter)
	avsUnavailableOrUnsupported        = filterId("AVS_UNAVAILABLE_OR_UNSUPPORTED")         // Address Verification Service Unavailable or Not Supported (advanced filter)
	cardSecurityCodeMismatch           = filterId("CARD_SECURITY_CODE_MISMATCH")            // advanced filter
	billingOrShippingAddressMismatch   = filterId("BILLING_OR_SHIPPING_ADDRESS_MISMATCH")   // advanced filter
	riskyZipCode                       = filterId("RISKY_ZIP_CODE")                         // high risk lists filter
	suspectedFreightForwarderCheck     = filterId("SUSPECTED_FREIGHT_FORWARDER_CHECK")      // high risk lists filter
	riskyEmailAddressDomainCheck       = filterId("RISKY_EMAIL_ADDRESS_DOMAIN_CHECK")       // high risk lists filter
	riskyBankIdentificationNumberCheck = filterId("RISKY_BANK_IDENTIFICATION_NUMBER_CHECK") // high risk lists filter
	riskyIpAddressRange                = filterId("RISKY_IP_ADDRESS_RANGE")                 // high risk lists filter
	largeOrderNumber                   = filterId("LARGE_ORDER_NUMBER")                     // transaction data filter
	totalPurchasePriceMinimum          = filterId("TOTAL_PURCHASE_PRICE_MINIMUM")           // transaction data filter
	ipAddressVelocity                  = filterId("IP_ADDRESS_VELOCITY")                    // transaction data filter
	paypalFraudModel                   = filterId("PAYPAL_FRAUD_MODEL")                     // transaction data filter
)
