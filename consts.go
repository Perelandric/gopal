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
@enum Server
	Live
	Sandbox

@enum sortBy --json=string
	CreateTime --string=create_time
	UpdateTime --string=update_time

@enum sortOrder --json=string
	DESC
	ASC

@enum relType --json=string
	Self
	ParentPayment --string=parent_payment
	Execute
	Refund
	ApprovalUrl --string=approval_url
	Suspend
	ReActivate --string=re_activate
	Cancel
	Void
	Authorization
	Capture
	Reauthorize
	Order
	Item
	Batch
	Delete
	Patch
	First
	Last
	Update
	Resend
	Next
	Previous
	Start
	NextPage --string=next_page
	PreviousPage --string=previous_page

@enum method --json=string
	Get --string=GET
	Post --string=POST
	Redirect --string=REDIRECT
	Delete --string=DELETE
	Patch --string=PATCH

@enum payerStatus --json=string
	Verified --string=VERIFIED
	Unverified --string=UNVERIFIED

@enum intent --json=string
	Sale --string=sale
	Authorize --string=authorize
	Order --string=order

@enum FailureReason --json=string
	UnableToCompleteTransaction --string=UNABLE_TO_COMPLETE_TRANSACTION
	InvalidPaymentMethod --string=INVALID_PAYMENT_METHOD
	PayerCannotPay --string=PAYER_CANNOT_PAY
	CannotPayThisPayee --string=CANNOT_PAY_THIS_PAYEE
	RedirectRequired --string=REDIRECT_REQUIRED
	PayeeFilterRestrictions --string=PAYEE_FILTER_RESTRICTIONS

@enum fmfFilter --json=string
	Accept  --string=ACCEPT --description="An ACCEPT filter is triggered only for the TOTAL_PURCHASE_PRICE_MINIMUM filter setting and is returned only in direct credit card payments where payment is accepted."
	Pending --string=PENDING --description="Triggers a PENDING filter action where you need to explicitly accept or deny the transaction."
	Deny  --string=DENY --description="Triggers a DENY action where payment is denied automatically."
	Report --string=REPORT --description="Triggers the Flag testing mode where payment is accepted."

@enum filterId
	MaximumTransactionAmount --string=MAXIMUM_TRANSACTION_AMOUNT --description="basic filter"
	UnconfirmedAddress --string=UNCONFIRMED_ADDRESS --description="basic filter"
	CountryMonitor --string=COUNTRY_MONITOR --description="basic filter"
	AvsNoMatch --string=AVS_NO_MATCH --description="Address Verification Service No Match (advanced filter)"
	AvsPartialMatch --string=AVS_PARTIAL_MATCH --description="Address Verification Service Partial Match (advanced filter)"
	AvsUnavailableOrUnsupported --string=AVS_UNAVAILABLE_OR_UNSUPPORTED --description="Address Verification Service Unavailable or Not Supported (advanced filter)"
	CardSecurityCodeMismatch --string=CARD_SECURITY_CODE_MISMATCH --description="advanced filter"
	BillingOrShippingAddressMismatch --string=BILLING_OR_SHIPPING_ADDRESS_MISMATCH --description="advanced filter"
	RiskyZipCode --string=RISKY_ZIP_CODE --description="high risk lists filter"
	SuspectedFreightForwarderCheck --string=SUSPECTED_FREIGHT_FORWARDER_CHECK --description="high risk lists filter"
	RiskyEmailAddressDomainCheck --string=RISKY_EMAIL_ADDRESS_DOMAIN_CHECK --description="high risk lists filter"
	RiskyBankIdentificationNumberCheck --string=RISKY_BANK_IDENTIFICATION_NUMBER_CHECK --description="high risk lists filter"
	RiskyIpAddressRange --string=RISKY_IP_ADDRESS_RANGE --description="high risk lists filter"
	LargeOrderNumber --string=LARGE_ORDER_NUMBER --description="transaction data filter"
	TotalPurchasePriceMinimum --string=TOTAL_PURCHASE_PRICE_MINIMUM --description="transaction data filter"
	IpAddressVelocity --string=IP_ADDRESS_VELOCITY --description="transaction data filter"
	PaypalFraudModel --string=PAYPAL_FRAUD_MODEL --description="transaction data filter"


