package gopal

import "net/http"
import "net/url"
import "time"
import "fmt"
import "path"
import "io"
import "io/ioutil"
import "strings"
import "strconv"
import "bytes"
import "encoding/json"

func NewConnection(live connection_type_i, id, secret, host string) (*Connection, error) {
	var hosturl, err = url.Parse(host)
	if err != nil {
		return nil, err
	}
	var conn = &Connection{live: live, id: id, secret: secret, hosturl: hosturl, client: http.Client{}}
	conn.Payments.connection = conn
	conn.Sales.connection = conn
	conn.Refunds.connection = conn
	conn.Authorizations.connection = conn
	conn.Captures.connection = conn

	err = conn.authenticate()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

type Connection struct {
	live       connection_type_i
	id, secret string
	hosturl    *url.URL
	client     http.Client
	tokeninfo  tokeninfo
    Payments Payments
    Sales Sales
    Refunds Refunds
    Authorizations Authorizations
	Captures Captures
//  Vault
//  Identity
}

func (self *Connection) authenticate() error {
	// If an error is returned, zero the tokeninfo
	var err error
	var duration time.Duration

	// No need to authenticate if the previous has not yet expired
	if time.Now().Before(self.tokeninfo.expiration) {
		return nil
	}

	defer func() {
		if err != nil {
			self.tokeninfo = tokeninfo{}
		}
	}()

	// (re)authenticate
	err = self.make_request("POST", "/oauth2/token", "grant_type=client_credentials", "", &self.tokeninfo, true)
	if err != nil {
		return err
	}

	// Set the duration to expire 3 minutes early to avoid expiration during a request cycle
	duration = time.Duration(self.tokeninfo.Expires_in)*time.Second - 3*time.Minute
	self.tokeninfo.expiration = time.Now().Add(duration)

	return nil
}

// Authorization response
type tokeninfo struct {
	Scope         string `json:"scope,omitempty"`
	Access_token  string `json:"access_token,omitempty"`
	Refresh_token string `json:"refresh_token,omitempty"`
	Token_type    string `json:"token_type,omitempty"`
	Expires_in    uint   `json:"expires_in,omitempty"`

	// Not sure about this field. Appears in response, but not in documentation.
	App_id string `json:"app_id,omitempty"`

	// Derived fields
	expiration time.Time

	// Handles the case where an error is received instead
	*identity_error
}

func (ti *tokeninfo) auth_token() string {
	return ti.Token_type + " " + ti.Access_token
}

func (pp *Connection) make_request(method, subdir string, body interface{}, idempotent_id string, jsn errorable, auth_req bool) error {
	var err error
	var result []byte
	var req *http.Request
	var resp *http.Response
	var body_reader io.Reader
	var url = "https://api"

	// use sandbox url if requested
	if pp.live == Sandbox {
		url += ".sandbox"
	}
	url = url + ".paypal.com/" + path.Join("v1", subdir)

	switch val := body.(type) {
	case string:
		body_reader = strings.NewReader(val)
	case []byte:
		body_reader = bytes.NewReader(val)
	case nil:
		body_reader = bytes.NewReader(nil)
	default:
		result, err = json.Marshal(val)
		if err != nil {
			return err
		}
		body_reader = bytes.NewReader(result)
		result = nil
	}

	req, err = http.NewRequest(method, url, body_reader)
	if err != nil {
		return err
	}

	if auth_req {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.SetBasicAuth(pp.id, pp.secret)
	} else {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", pp.tokeninfo.auth_token())

		// TODO: The UUID generation needs to be imporoved------v
		req.Header.Set("PayPal-Request-Id", idempotent_id + strconv.FormatInt(time.Now().UnixNano(), 36))
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en_US")

	resp, err = pp.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(result))
	err = json.Unmarshal(result, jsn)
	//var x,y = json.Marshal(jsn)
	//fmt.Printf("=++++++++++++++++++\n%s\n%v\n",x,y)
	if err != nil {
		return err
	}

	err = jsn.to_error()
	if err != nil {
		// Specific management for PayPal response errors
		return err
	}

	return nil
}
