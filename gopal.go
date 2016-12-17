package gopal

// http://play.golang.org/p/zUgqXPsjjg // Idea for generated enums

import (
	"Golific/gJson"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"
)

type connection struct {
	server     ServerEnum
	id, secret string
	tokeninfo  tokeninfo
	client     http.Client

	mux            sync.RWMutex
	reauthAttempts int
	expiration     time.Time
	tempBlock      bool
}

type Connection interface {
	FetchPayment(string) (Payment, error)
	GetAllPayments(int, sortByEnum, sortOrderEnum, ...time.Time) *PaymentBatcher
	FetchAuthorization(string) (*Authorization, error)
	FetchCapture(string) (*Capture, error)
	NewCreditCardPayment() *CreditCardPayment
	NewPaypalPayment(Redirects, *PaypalPayerInfo) (*PaypalPayment, error)
	Execute(*url.URL) error
	FetchRefund(string) (*Refund, error)
	FetchSale(string) (*Sale, error)

	authenticate() error
	send(*request) error
}

// Authorization response
type tokeninfo struct {
	// The access token issued by the authorization server.
	AccessToken string `json:"access_token,omitempty"`

	// The refresh token, which can be used to obtain new access tokens using the
	// same authorization grant as described in OAuth2.0 RFC6749 - Section 6.
	RefreshToken string `json:"refresh_token,omitempty"`

	// The type of the token issued as described in OAuth2.0 RFC6749 - Section 7.1
	// Value is case insensitive.
	TokenType string `json:"token_type,omitempty"`

	// The lifetime of the access token in seconds. After the access token
	// expires, use the refresh_token to refresh the access token.
	ExpiresIn uint `json:"expires_in,omitempty"`

	// Not sure what this is for. Appears once in a response example, but no
	// explanation is given anywhere.
	AppId string `json:"app_id,omitempty"`

	// THERE SEEM TO BE DIFFERENT DOCS IN DIFFERENT PLACES FOR THIS!!!
	//
	// Required if different from the scope requested by the client. For a list of
	// possible values, see the attributes table.
	// https://developer.paypal.com/docs/integration/direct/identity/attributes/
	//
	// Scopes expressed in the form of resource URL endpoints. The value of the
	// scope parameter is expressed as a list of space-delimited, case-sensitive
	// strings.
	// Value assigned by PayPal.
	Scope string `json:"scope,omitempty"`

	// Handles the case where an error is received instead
	*identity_error
}

func (c *connection) auth_token() string {
	c.mux.RLock()
	defer c.mux.RUnlock()

	return c.tokeninfo.TokenType + " " + c.tokeninfo.AccessToken
}

func NewConnection(s ServerEnum, id, secret string) (c *connection, err error) {
	c = &connection{
		server: s,
		id:     id,
		secret: secret,
	}

	log.Printf("\nCreating Paypal %q connection.\n", s)

	if err = c.authenticate(); err != nil {
		return nil, err
	}
	return
}

var authTimeout = fmt.Errorf(
	`Temporarily unable to connect to the payment server. Please try again shortly.`,
)

// From the docs, regarding [re]authentication:
//
// PayPal-issued access tokens can be used to access all the REST API endpoints.
// These tokens have a finite lifetime and you must write code to detect when an
// access token expires. You can do this either by keeping track of the
// `expires_in` value returned in the response from the token request (the value
// is expressed in seconds), or handle the error response (401 Unauthorized)
// from the API endpoint when an expired token is detected.
func (c *connection) authenticate() (err error) {
	c.mux.RLock()
	if c.tempBlock {
		c.mux.RUnlock()
		return authTimeout
	}
	var valid = time.Now().Before(c.expiration)
	c.mux.RUnlock()

	if valid {
		return nil
	}

	// Authentication expired
	c.mux.Lock()
	defer c.mux.Unlock()

	if c.tempBlock {
		// A request got in first and failed auth, so we're blocked for a moment
		return authTimeout
	}

	if time.Now().Before(c.expiration) {
		return nil // A request got in first and succeeded auth, so we're good
	}

	// Try to authenticate
	if c.reauthAttempts == 0 {
		log.Println("Attempting authentication...")
	}

	var n = float64(c.reauthAttempts)
	var limit = n + 3

	for { // (re)authenticate
		var req *http.Request
		var resp *http.Response

		var r = &request{
			method:   method.Post,
			path:     "/oauth2/token",
			body:     "grant_type=client_credentials",
			response: &c.tokeninfo,
		}

		if req, err = r.getHttpRequest(c.server); err != nil {
			goto FAIL
		}
		req.SetBasicAuth(c.id, c.secret)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		if resp, err = c.client.Do(req); err != nil {
			goto FAIL
		}

		defer resp.Body.Close()

		if err = r.processBody(resp); err == nil {
			var expiresIn = time.Duration(c.tokeninfo.ExpiresIn) * time.Second

			// Set to expire 2 minutes early to avoid expiration during a request cycle
			c.expiration = time.Now().Add(expiresIn - (2 * time.Minute))
			c.reauthAttempts = 0
			c.tempBlock = false

			log.Println("Authentication succeeded.")
			return nil
		}

	FAIL:
		n = n + 1

		if math.Mod(n, math.Pow(10, math.Floor(math.Log10(n)))) == 0 {
			// every 10, then every 100, then every 1000, etc...
			log.Printf("Authentication failed %d attempts.\n", n)
		}

		if n >= limit {
			break
		}

		time.Sleep(time.Second) // Pause before retry
	}

	// The incremented attempts all failed, so give up and return the error.
	c.reauthAttempts = int(n)
	c.tokeninfo = tokeninfo{}

	// Prevent buildups of waiting requests by blocking requests for a short time.
	c.tempBlock = true
	go func(c *connection) {
		<-time.After(time.Second)

		c.mux.Lock()
		c.tempBlock = false
		c.mux.Unlock()
	}(c)

	return err
}

