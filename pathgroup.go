package gopal

import "fmt"
import "net/http"

type PathGroup struct {
	connection *Connection
	return_url string
	cancel_url string
	Payments Payments
	Authorizations Authorizations
//	Vault
//	Identity
}

type Payments struct {
	pathGroup *PathGroup
	pending map[string]*Payment
}

func (self *Payments) get(req *http.Request) (*Payment, error) {
	var query = req.URL.Query()
	var uuid = query.Get("uuid")
	var pymt, _ = self.pending[uuid]

	if pymt == nil || pymt.uuid != uuid {
		return nil, fmt.Errorf("Unknown payment")
	}
	pymt.url_values = query
	return pymt, nil
}



type Authorizations struct {
	pathGroup *PathGroup
	pending map[string]*Authorization
}
type Authorization struct {
}

func (self *Authorizations) get(req *http.Request) (*Authorization, error) {
	return nil, nil
}
