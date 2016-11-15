package gopal

import "fmt"

//go:generate Golific $GOFILE

const (
	AuthorizationType = "authorization"
	CaptureType       = "capture"
	RefundType        = "refund"
	SaleType          = "sale"
)

const (
	_paymentsPath      = "payments/payment"
	_capturePath       = "payments/capture"
	_refundPath        = "payments/refund"
	_authorizationPath = "payments/authorization"
	_salePath          = "payments/sale"
	_execute           = "execute"
	_refund            = "refund"
	_reauthorize       = "reauthorize"
	_capture           = "capture"
	_void              = "void"
)

// Re-export public variants
var (
	Live    = Server.Live
	Sandbox = Server.Sandbox

	PayPal     = PaymentMethod.PayPal
	CreditCard = PaymentMethod.CreditCard
)

type enumValidator interface {
	validate() error
	IntValue() int
	Namespace() string
	IsDefault() bool
}

func validateEnum(enum enumValidator) error {
	if enum.IntValue() == 0 && enum.IsDefault() == false {
		return fmt.Errorf("%q is required", enum.Namespace())
	}
	return nil
}

/*
@enum
*/
type __Server struct {
	Live    int
	Sandbox int
}

/*
@enum json:"string"
*/
type __sortBy struct {
	CreateTime int `gString:"create_time"`
	UpdateTime int `gString:"update_time"`
}

/*
@enum json:"string"
*/
type __sortOrder struct {
	DESC int
	ASC  int
}

/*
@enum json:"string"
*/
type __relType struct {
	Self          int
	ParentPayment int `gString:"parent_payment"`
	Execute       int
	Refund        int
	ApprovalUrl   int `gString:"approval_url"`
	Suspend       int
	ReActivate    int `gString:"re_activate"`
	Cancel        int
	Void          int
	Authorization int
	Capture       int
	Reauthorize   int
	Order         int
	Item          int
	Batch         int
	Delete        int
	Patch         int
	First         int
	Last          int
	Update        int
	Resend        int
	Next          int
	Previous      int
	Start         int
	NextPage      int `gString:"next_page"`
	PreviousPage  int `gString:"previous_page"`
}

/*
@enum json:"string"
*/
type __method struct {
	Get      int `gString:"GET"`
	Post     int `gString:"POST"`
	Redirect int `gString:"REDIRECT"`
	Delete   int `gString:"DELETE"`
	Patch    int `gString:"PATCH"`
}

/*
@enum json:"string"
*/
type __payerStatus struct {
	Verified   int `gString:"VERIFIED"`
	Unverified int `gString:"UNVERIFIED"`
}

/*
@enum json:"string"
*/
type __intent struct {
	Sale      int `gString:"sale"`
	Authorize int `gString:"authorize"`
	Order     int `gString:"order"`
}

/*
@enum json:"string"
*/
type __FailureReason struct {
	UnableToCompleteTransaction int `gString:"UNABLE_TO_COMPLETE_TRANSACTION"`
	InvalidPaymentMethod        int `gString:"INVALID_PAYMENT_METHOD"`
	PayerCannotPay              int `gString:"PAYER_CANNOT_PAY"`
	CannotPayThisPayee          int `gString:"CANNOT_PAY_THIS_PAYEE"`
	RedirectRequired            int `gString:"REDIRECT_REQUIRED"`
	PayeeFilterRestrictions     int `gString:"PAYEE_FILTER_RESTRICTIONS"`
}

/*
@enum json:"string"
*/
type __fmfFilter struct {
	Accept  int `gString:"ACCEPT" gDescription:"An ACCEPT filter is triggered only for the TOTAL_PURCHASE_PRICE_MINIMUM filter setting and is returned only in direct credit card payments where payment is accepted."`
	Pending int `gString:"PENDING" gDescription:"Triggers a PENDING filter action where you need to explicitly accept or deny the transaction."`
	Deny    int `gString:"DENY" gDescription:"Triggers a DENY action where payment is denied automatically."`
	Report  int `gString:"REPORT" gDescription:"Triggers the Flag testing mode where payment is accepted."`
}