@enum normStatus --json=string
	Unknown --string=UNKNOWN
	UnnormalizedUserPreferred --string=UNNORMALIZED_USER_PREFERRED
	Normalized --string=NORMALIZED
	Unnormalized --string=UNNORMALIZED

@enum addressStatus --json=string
	Confirmed --string=CONFIRMED
	Unconfirmed --string=UNCONFIRMED

@enum AddressType --json=string
	Residential --string=residential
	Business --string=business
	Mailbox --string=mailbox

*/

// This should only be invoked when the AddressType is required.
func (self AddressTypeEnum) validate() error {
	return validateEnum(self)
}

/*
@enum PaymentMethod --json=string
	CreditCard --string=credit_card
	PayPal --string=paypal
*/

func (self PaymentMethodEnum) validate() error {
	return validateEnum(self)
}

/*
@enum TaxIdType --json=string
	BrCpf --string=BR_CPF
	BrCnpj --string=BR_CNPJ

@enum CreditCardType --json=string
	Visa --string=visa
	MasterCard --string=mastercard
	Discover --string=discover
	Amex --string=amex

@enum State --json=string
	Created --string=created
	Approved --string=approved
	Canceled --string=canceled
	InProgress --string=in_progress
	Failed --string=failed
	Pending --string=pending
	Completed --string=completed
	Refunded --string=refunded
	PartiallyRefunded --string=partially_refunded
	Expired --string=expired
	Ok --string=ok
	Authorized --string=authorized
	Captured --string=captured
	PartiallyCaptured --string=partially_captured
	Voided --string=voided

@enum reasonCode --json=string
	Chargeback --string=CHARGEBACK
	Guarantee --string=GUARANTEE
	BuyerComplaint --string=BUYER_COMPLAINT
	RefundCode --string=REFUND
	UnconfirmedShippingAddress --string=UNCONFIRMED_SHIPPING_ADDRESS
	EcheckCode --string=ECHECK
	InternationalWithdrawal --string=INTERNATIONAL_WITHDRAWAL
	ReceivingPreferenceMandatesManualAction --string=RECEIVING_PREFERENCE_MANDATES_MANUAL_ACTION
	PaymentReview --string=PAYMENT_REVIEW
	RegulatoryReview --string=REGULATORY_REVIEW
	Unilateral --string=UNILATERAL
	VerificationRequired --string=VERIFICATION_REQUIRED

@enum protectionElig --json=string
	Eligible --string=ELIGIBLE --description="Merchant is protected by PayPal's Seller Protection Policy for Unauthorized. Payments and Item Not Received."
	PartiallyEligible --string=PARTIALLY_ELIGIBLE --description="Merchant is protected by PayPal's Seller Protection Policy for Item Not Received or Unauthorized Payments. Refer to protection_eligibility_type for specifics."
	Ineligibile --string=INELIGIBLE --description="Merchant is not protected under the Seller Protection Policy."

@enum protectionEligType --json=string
	ItemNotReceivedEligible --string=ITEM_NOT_RECEIVED_ELIGIBLE --description="Sellers are protected against claims for items not received."
	UnauthorizedPaymentEligible --string=UNAUTHORIZED_PAYMENT_ELIGIBLE --description="Sellers are protected against claims for unauthorized payments."

@enum paymentMode --json=string
	InstantTransfer --string=INSTANT_TRANSFER
	ManualBankTransfer --string=MANUAL_BANK_TRANSFER
	DelayedTransfer --string=DELAYED_TRANSFER
	Echeck --string=ECHECK

@enum pendingReason --json=string
	PayerShippingUnconfirmed --string=PAYER-SHIPPING-UNCONFIRMED
	MultiCurrency --string=MULTI-CURRENCY
	RiskReview --string=RISK-REVIEW
	RegulatoryReview --string=REGULATORY-REVIEW
	VerificationRequired --string=VERIFICATION-REQUIRED
	OrderPending --string=ORDER
	OtherPending --string=OTHER

@enum CurrencyType --json=string
	AUD --string=AUD --description="Australian dollar"
	BRL --string=BRL --description="Brazilian real**"
	CAD --string=CAD --description="Canadian dollar"
	CZK --string=CZK --description="Czech koruna"
	DKK --string=DKK --description="Danish krone"
	EUR --string=EUR --description="Euro"
	HKD --string=HKD --description="Hong Kong dollar"
	HUF --string=HUF --description="Hungarian forint"
	ILS --string=ILS --description="Israeli new shekel"
	JPY --string=JPY --description="Japanese yen*"
	MYR --string=MYR --description="Malaysian ringgit**"
	MXN --string=MXN --description="Mexican peso"
	TWD --string=TWD --description="New Taiwan dollar*"
	NZD --string=NZD --description="New Zealand dollar"
	NOK --string=NOK --description="Norwegian krone"
	PHP --string=PHP --description="Philippine peso"
	PLN --string=PLN --description="Polish złoty"
	GBP --string=GBP --description="Pound sterling"
	SGD --string=SGD --description="Singapore dollar"
	SEK --string=SEK --description="Swedish krona"
	CHF --string=CHF --description="Swiss franc"
	THB --string=THB --description="Thai baht"
	TRY --string=TRY --description="Turkish lira**"
	USD --string=USD --description="United States dollar"
*/

