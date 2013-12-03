package gopal

// Types of PayPal connection (Sandbox or Live)
type c_type bool
func (self c_type) getType() bool {
    return bool(self)
}
type connection_type_i interface {
    getType() bool
}
const Sandbox = c_type(false)
const Live = c_type(true)




type Intent string
const Sale = Intent("sale")
const Authorize = Intent("authorize")




type AddressType string
const Residential = AddressType("residential")
const Business = AddressType("business")
const Mailbox = AddressType("mailbox")




type PaymentMethod string
func (self PaymentMethod) payment_method() PaymentMethod {
	return self
}
type payment_method_i interface {
	payment_method() PaymentMethod
}
const CreditCard = PaymentMethod("credit_card")
const PayPal = PaymentMethod("paypal")




type CreditCardType string
func (self CreditCardType) credit_card_type() CreditCardType {
	return self
}
type credit_card_type_i interface {
	credit_card_type() CreditCardType
}
const Visa = CreditCardType("visa")
const MasterCard = CreditCardType("mastercard")
const Discover = CreditCardType("discover")
const Amex = CreditCardType("amex")




type State string

// payment
const Created = State("created")
const Approved = State("approved")
const Canceled = State("canceled")

// payment/refund
const Failed = State("failed")

// refund/sale/capture/authorization
const Pending = State("pending")

// refund/sale/capture
const Completed = State("completed")

// sale/capture
const Refunded = State("refunded")
const PartiallyRefunded = State("partially_refunded")

// payment/credit_card/authorization
const Expired = State("expired")

// credit_card
const Ok = State("ok")

// authorization
const Authorized = State("authorized")
const Captured = State("captured")
const PartiallyCaptured = State("partially_captured")
const Voided = State("voided")




type CurrencyType string
func (self CurrencyType) currency_type() CurrencyType {
	return self
}
type currency_type_i interface {
	currency_type() CurrencyType
}
const AUD = CurrencyType("AUD") // Australian dollar
const BRL = CurrencyType("BRL") // Brazilian real**
const CAD = CurrencyType("CAD") // Canadian dollar
const CZK = CurrencyType("CZK") // Czech koruna
const DKK = CurrencyType("DKK") // Danish krone
const EUR = CurrencyType("EUR") // Euro
const HKD = CurrencyType("HKD") // Hong Kong dollar
const HUF = CurrencyType("HUF") // Hungarian forint
const ILS = CurrencyType("ILS") // Israeli new shekel
const JPY = CurrencyType("JPY") // Japanese yen*
const MYR = CurrencyType("MYR") // Malaysian ringgit**
const MXN = CurrencyType("MXN") // Mexican peso
const TWD = CurrencyType("TWD") // New Taiwan dollar*
const NZD = CurrencyType("NZD") // New Zealand dollar
const NOK = CurrencyType("NOK") // Norwegian krone
const PHP = CurrencyType("PHP") // Philippine peso
const PLN = CurrencyType("PLN") // Polish złoty
const GBP = CurrencyType("GBP") // Pound sterling
const SGD = CurrencyType("SGD") // Singapore dollar
const SEK = CurrencyType("SEK") // Swedish krona
const CHF = CurrencyType("CHF") // Swiss franc
const THB = CurrencyType("THB") // Thai baht
const TRY = CurrencyType("TRY") // Turkish lira**
const USD = CurrencyType("USD") // United States dollar