/*
@enum
*/
type __filterId struct {
	MaximumTransactionAmount           int `gString:"MAXIMUM_TRANSACTION_AMOUNT" gDescription:"basic filter"`
	UnconfirmedAddress                 int `gString:"UNCONFIRMED_ADDRESS" gDescription:"basic filter"`
	CountryMonitor                     int `gString:"COUNTRY_MONITOR" gDescription:"basic filter"`
	AvsNoMatch                         int `gString:"AVS_NO_MATCH" gDescription:"Address Verification Service No Match (advanced filter)"`
	AvsPartialMatch                    int `gString:"AVS_PARTIAL_MATCH" gDescription:"Address Verification Service Partial Match (advanced filter)"`
	AvsUnavailableOrUnsupported        int `gString:"AVS_UNAVAILABLE_OR_UNSUPPORTED" gDescription:"Address Verification Service Unavailable or Not Supported (advanced filter)"`
	CardSecurityCodeMismatch           int `gString:"CARD_SECURITY_CODE_MISMATCH" gDescription:"advanced filter"`
	BillingOrShippingAddressMismatch   int `gString:"BILLING_OR_SHIPPING_ADDRESS_MISMATCH" gDescription:"advanced filter"`
	RiskyZipCode                       int `gString:"RISKY_ZIP_CODE" gDescription:"high risk lists filter"`
	SuspectedFreightForwarderCheck     int `gString:"SUSPECTED_FREIGHT_FORWARDER_CHECK" gDescription:"high risk lists filter"`
	RiskyEmailAddressDomainCheck       int `gString:"RISKY_EMAIL_ADDRESS_DOMAIN_CHECK" gDescription:"high risk lists filter"`
	RiskyBankIdentificationNumberCheck int `gString:"RISKY_BANK_IDENTIFICATION_NUMBER_CHECK" gDescription:"high risk lists filter"`
	RiskyIpAddressRange                int `gString:"RISKY_IP_ADDRESS_RANGE" gDescription:"high risk lists filter"`
	LargeOrderNumber                   int `gString:"LARGE_ORDER_NUMBER" gDescription:"transaction data filter"`
	TotalPurchasePriceMinimum          int `gString:"TOTAL_PURCHASE_PRICE_MINIMUM" gDescription:"transaction data filter"`
	IpAddressVelocity                  int `gString:"IP_ADDRESS_VELOCITY" gDescription:"transaction data filter"`
	PaypalFraudModel                   int `gString:"PAYPAL_FRAUD_MODEL" gDescription:"transaction data filter"`
}

/*
@enum json:"string"
*/
type __normStatus struct {
	Unknown                   int `gString:"UNKNOWN"`
	UnnormalizedUserPreferred int `gString:"UNNORMALIZED_USER_PREFERRED"`
	Normalized                int `gString:"NORMALIZED"`
	Unnormalized              int `gString:"UNNORMALIZED"`
}

/*
@enum json:"string"
*/
type __addressStatus struct {
	Confirmed   int `gString:"CONFIRMED"`
	Unconfirmed int `gString:"UNCONFIRMED"`
}

/*
@enum json:"string"
*/
type __AddressType struct {
	Residential int `gString:"residential"`
	Business    int `gString:"business"`
	Mailbox     int `gString:"mailbox"`
}

// This should only be invoked when the AddressType is required.
func (self AddressTypeEnum) validate() error {
	return validateEnum(self)
}

/*
@enum json:"string"
*/
type __PaymentMethod struct {
	CreditCard int `gString:"credit_card"`
	PayPal     int `gString:"paypal"`
}

func (self PaymentMethodEnum) validate() error {
	return validateEnum(self)
}

/*
@enum json:"string"
*/
type __TaxIdType struct {
	BrCpf  int `gString:"BR_CPF"`
	BrCnpj int `gString:"BR_CNPJ"`
}