// This should only be invoked when the AddressType is required.
func (self CurrencyTypeEnum) validate() error {
	return validateEnum(self)
}

/*
@enum CountryCode --json=string
	AL --description="ALBANIA"
	DZ --description="ALGERIA"
	AD --description="ANDORRA"
	AO --description="ANGOLA"
	AI --description="ANGUILLA"
	AG --description="ANTIGUA AND BARBUDA"
	AR --description="ARGENTINA"
	AM --description="ARMENIA"
	AW --description="ARUBA"
	AU --description="AUSTRALIA"
	AT --description="AUSTRIA"
	AZ --description="AZERBAIJAN"
	BS --description="BAHAMAS"
	BH --description="BAHRAIN"
	BB --description="BARBADOS"
	BE --description="BELGIUM"
	BZ --description="BELIZE"
	BJ --description="BENIN"
	BM --description="BERMUDA"
	BT --description="BHUTAN"
	BO --description="BOLIVIA"
	BA --description="BOSNIA-HERZEGOVINA"
	BW --description="BOTSWANA"
	BR --description="BRAZIL"
	BN --description="BRUNEI DARUSSALAM"
	BG --description="BULGARIA"
	BF --description="BURKINA FASO"
	BI --description="BURUNDI"
	KH --description="CAMBODIA"
	CA --description="CANADA"
	CV --description="CAPE VERDE"
	KY --description="CAYMAN ISLANDS"
	TD --description="CHAD"
	CL --description="CHILE"
	CN --description="CHINA (For domestic Chinese bank transactions only)"
	C2 --description="CHINA (For CUP, bank card and cross-border transactions)"
	CO --description="COLOMBIA"
	KM --description="COMOROS"
	CD --description="DEMOCRATIC REPUBLIC OF CONGO"
	CG --description="CONGO"
	CK --description="COOK ISLANDS"
	CR --description="COSTA RICA"
	HR --description="CROATIA"
	CY --description="CYPRUS"
	CZ --description="CZECH REPUBLIC"
	DK --description="DENMARK"
	DJ --description="DJIBOUTI"
	DM --description="DOMINICA"
	DO --description="DOMINICAN REPUBLIC"
	EC --description="ECUADOR"
	EG --description="EGYPT"
	SV --description="EL SALVADOR"
	ER --description="ERITERIA"
	EE --description="ESTONIA"
	ET --description="ETHIOPIA"
	FK --description="FALKLAND ISLANDS (MALVINAS)"
	FJ --description="FIJI"
	FI --description="FINLAND"
	FR --description="FRANCE"
	GF --description="FRENCH GUIANA"
	PF --description="FRENCH POLYNESIA"
	GA --description="GABON"
	GM --description="GAMBIA"
	GE --description="GEORGIA"
	DE --description="GERMANY"
	GI --description="GIBRALTAR"
	GR --description="GREECE"
	GL --description="GREENLAND"
	GD --description="GRENADA"
	GP --description="GUADELOUPE"
	GU --description="GUAM"
	GT --description="GUATEMALA"
	GN --description="GUINEA"
	GW --description="GUINEA BISSAU"
	GY --description="GUYANA"
	VA --description="HOLY SEE (VATICAN CITY STATE)"
	HN --description="HONDURAS"
	HK --description="HONG KONG"
	HU --description="HUNGARY"
	IS --description="ICELAND"
	IN --description="INDIA"
	ID --description="INDONESIA"
	IE --description="IRELAND"
	IL --description="ISRAEL"
	IT --description="ITALY"
	JM --description="JAMAICA"
	JP --description="JAPAN"
	JO --description="JORDAN"
	KZ --description="KAZAKHSTAN"
	KE --description="KENYA"
	KI --description="KIRIBATI"
	KR --description="KOREA, REPUBLIC OF (or SOUTH KOREA)"
	KW --description="KUWAIT"
	KG --description="KYRGYZSTAN"
	LA --description="LAOS"
	LV --description="LATVIA"
	LS --description="LESOTHO"
	LI --description="LIECHTENSTEIN"
	LT --description="LITHUANIA"
	LU --description="LUXEMBOURG"
	MG --description="MADAGASCAR"
	MW --description="MALAWI"
	MY --description="MALAYSIA"
	MV --description="MALDIVES"
	ML --description="MALI"
	MT --description="MALTA"
	MH --description="MARSHALL ISLANDS"
	MQ --description="MARTINIQUE"
	MR --description="MAURITANIA"
	MU --description="MAURITIUS"
	YT --description="MAYOTTE"
	MX --description="MEXICO"
	FM --description="MICRONESIA, FEDERATED STATES OF"
	MN --description="MONGOLIA"
	MS --description="MONTSERRAT"
	MA --description="MOROCCO"
	MZ --description="MOZAMBIQUE"
	NA --description="NAMIBIA"
	NR --description="NAURU"
	NP --description="NEPAL"
	NL --description="NETHERLANDS"
	AN --description="NETHERLANDS ANTILLES"
	NC --description="NEW CALEDONIA"
	NZ --description="NEW ZEALAND"
	NI --description="NICARAGUA"
	NE --description="NIGER"
	NU --description="NIUE"
	NF --description="NORFOLK ISLAND"
	NO --description="NORWAY"
	OM --description="OMAN"
	PW --description="PALAU"
	PA --description="PANAMA"
	PG --description="PAPUA NEW GUINEA"
	PE --description="PERU"
	PH --description="PHILIPPINES"
	PN --description="PITCAIRN"
	PL --description="POLAND"
	PT --description="PORTUGAL"
	QA --description="QATAR"
	RE --description="REUNION"
	RO --description="ROMANIA"
	RU --description="RUSSIAN FEDERATION"
	RW --description="RWANDA"
	SH --description="SAINT HELENA"
	KN --description="SAINT KITTS AND NEVIS"
	LC --description="SAINT LUCIA"
	PM --description="SAINT PIERRE AND MIQUELON"
	VC --description="SAINT VINCENT AND THE GRENADINES"
	WS --description="SAMOA"
	SM --description="SAN MARINO"
	ST --description="SAO TOME AND PRINCIPE"
	SA --description="SAUDI ARABIA"
	SN --description="SENEGAL"
	RS --description="SERBIA"
	SC --description="SEYCHELLES"
	SL --description="SIERRA LEONE"
	SG --description="SINGAPORE"
	SK --description="SLOVAKIA"
	SI --description="SLOVENIA"
	SB --description="SOLOMON ISLANDS"
	SO --description="SOMALIA"
	ZA --description="SOUTH AFRICA"
	ES --description="SPAIN"
	LK --description="SRI LANKA"
	SR --description="SURINAME"
	SJ --description="SVALBARD AND JAN MAYEN"
	SZ --description="SWAZILAND"
	SE --description="SWEDEN"
	CH --description="SWITZERLAND"
	TW --description="TAIWAN, PROVINCE OF CHINA"
	TJ --description="TAJIKISTAN"
	TZ --description="TANZANIA, UNITED REPUBLIC OF"
	TH --description="THAILAND"
	TG --description="TOGO"
	TO --description="TONGA"
	TT --description="TRINIDAD AND TOBAGO"
	TN --description="TUNISIA"
	TR --description="TURKEY"
	TM --description="TURKMENISTAN"
	TC --description="TURKS AND CAICOS ISLANDS"
	TV --description="TUVALU"
	UG --description="UGANDA"
	UA --description="UKRAINE"
	AE --description="UNITED ARAB EMIRATES"
	GB --description="UNITED KINGDOM"
	US --description="UNITED STATES"
	UY --description="URUGUAY"
	VU --description="VANUATU"
	VE --description="VENEZUELA"
	VN --description="VIETNAM"
	VG --description="VIRGIN ISLANDS, BRITISH"
	WF --description="WALLIS AND FUTUNA"
	YE --description="YEMEN"
	ZM --description="ZAMBIA"
*/

