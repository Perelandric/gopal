package gopal

// http://play.golang.org/p/rz7Qwt4mwH // Consts idea

// Sort order
type sort_by bool

func (self sort_by) getSortBy() bool {
	return bool(self)
}
func (self sort_by) String() string {
	if self.getSortBy() {
		return "create_time"
	} else {
		return "update_time"
	}
}

type sort_by_i interface {
	getSortBy() bool
	String() string
}

const CreateTime = sort_by(true)
const UpdateTime = sort_by(false)

// Ascending or descending order
type sort_order bool

func (self sort_order) getSortOrder() bool {
	return bool(self)
}
func (self sort_order) String() string {
	if self.getSortOrder() {
		return "DESC"
	} else {
		return "ASC"
	}
}

type sort_order_i interface {
	getSortOrder() bool
	String() string
}

const DESC = sort_order(true)
const ASC = sort_order(false)

type relType string

const (
	self          = relType("self")
	parentPayment = relType("parent_payment")
	execute       = relType("execute")
	refund        = relType("refund")
)

type method string

const (
	get  = method("GET")
	post = method("POST")
)

type payerStatus string

const (
	verified   = payerStatus("VERIFIED")
	unverified = payerStatus("UNVERIFIED")
)

type Intent string

const (
	sale      = Intent("sale")
	authorize = Intent("authorize")
	order     = Intent("order")
)

type normalizationStatus struct{ value string }

var (
	Unknown                   = normalizationStatus{"UNKNOWN"}
	UnnormalizedUserPreferred = normalizationStatus{"UNNORMALIZED_USER_PREFERRED"}
	Normalized                = normalizationStatus{"NORMALIZED"}
	Unnormalized              = normalizationStatus{"UNNORMALIZED"}
)

type addressStatus struct{ value string }

var (
	Confirmed   = addressStatus{"CONFIRMED"}
	Unconfirmed = addressStatus{"UNCONFIRMED"}
)

type addressType struct{ value string }

var (
	Residential = addressType{"residential"}
	Business    = addressType{"business"}
	Mailbox     = addressType{"mailbox"}
)

type PaymentMethod string

const (
	CreditCard = PaymentMethod("credit_card")
	PayPal     = PaymentMethod("paypal")
)

type TaxIdType string

const (
	br_cpf  = TaxIdType("BR_CPF")
	br_cnpj = TaxIdType("BR_CNPJ")
)

type CreditCardType string

const (
	Visa       = CreditCardType("visa")
	MasterCard = CreditCardType("mastercard")
	Discover   = CreditCardType("discover")
	Amex       = CreditCardType("amex")
)

type state string

// payment
const (
	Created    = state("created")
	Approved   = state("approved")
	Canceled   = state("canceled")
	InProgress = state("in_progress")
)

// payment/refund
const Failed = state("failed")

// payment/refund/sale/capture/authorization
const Pending = state("pending")

// refund/sale/capture
const Completed = state("completed")

// sale/capture
const Refunded = state("refunded")
const PartiallyRefunded = state("partially_refunded")

// payment/credit_card/authorization
const Expired = state("expired")

// credit_card
const Ok = state("ok")

// authorization
const Authorized = state("authorized")
const Captured = state("captured")
const PartiallyCaptured = state("partially_captured")
const Voided = state("voided")

type reasonCode string

const (
	chargeback                              = "CHARGEBACK"
	guarantee                               = "GUARANTEE"
	buyerComplaint                          = "BUYER_COMPLAINT"
	refundCode                              = "REFUND"
	unconfirmedShippingAddress              = "UNCONFIRMED_SHIPPING_ADDRESS"
	echeckCode                              = "ECHECK"
	internationalWithdrawal                 = "INTERNATIONAL_WITHDRAWAL"
	receivingPreferenceMandatesManualAction = "RECEIVING_PREFERENCE_MANDATES_MANUAL_ACTION"
	paymentReview                           = "PAYMENT_REVIEW"
	regulatoryReview                        = "REGULATORY_REVIEW"
	unilateral                              = "UNILATERAL"
	verificationRequired                    = "VERIFICATION_REQUIRED"
)

type protectionElig string

const (
	// Merchant is protected by PayPal's Seller Protection Policy for Unauthorized
	// Payments and Item Not Received.
	eligible = "ELIGIBLE"

	// Merchant is protected by PayPal's Seller Protection Policy for Item Not
	// Received or Unauthorized Payments. Refer to protection_eligibility_type for
	// specifics.
	partiallyEligible = "PARTIALLY_ELIGIBLE"

	// Merchant is not protected under the Seller Protection Policy.
	ineligibile = "INELIGIBLE"
)