/*
@enum json:"string"
*/
type __CreditCardType struct {
	Visa       int `gString:"visa"`
	MasterCard int `gString:"mastercard"`
	Discover   int `gString:"discover"`
	Amex       int `gString:"amex"`
}

/*
@enum json:"string"
*/
type __State struct {
	Created           int `gString:"created"`
	Approved          int `gString:"approved"`
	Canceled          int `gString:"canceled"`
	InProgress        int `gString:"in_progress"`
	Failed            int `gString:"failed"`
	Pending           int `gString:"pending"`
	Completed         int `gString:"completed"`
	Refunded          int `gString:"refunded"`
	PartiallyRefunded int `gString:"partially_refunded"`
	Expired           int `gString:"expired"`
	Ok                int `gString:"ok"`
	Authorized        int `gString:"authorized"`
	Captured          int `gString:"captured"`
	PartiallyCaptured int `gString:"partially_captured"`
	Voided            int `gString:"voided"`
}

/*
@enum json:"string"
*/
type __reasonCode struct {
	Chargeback                              int `gString:"CHARGEBACK"`
	Guarantee                               int `gString:"GUARANTEE"`
	BuyerComplaint                          int `gString:"BUYER_COMPLAINT"`
	RefundCode                              int `gString:"REFUND"`
	UnconfirmedShippingAddress              int `gString:"UNCONFIRMED_SHIPPING_ADDRESS"`
	EcheckCode                              int `gString:"ECHECK"`
	InternationalWithdrawal                 int `gString:"INTERNATIONAL_WITHDRAWAL"`
	ReceivingPreferenceMandatesManualAction int `gString:"RECEIVING_PREFERENCE_MANDATES_MANUAL_ACTION"`
	PaymentReview                           int `gString:"PAYMENT_REVIEW"`
	RegulatoryReview                        int `gString:"REGULATORY_REVIEW"`
	Unilateral                              int `gString:"UNILATERAL"`
	VerificationRequired                    int `gString:"VERIFICATION_REQUIRED"`
}

/*
@enum json:"string"
*/
type __protectionElig struct {
	Eligible          int `gString:"ELIGIBLE" gDescription:"Merchant is protected by PayPal's Seller Protection Policy for Unauthorized. Payments and Item Not Received."`
	PartiallyEligible int `gString:"PARTIALLY_ELIGIBLE" gDescription:"Merchant is protected by PayPal's Seller Protection Policy for Item Not Received or Unauthorized Payments. Refer to protection_eligibility_type for specifics."`
	Ineligibile       int `gString:"INELIGIBLE" gDescription:"Merchant is not protected under the Seller Protection Policy."`
}

/*
@enum json:"string"
*/
type __protectionEligType struct {
	ItemNotReceivedEligible     int `gString:"ITEM_NOT_RECEIVED_ELIGIBLE" gDescription:"Sellers are protected against claims for items not received."`
	UnauthorizedPaymentEligible int `gString:"UNAUTHORIZED_PAYMENT_ELIGIBLE" gDescription:"Sellers are protected against claims for unauthorized payments."`
}

/*
@enum json:"string"
*/
type __paymentMode struct {
	InstantTransfer    int `gString:"INSTANT_TRANSFER"`
	ManualBankTransfer int `gString:"MANUAL_BANK_TRANSFER"`
	DelayedTransfer    int `gString:"DELAYED_TRANSFER"`
	Echeck             int `gString:"ECHECK"`
}

/*
@enum json:"string"
*/
type __pendingReason struct {
	PayerShippingUnconfirmed int `gString:"PAYER-SHIPPING-UNCONFIRMED"`
	MultiCurrency            int `gString:"MULTI-CURRENCY"`
	RiskReview               int `gString:"RISK-REVIEW"`
	RegulatoryReview         int `gString:"REGULATORY-REVIEW"`
	VerificationRequired     int `gString:"VERIFICATION-REQUIRED"`
	OrderPending             int `gString:"ORDER"`
	OtherPending             int `gString:"OTHER"`
}