// This should only be invoked when the CountryCode is required.
func (self CountryCodeEnum) validate() error {
	return validateEnum(self)
}

/*
@enum StateCode --json=string
	AL --description="Alabama"
	AK --description="Alaska"
	AZ --description="Arizona"
	AR --description="Arkansas"
	CA --description="California"
	CO --description="Colorado"
	CT --description="Connecticut"
	DE --description="Delaware"
	FL --description="Florida"
	GA --description="Georgia"
	HI --description="Hawaii"
	ID --description="Idaho"
	IL --description="Illinois"
	IN --description="Indiana"
	IA --description="Iowa"
	KS --description="Kansas"
	KY --description="Kentucky"
	LA --description="Louisiana"
	ME --description="Maine"
	MD --description="Maryland"
	MA --description="Massachusetts"
	MI --description="Michigan"
	MN --description="Minnesota"
	MS --description="Mississippi"
	MO --description="Missouri"
	MT --description="Montana"
	NE --description="Nebraska"
	NV --description="Nevada"
	NH --description="New Hampshire"
	NJ --description="New Jersey"
	NM --description="New Mexico"
	NY --description="New York"
	NC --description="North Carolina"
	ND --description="North Dakota"
	OH --description="Ohio"
	OK --description="Oklahoma"
	OR --description="Oregon"
	PA --description="Pennsylvania"
	RI --description="Rhode Island"
	SC --description="South Carolina"
	SD --description="South Dakota"
	TN --description="Tennessee"
	TX --description="Texas"
	UT --description="Utah"
	VT --description="Vermont"
	VA --description="Virginia"
	WA --description="Washington"
	WV --description="West Virginia"
	WI --description="Wisconsin"
	WY --description="Wyoming"

	AS --description="American Samoa"
	DC --description="District of Columbia"
	FM --description="Federated States of Micronesia"
	GU --description="Guam"
	MH --description="Marshall Islands"
	MP --description="Northern Mariana Islands"
	PW --description="Palau"
	PR --description="Puerto Rico"
*/
