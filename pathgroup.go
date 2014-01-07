package gopal

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