type protectionEligType string

const ( // TODO: Both values may be represented, so a bitflag approach is needed

	// Sellers are protected against claims for items not received.
	itemNotReceivedEligible = protectionEligType("ITEM_NOT_RECEIVED_ELIGIBLE")

	// Sellers are protected against claims for unauthorized payments.
	unauthorizedPaymentEligible = protectionEligType("UNAUTHORIZED_PAYMENT_ELIGIBLE")
)

type paymentMode string

const (
	instantTransfer    = paymentMode("INSTANT_TRANSFER")
	manualBankTransfer = paymentMode("MANUAL_BANK_TRANSFER")
	delayedTransfer    = paymentMode("DELAYED_TRANSFER")
	echeck             = paymentMode("ECHECK")
)

type pendingReason string

const (
	payerShippingUnconfirmed   = pendingReason("PAYER-SHIPPING-UNCONFIRMED")
	multiCurrency              = pendingReason("MULTI-CURRENCY")
	riskReview                 = pendingReason("RISK-REVIEW")
	regulatoryReviewReason     = pendingReason("REGULATORY-REVIEW")     // *Modified to resolve conflict
	verificationRequiredReason = pendingReason("VERIFICATION-REQUIRED") // *Modified to resolve conflict
	orderPendingReason         = pendingReason("ORDER")                 // *Modified to resolve conflict
	otherPendingReason         = pendingReason("OTHER")                 // *Modified to resolve conflict
)

type currency_type string

const AUD = currency_type("AUD") // Australian dollar
const BRL = currency_type("BRL") // Brazilian real**
const CAD = currency_type("CAD") // Canadian dollar
const CZK = currency_type("CZK") // Czech koruna
const DKK = currency_type("DKK") // Danish krone
const EUR = currency_type("EUR") // Euro
const HKD = currency_type("HKD") // Hong Kong dollar
const HUF = currency_type("HUF") // Hungarian forint
const ILS = currency_type("ILS") // Israeli new shekel
const JPY = currency_type("JPY") // Japanese yen*
const MYR = currency_type("MYR") // Malaysian ringgit**
const MXN = currency_type("MXN") // Mexican peso
const TWD = currency_type("TWD") // New Taiwan dollar*
const NZD = currency_type("NZD") // New Zealand dollar
const NOK = currency_type("NOK") // Norwegian krone
const PHP = currency_type("PHP") // Philippine peso
const PLN = currency_type("PLN") // Polish złoty
const GBP = currency_type("GBP") // Pound sterling
const SGD = currency_type("SGD") // Singapore dollar
const SEK = currency_type("SEK") // Swedish krona
const CHF = currency_type("CHF") // Swiss franc
const THB = currency_type("THB") // Thai baht
const TRY = currency_type("TRY") // Turkish lira**
const USD = currency_type("USD") // United States dollar

type country_code string

