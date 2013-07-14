package gopal

import "fmt"

var AmountMismatchError = fmt.Errorf("Sum of values doesn't match total amount.")

type identity_error struct {
    Error string				`json:"error,omitempty"`
    Error_description string	`json:"error_description,omitempty"`
    Error_uri string			`json:"error_uri,omitempty"`
}
func (ie *identity_error) to_error() error {
	return fmt.Errorf("PayPal error response:%q, Description:%q", ie.Error, ie.Error_description)
}

func (pe *payment_error) to_error() error {
	return fmt.Errorf("Payment error response: %q, Message: %q", pe.Name, pe.Message)
}

type http_status_error struct {
    Name string				`json:"name,omitempty"`
    Message string			`json:"message,omitempty"`
    Information_link string	`json:"information_link,omitempty"`
}

type payment_error struct {
	http_status_error
    Debug_id string			`json:"debug_id,omitempty"`
    Details error_details	`json:"details,omitempty"`
}
type error_details struct {
    Field string	`json:"field,omitempty"`
    Issue string	`json:"issue,omitempty"`
}