/*
@enum json:"string"
*/
type __CurrencyType struct {
	AUD int `gString:"AUD" gDescription:"Australian dollar"`
	BRL int `gString:"BRL" gDescription:"Brazilian real**"`
	CAD int `gString:"CAD" gDescription:"Canadian dollar"`
	CZK int `gString:"CZK" gDescription:"Czech koruna"`
	DKK int `gString:"DKK" gDescription:"Danish krone"`
	EUR int `gString:"EUR" gDescription:"Euro"`
	HKD int `gString:"HKD" gDescription:"Hong Kong dollar"`
	HUF int `gString:"HUF" gDescription:"Hungarian forint"`
	ILS int `gString:"ILS" gDescription:"Israeli new shekel"`
	JPY int `gString:"JPY" gDescription:"Japanese yen*"`
	MYR int `gString:"MYR" gDescription:"Malaysian ringgit**"`
	MXN int `gString:"MXN" gDescription:"Mexican peso"`
	TWD int `gString:"TWD" gDescription:"New Taiwan dollar*"`
	NZD int `gString:"NZD" gDescription:"New Zealand dollar"`
	NOK int `gString:"NOK" gDescription:"Norwegian krone"`
	PHP int `gString:"PHP" gDescription:"Philippine peso"`
	PLN int `gString:"PLN" gDescription:"Polish złoty"`
	GBP int `gString:"GBP" gDescription:"Pound sterling"`
	SGD int `gString:"SGD" gDescription:"Singapore dollar"`
	SEK int `gString:"SEK" gDescription:"Swedish krona"`
	CHF int `gString:"CHF" gDescription:"Swiss franc"`
	THB int `gString:"THB" gDescription:"Thai baht"`
	TRY int `gString:"TRY" gDescription:"Turkish lira**"`
	USD int `gString:"USD" gDescription:"United States dollar"`
}

// This should only be invoked when the AddressType is required.
func (self CurrencyTypeEnum) validate() error {
	return validateEnum(self)
}

