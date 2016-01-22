#GoPayPal

GoPayPal is an abstraction over PayPal's REST API for the Go programming language. It aims to provide representation for all data and functionality of the API in a lightweight, safe and fast way.

Because the PayPal REST API is always being enhanced and improved, please file an issue if you find anything missing.

#Quick Example

``` go
package main

import (
  "log"
  p "github.com/Perelandric/GoPayPal"
)

func main() {
  // Establish your connection to PayPal
  var conn, err = p.NewConnection(p.Server.Sandbox, "(my id)", "(my secret)")
  if err != nil {
    log.Fatal(err)
  }

  // Create a payment, providing type of payment method, and redirect urls
  var pymt = conn.CreatePayment(p.PaymentMethod.PayPal,
    p.Redirects {
      Return: "http://www.example.com/page_to_handle_success.html",
      Cancel: "http://www.example.com/page_to_handle_cancelation.html",
    })

  // Create a transaction for the payment. Only the CurrencyType is required.
  var trans = p.NewTransaction(p.CurrencyType.USD,
    "The description of the transaction",
    &p.ShippingAddress {
      RecipientName: "Bob Woo",
      Type: p.AddressType.Residential,
      Address: Address {
      	Line1: "12345 67th Place",
        Line2: "Apt #890",
        City: "Whoville",
        CountryCode: CountryCode.US,
        PostalCode: "90210",
        State: "CA",
        Phone: "555-555-5555",
      },
    })
}
```
