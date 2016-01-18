package gopal

// http://play.golang.org/p/zUgqXPsjjg // Idea for generated enums

import "bytes"
import "encoding/json"
import "io"
import "io/ioutil"
import "net/http"

import "path"

import "strings"
import "strconv"
import "time"

type request struct {
	method    method
	path      string
	body      interface{}
	response  errorable
	isAuthReq bool
}

type connection struct {
	live       bool
	id, secret string
	tokeninfo  tokeninfo
	client     http.Client
}

func NewConnection(live bool, id, secret string) (conn *connection, err error) {
	conn = &connection{
		live:   live,
		id:     id,
		secret: secret,
	}

	if err = conn.authenticate(); err != nil {
		return nil, err
	}
	return
}

func (self *connection) authenticate() error {
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
	if err = self.send(&request{
		method:    post,
		path:      "/oauth2/token",
		body:      "grant_type=client_credentials",
		response:  &self.tokeninfo,
		isAuthReq: true,
	}); err != nil {
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

	RawData []byte `json:"-"`

	// Handles the case where an error is received instead
	*identity_error
}

func (ti *tokeninfo) auth_token() string {
	return ti.Token_type + " " + ti.Access_token
}

func (pp *connection) send(reqData *request) error {
	var err error
	var result []byte
	var req *http.Request
	var resp *http.Response
	var body_reader io.Reader
	var url string

	// use sandbox url if requested
	if pp.live {
		url = path.Join("https://api.paypal.com/v1", reqData.path)
	} else {
		url = path.Join("https://api.sandbox.paypal.com/v1", reqData.path)
	}

	switch val := reqData.body.(type) {
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

	// TODO: Paypal docs mention a `nonce`. Research that.

	req, err = http.NewRequest(string(reqData.method), url, body_reader)
	if err != nil {
		return err
	}

	if reqData.isAuthReq {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.SetBasicAuth(pp.id, pp.secret)
	} else {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", pp.tokeninfo.auth_token())

		// TODO: How to include idempotent id
		// TODO: The UUID generation needs to be improved------v
		req.Header.Set(
			"PayPal-Request-Id", strconv.FormatInt(time.Now().UnixNano(), 36))
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

	/*
		fmt.Println("RESPONSE:",string(result))
	*/

	// If there was no Body, we can return
	if len(bytes.TrimSpace(result)) == 0 || reqData.response == nil {
		return nil
	}

	if err = json.Unmarshal(result, reqData.response); err != nil {
		return err
	}

	return reqData.response.to_error()
}