const (
	_AL = country_code("AL") // ALBANIA
	_DZ = country_code("DZ") // ALGERIA
	_AD = country_code("AD") // ANDORRA
	_AO = country_code("AO") // ANGOLA
	_AI = country_code("AI") // ANGUILLA
	_AG = country_code("AG") // ANTIGUA AND BARBUDA
	_AR = country_code("AR") // ARGENTINA
	_AM = country_code("AM") // ARMENIA
	_AW = country_code("AW") // ARUBA
	_AU = country_code("AU") // AUSTRALIA
	_AT = country_code("AT") // AUSTRIA
	_AZ = country_code("AZ") // AZERBAIJAN
	_BS = country_code("BS") // BAHAMAS
	_BH = country_code("BH") // BAHRAIN
	_BB = country_code("BB") // BARBADOS
	_BE = country_code("BE") // BELGIUM
	_BZ = country_code("BZ") // BELIZE
	_BJ = country_code("BJ") // BENIN
	_BM = country_code("BM") // BERMUDA
	_BT = country_code("BT") // BHUTAN
	_BO = country_code("BO") // BOLIVIA
	_BA = country_code("BA") // BOSNIA-HERZEGOVINA
	_BW = country_code("BW") // BOTSWANA
	_BR = country_code("BR") // BRAZIL
	_BN = country_code("BN") // BRUNEI DARUSSALAM
	_BG = country_code("BG") // BULGARIA
	_BF = country_code("BF") // BURKINA FASO
	_BI = country_code("BI") // BURUNDI
	_KH = country_code("KH") // CAMBODIA
	_CA = country_code("CA") // CANADA
	_CV = country_code("CV") // CAPE VERDE
	_KY = country_code("KY") // CAYMAN ISLANDS
	_TD = country_code("TD") // CHAD
	_CL = country_code("CL") // CHILE
	_CN = country_code("CN") // CHINA (For domestic Chinese bank transactions only)
	_C2 = country_code("C2") // CHINA (For CUP, bank card and cross-border transactions)
	_CO = country_code("CO") // COLOMBIA
	_KM = country_code("KM") // COMOROS
	_CD = country_code("CD") // DEMOCRATIC REPUBLIC OF CONGO
	_CG = country_code("CG") // CONGO
	_CK = country_code("CK") // COOK ISLANDS
	_CR = country_code("CR") // COSTA RICA
	_HR = country_code("HR") // CROATIA
	_CY = country_code("CY") // CYPRUS
	_CZ = country_code("CZ") // CZECH REPUBLIC
	_DK = country_code("DK") // DENMARK
	_DJ = country_code("DJ") // DJIBOUTI
	_DM = country_code("DM") // DOMINICA
	_DO = country_code("DO") // DOMINICAN REPUBLIC
	_EC = country_code("EC") // ECUADOR
	_EG = country_code("EG") // EGYPT
	_SV = country_code("SV") // EL SALVADOR
	_ER = country_code("ER") // ERITERIA
	_EE = country_code("EE") // ESTONIA
	_ET = country_code("ET") // ETHIOPIA
	_FK = country_code("FK") // FALKLAND ISLANDS (MALVINAS)
	_FJ = country_code("FJ") // FIJI
	_FI = country_code("FI") // FINLAND
	_FR = country_code("FR") // FRANCE
	_GF = country_code("GF") // FRENCH GUIANA
	_PF = country_code("PF") // FRENCH POLYNESIA
	_GA = country_code("GA") // GABON
	_GM = country_code("GM") // GAMBIA
	_GE = country_code("GE") // GEORGIA
	_DE = country_code("DE") // GERMANY
	_GI = country_code("GI") // GIBRALTAR
	_GR = country_code("GR") // GREECE
	_GL = country_code("GL") // GREENLAND
	_GD = country_code("GD") // GRENADA
	_GP = country_code("GP") // GUADELOUPE
	_GU = country_code("GU") // GUAM
	_GT = country_code("GT") // GUATEMALA
	_GN = country_code("GN") // GUINEA
	_GW = country_code("GW") // GUINEA BISSAU
	_GY = country_code("GY") // GUYANA
	_VA = country_code("VA") // HOLY SEE (VATICAN CITY STATE)
	_HN = country_code("HN") // HONDURAS
	_HK = country_code("HK") // HONG KONG
	_HU = country_code("HU") // HUNGARY
	_IS = country_code("IS") // ICELAND
	_IN = country_code("IN") // INDIA
	_ID = country_code("ID") // INDONESIA
	_IE = country_code("IE") // IRELAND
	_IL = country_code("IL") // ISRAEL
	_IT = country_code("IT") // ITALY
	_JM = country_code("JM") // JAMAICA
	_JP = country_code("JP") // JAPAN
	_JO = country_code("JO") // JORDAN
	_KZ = country_code("KZ") // KAZAKHSTAN
	_KE = country_code("KE") // KENYA
	_KI = country_code("KI") // KIRIBATI
	_KR = country_code("KR") // KOREA, REPUBLIC OF (or SOUTH KOREA)
	_KW = country_code("KW") // KUWAIT
	_KG = country_code("KG") // KYRGYZSTAN
	_LA = country_code("LA") // LAOS
	_LV = country_code("LV") // LATVIA
	_LS = country_code("LS") // LESOTHO
	_LI = country_code("LI") // LIECHTENSTEIN
	_LT = country_code("LT") // LITHUANIA
	_LU = country_code("LU") // LUXEMBOURG
	_MG = country_code("MG") // MADAGASCAR
	_MW = country_code("MW") // MALAWI
	_MY = country_code("MY") // MALAYSIA
	_MV = country_code("MV") // MALDIVES
	_ML = country_code("ML") // MALI
	_MT = country_code("MT") // MALTA
	_MH = country_code("MH") // MARSHALL ISLANDS
	_MQ = country_code("MQ") // MARTINIQUE
	_MR = country_code("MR") // MAURITANIA
	_MU = country_code("MU") // MAURITIUS
	_YT = country_code("YT") // MAYOTTE
	_MX = country_code("MX") // MEXICO
	_FM = country_code("FM") // MICRONESIA, FEDERATED STATES OF
	_MN = country_code("MN") // MONGOLIA
	_MS = country_code("MS") // MONTSERRAT
	_MA = country_code("MA") // MOROCCO
	_MZ = country_code("MZ") // MOZAMBIQUE
	_NA = country_code("NA") // NAMIBIA
	_NR = country_code("NR") // NAURU
	_NP = country_code("NP") // NEPAL
	_NL = country_code("NL") // NETHERLANDS
	_AN = country_code("AN") // NETHERLANDS ANTILLES
	_NC = country_code("NC") // NEW CALEDONIA
	_NZ = country_code("NZ") // NEW ZEALAND
	_NI = country_code("NI") // NICARAGUA
	_NE = country_code("NE") // NIGER
	_NU = country_code("NU") // NIUE
	_NF = country_code("NF") // NORFOLK ISLAND
	_NO = country_code("NO") // NORWAY
	_OM = country_code("OM") // OMAN
	_PW = country_code("PW") // PALAU
	_PA = country_code("PA") // PANAMA
	_PG = country_code("PG") // PAPUA NEW GUINEA
	_PE = country_code("PE") // PERU
	_PH = country_code("PH") // PHILIPPINES
	_PN = country_code("PN") // PITCAIRN
	_PL = country_code("PL") // POLAND
	_PT = country_code("PT") // PORTUGAL
	_QA = country_code("QA") // QATAR
	_RE = country_code("RE") // REUNION
	_RO = country_code("RO") // ROMANIA
	_RU = country_code("RU") // RUSSIAN FEDERATION
	_RW = country_code("RW") // RWANDA
	_SH = country_code("SH") // SAINT HELENA
	_KN = country_code("KN") // SAINT KITTS AND NEVIS
	_LC = country_code("LC") // SAINT LUCIA
	_PM = country_code("PM") // SAINT PIERRE AND MIQUELON
	_VC = country_code("VC") // SAINT VINCENT AND THE GRENADINES
	_WS = country_code("WS") // SAMOA
	_SM = country_code("SM") // SAN MARINO
	_ST = country_code("ST") // SAO TOME AND PRINCIPE
	_SA = country_code("SA") // SAUDI ARABIA
	_SN = country_code("SN") // SENEGAL
	_RS = country_code("RS") // SERBIA
	_SC = country_code("SC") // SEYCHELLES
	_SL = country_code("SL") // SIERRA LEONE
	_SG = country_code("SG") // SINGAPORE
	_SK = country_code("SK") // SLOVAKIA
	_SI = country_code("SI") // SLOVENIA
	_SB = country_code("SB") // SOLOMON ISLANDS
	_SO = country_code("SO") // SOMALIA
	_ZA = country_code("ZA") // SOUTH AFRICA
	_ES = country_code("ES") // SPAIN
	_LK = country_code("LK") // SRI LANKA
	_SR = country_code("SR") // SURINAME
	_SJ = country_code("SJ") // SVALBARD AND JAN MAYEN
	_SZ = country_code("SZ") // SWAZILAND
	_SE = country_code("SE") // SWEDEN
	_CH = country_code("CH") // SWITZERLAND
	_TW = country_code("TW") // TAIWAN, PROVINCE OF CHINA
	_TJ = country_code("TJ") // TAJIKISTAN
	_TZ = country_code("TZ") // TANZANIA, UNITED REPUBLIC OF
	_TH = country_code("TH") // THAILAND
	_TG = country_code("TG") // TOGO
	_TO = country_code("TO") // TONGA
	_TT = country_code("TT") // TRINIDAD AND TOBAGO
	_TN = country_code("TN") // TUNISIA
	_TR = country_code("TR") // TURKEY
	_TM = country_code("TM") // TURKMENISTAN
	_TC = country_code("TC") // TURKS AND CAICOS ISLANDS
	_TV = country_code("TV") // TUVALU
	_UG = country_code("UG") // UGANDA
	_UA = country_code("UA") // UKRAINE
	_AE = country_code("AE") // UNITED ARAB EMIRATES
	_GB = country_code("GB") // UNITED KINGDOM
	_US = country_code("US") // UNITED STATES
	_UY = country_code("UY") // URUGUAY
	_VU = country_code("VU") // VANUATU
	_VE = country_code("VE") // VENEZUELA
	_VN = country_code("VN") // VIETNAM
	_VG = country_code("VG") // VIRGIN ISLANDS, BRITISH
	_WF = country_code("WF") // WALLIS AND FUTUNA
	_YE = country_code("YE") // YEMEN
	_ZM = country_code("ZM") // ZAMBIA
)
