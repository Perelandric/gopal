package gopal

type payer struct {
	Payment_method      PaymentMethod        `json:"payment_method,omitempty"`
	Funding_instruments []funding_instrument `json:"funding_instruments,omitempty"`
	Payer_info          *payer_info          `json:"payer_info,omitempty"`
}
type payer_info struct {
	Email            string           `json:"email,omitempty"`
	First_name       string           `json:"first_name,omitempty"`
	Last_name        string           `json:"last_name,omitempty"`
	Payer_id         string           `json:"payer_id,omitempty"`
	Phone            string           `json:"phone,omitempty"`
	Shipping_address *ShippingAddress `json:"shipping_address,omitempty"`
}
type Address struct {
	Line1        string `json:"line1,omitempty"`
	Line2        string `json:"line2,omitempty"`
	City         string `json:"city,omitempty"`
	Country_code string `json:"country_code,omitempty"`
	Postal_code  string `json:"postal_code,omitempty"`
	State        string `json:"state,omitempty"`
	Phone        string `json:"phone,omitempty"`
}
type ShippingAddress struct {
	Recipient_name string      `json:"recipient_name,omitempty"`
	Type           AddressType `json:"type,omitempty"`
	Address
}

type identity_address struct {
	Street_address string `json:"street_address,omitempty"`
	Locality       string `json:"locality,omitempty"`
	Region         string `json:"region,omitempty"`
	Postal_code    string `json:"postal_code,omitempty"`
	Country        string `json:"country,omitempty"`
}

type userinfo struct {
	User_id     string `json:"user_id,omitempty"`
	Sub         string `json:"sub,omitempty"`
	Name        string `json:"name,omitempty"`
	Given_name  string `json:"given_name,omitempty"`
	Family_name string `json:"family_name,omitempty"`
	Middle_name string `json:"middle_name,omitempty"`

	Email          string `json:"email,omitempty"`
	Email_verified string `json:"email_verified,omitempty"`
	Phone_number   string `json:"phone_number,omitempty"`

	Address  identity_address `json:"address,omitempty"`
	Zoneinfo string           `json:"zoneinfo,omitempty"`
	Locale   string           `json:"locale,omitempty"`

	Gender    string `json:"gender,omitempty"`
	Birthdate string `json:"birthdate,omitempty"`
	Age_range string `json:"age_range,omitempty"`
	Picture   string `json:"picture,omitempty"`

	Verified_account bool   `json:"verified_account,omitempty"`
	Account_type     string `json:"account_type,omitempty"`
	Payer_id         string `json:"payer_id,omitempty"`
}

type funding_instrument struct {
	Credit_card       *credit_card       `json:"credit_card,omitempty"`
	Credit_card_token *credit_card_token `json:"credit_card_token,omitempty"`
}
type _cc_details struct {
	Payer_id     string         `json:"payer_id,omitempty"`
	Type         CreditCardType `json:"type,omitempty"`
	Expire_month string         `json:"expire_month,omitempty"`
	Expire_year  string         `json:"expire_year,omitempty"`
}
type credit_card struct {
	Id     string `json:"id,omitempty"`
	Number string `json:"number,omitempty"`
	_cc_details
	Cvv2            string   `json:"cvv2,omitempty"`
	First_name      string   `json:"first_name,omitempty"`
	Last_name       string   `json:"last_name,omitempty"`
	Billing_address *Address `json:"billing_address,omitempty"`
	State           State    `json:"state,omitempty"`
	Valid_until     string   `json:"valid_until,omitempty"`
}
type credit_card_token struct {
	Credit_card_id string `json:"credit_card_id,omitempty"`
	Last4          string `json:"last4,omitempty"`
	_cc_details
}
