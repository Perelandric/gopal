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
  var pymt = conn.CreatePaypalPayment(
    p.Payer {...},
    p.Redirects {
      Return: "http://www.example.com/page_to_handle_success.html",
      Cancel: "http://www.example.com/page_to_handle_cancelation.html",
    })

  // Create a transaction for the payment. Only the CurrencyType is required.
  var trans = pymt.AddTransaction(
    p.CurrencyType.USD,
    &p.ShippingAddress {
      RecipientName: "Bob Woo",
      Type: p.AddressType.Residential,
      Address: Address {
      	Line1: "12345 67th Place",
        Line2: "Apt #890",
        City: "Whoville",
        State: "CA",
        PostalCode: "90210",
        CountryCode: CountryCode.US,
        Phone: "555-555-5555",
      },
    })

    // Add optional items to the Transaction
    trans.Description = "A description of the transaction."

  // Add items to the Transaction. Quantities < 1 are ignored. Price is to be
  // given according to the CurrencyType that was provided to NewTransaction().
  trans.AddItem(1, 23.45, "Roast Beast", "987654321")

}
```