/*
@enum json:"string"
*/
type __CountryCode struct {
	AL int `gDescription:"ALBANIA"`
	DZ int `gDescription:"ALGERIA"`
	AD int `gDescription:"ANDORRA"`
	AO int `gDescription:"ANGOLA"`
	AI int `gDescription:"ANGUILLA"`
	AG int `gDescription:"ANTIGUA AND BARBUDA"`
	AR int `gDescription:"ARGENTINA"`
	AM int `gDescription:"ARMENIA"`
	AW int `gDescription:"ARUBA"`
	AU int `gDescription:"AUSTRALIA"`
	AT int `gDescription:"AUSTRIA"`
	AZ int `gDescription:"AZERBAIJAN"`
	BS int `gDescription:"BAHAMAS"`
	BH int `gDescription:"BAHRAIN"`
	BB int `gDescription:"BARBADOS"`
	BE int `gDescription:"BELGIUM"`
	BZ int `gDescription:"BELIZE"`
	BJ int `gDescription:"BENIN"`
	BM int `gDescription:"BERMUDA"`
	BT int `gDescription:"BHUTAN"`
	BO int `gDescription:"BOLIVIA"`
	BA int `gDescription:"BOSNIA-HERZEGOVINA"`
	BW int `gDescription:"BOTSWANA"`
	BR int `gDescription:"BRAZIL"`
	BN int `gDescription:"BRUNEI DARUSSALAM"`
	BG int `gDescription:"BULGARIA"`
	BF int `gDescription:"BURKINA FASO"`
	BI int `gDescription:"BURUNDI"`
	KH int `gDescription:"CAMBODIA"`
	CA int `gDescription:"CANADA"`
	CV int `gDescription:"CAPE VERDE"`
	KY int `gDescription:"CAYMAN ISLANDS"`
	TD int `gDescription:"CHAD"`
	CL int `gDescription:"CHILE"`
	CN int `gDescription:"CHINA (For domestic Chinese bank transactions only)"`
	C2 int `gDescription:"CHINA (For CUP, bank card and cross-border transactions)"`
	CO int `gDescription:"COLOMBIA"`
	KM int `gDescription:"COMOROS"`
	CD int `gDescription:"DEMOCRATIC REPUBLIC OF CONGO"`
	CG int `gDescription:"CONGO"`
	CK int `gDescription:"COOK ISLANDS"`
	CR int `gDescription:"COSTA RICA"`
	HR int `gDescription:"CROATIA"`
	CY int `gDescription:"CYPRUS"`
	CZ int `gDescription:"CZECH REPUBLIC"`
	DK int `gDescription:"DENMARK"`
	DJ int `gDescription:"DJIBOUTI"`
	DM int `gDescription:"DOMINICA"`
	DO int `gDescription:"DOMINICAN REPUBLIC"`
	EC int `gDescription:"ECUADOR"`
	EG int `gDescription:"EGYPT"`
	SV int `gDescription:"EL SALVADOR"`
	ER int `gDescription:"ERITERIA"`
	EE int `gDescription:"ESTONIA"`
	ET int `gDescription:"ETHIOPIA"`
	FK int `gDescription:"FALKLAND ISLANDS (MALVINAS)"`
	FJ int `gDescription:"FIJI"`
	FI int `gDescription:"FINLAND"`
	FR int `gDescription:"FRANCE"`
	GF int `gDescription:"FRENCH GUIANA"`
	PF int `gDescription:"FRENCH POLYNESIA"`
	GA int `gDescription:"GABON"`
	GM int `gDescription:"GAMBIA"`
	GE int `gDescription:"GEORGIA"`
	DE int `gDescription:"GERMANY"`
	GI int `gDescription:"GIBRALTAR"`
	GR int `gDescription:"GREECE"`
	GL int `gDescription:"GREENLAND"`
	GD int `gDescription:"GRENADA"`
	GP int `gDescription:"GUADELOUPE"`
	GU int `gDescription:"GUAM"`
	GT int `gDescription:"GUATEMALA"`
	GN int `gDescription:"GUINEA"`
	GW int `gDescription:"GUINEA BISSAU"`
	GY int `gDescription:"GUYANA"`
	VA int `gDescription:"HOLY SEE (VATICAN CITY STATE)"`
	HN int `gDescription:"HONDURAS"`
	HK int `gDescription:"HONG KONG"`
	HU int `gDescription:"HUNGARY"`
	IS int `gDescription:"ICELAND"`
	IN int `gDescription:"INDIA"`
	ID int `gDescription:"INDONESIA"`
	IE int `gDescription:"IRELAND"`
	IL int `gDescription:"ISRAEL"`
	IT int `gDescription:"ITALY"`
	JM int `gDescription:"JAMAICA"`
	JP int `gDescription:"JAPAN"`
	JO int `gDescription:"JORDAN"`
	KZ int `gDescription:"KAZAKHSTAN"`
	KE int `gDescription:"KENYA"`
	KI int `gDescription:"KIRIBATI"`
	KR int `gDescription:"KOREA, REPUBLIC OF (or SOUTH KOREA)"`
	KW int `gDescription:"KUWAIT"`
	KG int `gDescription:"KYRGYZSTAN"`
	LA int `gDescription:"LAOS"`
	LV int `gDescription:"LATVIA"`
	LS int `gDescription:"LESOTHO"`
	LI int `gDescription:"LIECHTENSTEIN"`
	LT int `gDescription:"LITHUANIA"`
	LU int `gDescription:"LUXEMBOURG"`
	MG int `gDescription:"MADAGASCAR"`
	MW int `gDescription:"MALAWI"`
	MY int `gDescription:"MALAYSIA"`
	MV int `gDescription:"MALDIVES"`
	ML int `gDescription:"MALI"`
	MT int `gDescription:"MALTA"`
	MH int `gDescription:"MARSHALL ISLANDS"`
	MQ int `gDescription:"MARTINIQUE"`
	MR int `gDescription:"MAURITANIA"`
	MU int `gDescription:"MAURITIUS"`
	YT int `gDescription:"MAYOTTE"`
	MX int `gDescription:"MEXICO"`
	FM int `gDescription:"MICRONESIA, FEDERATED STATES OF"`
	MN int `gDescription:"MONGOLIA"`
	MS int `gDescription:"MONTSERRAT"`
	MA int `gDescription:"MOROCCO"`
	MZ int `gDescription:"MOZAMBIQUE"`
	NA int `gDescription:"NAMIBIA"`
	NR int `gDescription:"NAURU"`
	NP int `gDescription:"NEPAL"`
	NL int `gDescription:"NETHERLANDS"`
	AN int `gDescription:"NETHERLANDS ANTILLES"`
	NC int `gDescription:"NEW CALEDONIA"`
	NZ int `gDescription:"NEW ZEALAND"`
	NI int `gDescription:"NICARAGUA"`
	NE int `gDescription:"NIGER"`
	NU int `gDescription:"NIUE"`
	NF int `gDescription:"NORFOLK ISLAND"`
	NO int `gDescription:"NORWAY"`
	OM int `gDescription:"OMAN"`
	PW int `gDescription:"PALAU"`
	PA int `gDescription:"PANAMA"`
	PG int `gDescription:"PAPUA NEW GUINEA"`
	PE int `gDescription:"PERU"`
	PH int `gDescription:"PHILIPPINES"`
	PN int `gDescription:"PITCAIRN"`
	PL int `gDescription:"POLAND"`
	PT int `gDescription:"PORTUGAL"`
	QA int `gDescription:"QATAR"`
	RE int `gDescription:"REUNION"`
	RO int `gDescription:"ROMANIA"`
	RU int `gDescription:"RUSSIAN FEDERATION"`
	RW int `gDescription:"RWANDA"`
	SH int `gDescription:"SAINT HELENA"`
	KN int `gDescription:"SAINT KITTS AND NEVIS"`
	LC int `gDescription:"SAINT LUCIA"`
	PM int `gDescription:"SAINT PIERRE AND MIQUELON"`
	VC int `gDescription:"SAINT VINCENT AND THE GRENADINES"`
	WS int `gDescription:"SAMOA"`
	SM int `gDescription:"SAN MARINO"`
	ST int `gDescription:"SAO TOME AND PRINCIPE"`
	SA int `gDescription:"SAUDI ARABIA"`
	SN int `gDescription:"SENEGAL"`
	RS int `gDescription:"SERBIA"`
	SC int `gDescription:"SEYCHELLES"`
	SL int `gDescription:"SIERRA LEONE"`
	SG int `gDescription:"SINGAPORE"`
	SK int `gDescription:"SLOVAKIA"`
	SI int `gDescription:"SLOVENIA"`
	SB int `gDescription:"SOLOMON ISLANDS"`
	SO int `gDescription:"SOMALIA"`
	ZA int `gDescription:"SOUTH AFRICA"`
	ES int `gDescription:"SPAIN"`
	LK int `gDescription:"SRI LANKA"`
	SR int `gDescription:"SURINAME"`
	SJ int `gDescription:"SVALBARD AND JAN MAYEN"`
	SZ int `gDescription:"SWAZILAND"`
	SE int `gDescription:"SWEDEN"`
	CH int `gDescription:"SWITZERLAND"`
	TW int `gDescription:"TAIWAN, PROVINCE OF CHINA"`
	TJ int `gDescription:"TAJIKISTAN"`
	TZ int `gDescription:"TANZANIA, UNITED REPUBLIC OF"`
	TH int `gDescription:"THAILAND"`
	TG int `gDescription:"TOGO"`
	TO int `gDescription:"TONGA"`
	TT int `gDescription:"TRINIDAD AND TOBAGO"`
	TN int `gDescription:"TUNISIA"`
	TR int `gDescription:"TURKEY"`
	TM int `gDescription:"TURKMENISTAN"`
	TC int `gDescription:"TURKS AND CAICOS ISLANDS"`
	TV int `gDescription:"TUVALU"`
	UG int `gDescription:"UGANDA"`
	UA int `gDescription:"UKRAINE"`
	AE int `gDescription:"UNITED ARAB EMIRATES"`
	GB int `gDescription:"UNITED KINGDOM"`
	US int `gDescription:"UNITED STATES"`
	UY int `gDescription:"URUGUAY"`
	VU int `gDescription:"VANUATU"`
	VE int `gDescription:"VENEZUELA"`
	VN int `gDescription:"VIETNAM"`
	VG int `gDescription:"VIRGIN ISLANDS, BRITISH"`
	WF int `gDescription:"WALLIS AND FUTUNA"`
	YE int `gDescription:"YEMEN"`
	ZM int `gDescription:"ZAMBIA"`
}

