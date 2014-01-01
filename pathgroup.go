package gopal

import "fmt"
import "net/http"

type PathGroup struct {
	connection *Connection
	return_url string
	cancel_url string
	Payments Payments
	Sales Sales
	Refunds Refunds
	Authorizations Authorizations
//	Vault
//	Identity
}

type Payments struct {
	pathGroup *PathGroup
	pending map[string]*PaymentObject
}

func (self *Payments) get(req *http.Request) (*PaymentObject, error) {
	var query = req.URL.Query()
	var uuid = query.Get("uuid")
	var pymt, _ = self.pending[uuid]

	if pymt == nil || pymt.uuid != uuid {
		return nil, fmt.Errorf("Unknown payment")
	}
	pymt.url_values = query
	return pymt, nil
}

func (self *Payments) Get(payment_id string) (*PaymentObject, error) {
	var pymt = new(PaymentObject)
    var err = self.pathGroup.connection.make_request("POST", "payments/payment/" + payment_id, nil, "", pymt, false)
	if err != nil {
		return nil, err
	}
	return pymt, nil
}
