package gopal

import "fmt"

type errorable interface {
	IsError() bool
	to_error() error
}

var ErrNoResults = fmt.Errorf("No results.")

var UnexpectedResponse = fmt.Errorf("Paypal server returned an unexpected response.")
var AmountMismatchError = fmt.Errorf("Sum of values doesn't match total amount.")

type identity_error struct {
	Error             string `json:"error,omitempty"`
	Error_description string `json:"error_description,omitempty"`
	Error_uri         string `json:"error_uri,omitempty"`

	RawData		[]byte `json:"-"`
}

func (self *identity_error) IsError() bool {
	return self != nil
}
func (ie *identity_error) to_error() error {
	if ie.IsError() {
		return fmt.Errorf("PayPal error response:%q, Description:%q", ie.Error, ie.Error_description)
	}
	return nil
}

type http_status_error struct {
	Name             string `json:"name,omitempty"`
	Message          string `json:"message,omitempty"`
	Information_link string `json:"information_link,omitempty"`

	RawData		[]byte `json:"-"`
}

func (self *http_status_error) IsError() bool {
	return self != nil
}
func (hse *http_status_error) to_error() error {
	if hse.IsError() {
		return fmt.Errorf("Http status error: %q, Message: %q, Info link: %q", hse.Name, hse.Message, hse.Information_link)
	}
	return nil
}

type payment_error struct {
	http_status_error
	Debug_id string        `json:"debug_id,omitempty"`
	Details  error_details `json:"details,omitempty"`

	RawData		[]byte `json:"-"`
}

func (self *payment_error) IsError() bool {
	return self != nil
}
func (pe *payment_error) to_error() error {
	if pe.IsError() {
		return fmt.Errorf("Payment error response: %q, Message: %q", pe.Name, pe.Message)
	}
	return nil
}


type error_details struct {
	Field string `json:"field,omitempty"`
	Issue string `json:"issue,omitempty"`
}