/*

TODO:

We use a `mux` above but not below. This is probably safe because most of the
stuff that would be written while locked above is not the same stuff as below.
In the cases where we would need to read something below that could be written
above, I think it would only be read if the authentication passed, meaning it
would not be writing above.

But what if auth passed for a request, but then a nanosecond later we expired and
the auth failed for the next request, which then attempted authentication above.
If that authentication took place while the first reqest was still working below,
is it possible that something below on `c` could be being read or written?

I think that the `auth_token` call below would be the only thing that could conflict
in that case because it reads the `tokeninfo` data that may be being set above.

One possibility would be to have that`auth_token()` function do a RLock() while
reading that data. I would just need to make sure that there couldn't be a
deadlock.

Another possibility would be to not reuse the tokeninfo, so that when a new
authentication is taking place, the first request that made it through has the
old `tokeninfo`. But is that right? I mean the `connection` would need to write
the whole new object, which would have the same race condition possibility.

Does it make any sense to put all this authentication stuff on a separate object
that's held by the connection? I don't think that really changes anything, but
it may make it clearer that the only "auth" stuff referenced below is what is
given by the `auth_token()` call.
*/

func (c *connection) send(reqData *request) (err error) {
	var req *http.Request
	var resp *http.Response

	// Make sure we're still authenticated. Will refresh if not.
	if err = c.authenticate(); err != nil {
		return err
	}

	if req, err = reqData.getHttpRequest(c.server); err != nil {
		return err
	}

	// TODO: How to include idempotent id
	// TODO: The UUID generation needs to be improved------v
	req.Header.Set(
		"PayPal-Request-Id", strconv.FormatInt(time.Now().UnixNano(), 36))
	req.Header.Set("Authorization", c.auth_token())
	req.Header.Set("Content-Type", "application/json")

	if resp, err = c.client.Do(req); err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		// Unauthorized, yet made it past the `authenticate()` check above.
		log.Println("Unexpected 401 response. Attempting re-auth.")

		if err = c.authenticate(); err != nil {
			log.Printf("Re-auth after 401 resulted in err: %s\n", err)
			return err
		}

		log.Println("Re-auth succeeded. Attempting recursive send()")
		return c.send(reqData)
	}

	return reqData.processBody(resp)
}

type request struct {
	method       methodEnum
	path         string
	body         interface{}
	response     errorable
	responseData []byte
}

func (r *request) getHttpRequest(s ServerEnum) (*http.Request, error) {
	var body_reader io.Reader

	switch val := r.body.(type) {
	case string:

		//		fmt.Printf("req:\n%s\n\n", val)

		body_reader = strings.NewReader(val)

	case []byte:

		//		fmt.Printf("req:\n%s\n\n", val)

		body_reader = bytes.NewReader(val)

	case nil:
		body_reader = bytes.NewReader(nil)

	case gJson.JSONEncodable:
		var enc = gJson.Encoder{}
		val.JSONEncode(&enc)
		body_reader = bytes.NewReader(enc.Bytes())
		//		fmt.Printf("req:\n%s\n\n", enc.Bytes())

	default:
		if result, err := json.Marshal(val); err != nil {
			return nil, err
		} else {
			//			fmt.Printf("req:\n%s\n\n", result)
			body_reader = bytes.NewReader(result)
		}
	}

	// TODO: Paypal docs mention a `nonce`. Research that.

	var url = "https://" // Can't include this when doing the `.Join()` below

	// Use sandbox url if requested
	if s == Server.Sandbox {
		url += path.Join("api.sandbox.paypal.com/v1", r.path)
	} else {
		url += path.Join("api.paypal.com/v1", r.path)
	}

	req, err := http.NewRequest(r.method.String(), url, body_reader)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en_US")

	return req, nil
}

func (r *request) processBody(resp *http.Response) error {
	if r.response == nil {
		return nil
	}

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	r.responseData = bytes.TrimSpace(result)

	//	fmt.Printf("response:\n%s\n\n", result)

	// If there was no Body, we can return
	if len(r.responseData) == 0 {
		return nil
	}

	if err = json.Unmarshal(result, r.response); err != nil {
		return err
	}

	return r.response.to_error()
}