// This should only be invoked when the CountryCode is required.
func (self CountryCodeEnum) validate() error {
	return validateEnum(self)
}

/*
@enum json:"string"
*/
type __StateCode struct {
	AL int `gDescription:"Alabama"`
	AK int `gDescription:"Alaska"`
	AZ int `gDescription:"Arizona"`
	AR int `gDescription:"Arkansas"`
	CA int `gDescription:"California"`
	CO int `gDescription:"Colorado"`
	CT int `gDescription:"Connecticut"`
	DE int `gDescription:"Delaware"`
	FL int `gDescription:"Florida"`
	GA int `gDescription:"Georgia"`
	HI int `gDescription:"Hawaii"`
	ID int `gDescription:"Idaho"`
	IL int `gDescription:"Illinois"`
	IN int `gDescription:"Indiana"`
	IA int `gDescription:"Iowa"`
	KS int `gDescription:"Kansas"`
	KY int `gDescription:"Kentucky"`
	LA int `gDescription:"Louisiana"`
	ME int `gDescription:"Maine"`
	MD int `gDescription:"Maryland"`
	MA int `gDescription:"Massachusetts"`
	MI int `gDescription:"Michigan"`
	MN int `gDescription:"Minnesota"`
	MS int `gDescription:"Mississippi"`
	MO int `gDescription:"Missouri"`
	MT int `gDescription:"Montana"`
	NE int `gDescription:"Nebraska"`
	NV int `gDescription:"Nevada"`
	NH int `gDescription:"New Hampshire"`
	NJ int `gDescription:"New Jersey"`
	NM int `gDescription:"New Mexico"`
	NY int `gDescription:"New York"`
	NC int `gDescription:"North Carolina"`
	ND int `gDescription:"North Dakota"`
	OH int `gDescription:"Ohio"`
	OK int `gDescription:"Oklahoma"`
	OR int `gDescription:"Oregon"`
	PA int `gDescription:"Pennsylvania"`
	RI int `gDescription:"Rhode Island"`
	SC int `gDescription:"South Carolina"`
	SD int `gDescription:"South Dakota"`
	TN int `gDescription:"Tennessee"`
	TX int `gDescription:"Texas"`
	UT int `gDescription:"Utah"`
	VT int `gDescription:"Vermont"`
	VA int `gDescription:"Virginia"`
	WA int `gDescription:"Washington"`
	WV int `gDescription:"West Virginia"`
	WI int `gDescription:"Wisconsin"`
	WY int `gDescription:"Wyoming"`

	AS int `gDescription:"American Samoa"`
	DC int `gDescription:"District of Columbia"`
	FM int `gDescription:"Federated States of Micronesia"`
	GU int `gDescription:"Guam"`
	MH int `gDescription:"Marshall Islands"`
	MP int `gDescription:"Northern Mariana Islands"`
	PW int `gDescription:"Palau"`
	PR int `gDescription:"Puerto Rico"`
}
