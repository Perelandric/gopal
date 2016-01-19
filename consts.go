package gopal

//go:generate GoEnum $GOFILE

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

/*
@enum
--name=sortBy --marshaler=string --unmarshaler=string
CreateTime --string=create_time
UpdateTime --string=update_time

@enum
--name=sortOrder
DESC
ASC

@enum
--name=relType
self
parentPayment --string=parent_payment
execute
refund
approvalUrl --string=approval_url
suspend
reActivate --string=re_activate
cancel
void
authorization
capture
reauthorize
order
item
batch
delete
patch
first
last
update
resend
next
previous
start
nextPage --string=next_page
previousPage --string=previous_page

@enum
--name=method
get --string=GET
post --string=POST
redirect --string=REDIRECT
delete --string=DELETE
patch --string=PATCH

@enum
--name=payerStatus
verified --string=VERIFIED
unverified --string=UNVERIFIED

@enum
--name=intent
sale
authorize
order

@enum
--name=normStatus
unknown --string=UNKNOWN
unnormalizedUserPreferred --string=UNNORMALIZED_USER_PREFERRED
normalized --string=NORMALIZED
unnormalized --string=UNNORMALIZED

@enum
--name=addressStatus
confirmed --string=CONFIRMED
unconfirmed --string=UNCONFIRMED

@enum
--name=AddressType
Residential --string=residential
Business --string=business
Mailbox --string=mailbox

@enum
--name=PaymentMethod
CreditCard --string=credit_card
PayPal --string=paypal

@enum
--name=TaxIdType
BrCpf --string=BR_CPF
BrCnpj --string=BR_CNPJ

@enum
--name=CreditCardType
Visa --string=visa
MasterCard --string=mastercard
Discover --string=discover
Amex --string=amex

@enum
--name=state
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

@enum
--name=reasonCode
chargeback --string=CHARGEBACK
guarantee --string=GUARANTEE
buyerComplaint --string=BUYER_COMPLAINT
refundCode --string=REFUND
unconfirmedShippingAddress --string=UNCONFIRMED_SHIPPING_ADDRESS
echeckCode --string=ECHECK
internationalWithdrawal --string=INTERNATIONAL_WITHDRAWAL
receivingPreferenceMandatesManualAction --string=RECEIVING_PREFERENCE_MANDATES_MANUAL_ACTION
paymentReview --string=PAYMENT_REVIEW
regulatoryReview --string=REGULATORY_REVIEW
unilateral --string=UNILATERAL
verificationRequired --string=VERIFICATION_REQUIRED

@enum
--name=protectionElig
eligible --string=ELIGIBLE --description="Merchant is protected by PayPal's Seller Protection Policy for Unauthorized. Payments and Item Not Received."
partiallyEligible --string=PARTIALLY_ELIGIBLE --description="Merchant is protected by PayPal's Seller Protection Policy for Item Not Received or Unauthorized Payments. Refer to protection_eligibility_type for specifics."
ineligibile --string=INELIGIBLE --description="Merchant is not protected under the Seller Protection Policy."

@enum
--name=protectionEligType
itemNotReceivedEligible --string=ITEM_NOT_RECEIVED_ELIGIBLE --description="Sellers are protected against claims for items not received."
unauthorizedPaymentEligible --string=UNAUTHORIZED_PAYMENT_ELIGIBLE --description="Sellers are protected against claims for unauthorized payments."

@enum
--name=paymentMode
instantTransfer --string=INSTANT_TRANSFER
manualBankTransfer --string=MANUAL_BANK_TRANSFER
delayedTransfer --string=DELAYED_TRANSFER
echeck --string=ECHECK

@enum
--name=pendingReason
payerShippingUnconfirmed --string=PAYER-SHIPPING-UNCONFIRMED
multiCurrency --string=MULTI-CURRENCY
riskReview --string=RISK-REVIEW
regulatoryReview --string=REGULATORY-REVIEW
verificationRequired --string=VERIFICATION-REQUIRED
orderPending --string=ORDER
otherPending --string=OTHER

@enum
--name=CurrencyType
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

@enum
--name=CountryCode
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
