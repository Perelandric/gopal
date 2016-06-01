package gopal

// http://play.golang.org/p/zUgqXPsjjg // Idea for generated enums

import (
	"bytes"
	"fmt"
	"log"
)

import "encoding/json"
import "io"
import "io/ioutil"
import "net/http"

import "path"

import "strings"
import "strconv"
import "time"

func NewConnection(s ServerEnum, id, secret string) (c *connection, err error) {
	c = &connection{
		server: s,
		id:     id,
		secret: secret,
	}

	log.Println()
	log.Printf("Creating Paypal %q connection.\n", s)

	if err = c.authenticate(); err != nil {
		return nil, err
	}
	return
}

// From the docs, regarding [re]authentication:
//
// PayPal-issued access tokens can be used to access all the REST API endpoints.
// These tokens have a finite lifetime and you must write code to detect when an
// access token expires. You can do this either by keeping track of the
// `expires_in` value returned in the response from the token request (the value
// is expressed in seconds), or handle the error response (401 Unauthorized)
// from the API endpoint when an expired token is detected.
func (self *connection) authenticate() error {
	// If an error is returned, zero the tokeninfo
	var err error
	var duration time.Duration

	self.tokeninfo.mux.Lock()
	defer self.tokeninfo.mux.Unlock()

	// Check this here as well, since the current auth request could have been
	// waiting for another. If it succeeded, no sense in sending it again.
	if time.Now().Before(self.tokeninfo.expiration) {
		return nil // Another request must have updated the tokeninfo.
	}

	// Don't log each re-attempt after failure
	if self.tokeninfo.reauthAttempts == 0 {
		log.Println("Attempting authentication...")
	}

	// (re)authenticate
	err = self.send(&request{
		method:    method.Post,
		path:      "/oauth2/token",
		body:      "grant_type=client_credentials",
		response:  &self.tokeninfo,
		isAuthReq: true,
	})

	if err != nil {
		self.tokeninfo.reauthAttempts += 1

		if self.tokeninfo.reauthAttempts == 1 {
			// The first client to have an auth failure on this conn.
			log.Println("Authentication failure")

			for self.tokeninfo.reauthAttempts < 3 { // Retry in 5 second increments.
				time.Sleep(5 * time.Second)

				// (re)authenticate
				err = self.send(&request{
					method:    method.Post,
					path:      "/oauth2/token",
					body:      "grant_type=client_credentials",
					response:  &self.tokeninfo,
					isAuthReq: true,
				})

				if err != nil {
					self.tokeninfo.reauthAttempts += 1

				} else { // Success! Skip ahead to set new duration.
					goto OK
				}
			}

			// The incremented attempts all failed, so give up and return the error.
			self.tokeninfo = tokeninfo{
				reauthAttempts: self.tokeninfo.reauthAttempts,
			}
		}

		var attempts = self.tokeninfo.reauthAttempts
		var logEvery = uint(1)

		for attempts > 10 {
			logEvery *= 10
			attempts /= 10
		}

		if self.tokeninfo.reauthAttempts%logEvery == 0 {
			log.Printf(
				"Authentication failed %d consecutive attempts.\n",
				self.tokeninfo.reauthAttempts,
			)
		}

		return err
	}

OK: // Set to expire 3 minutes early to avoid expiration during a request cycle
	duration = time.Duration(self.tokeninfo.ExpiresIn)*time.Second - 3*time.Minute
	self.tokeninfo.expiration = time.Now().Add(duration)
	self.tokeninfo.reauthAttempts = 0

	log.Println("Authentication succeeded.")

	return nil
}

func (self *tokeninfo) auth_token() string {
	return self.TokenType + " " + self.AccessToken
}

func (self *connection) send(reqData *request) error {
	var err error
	var result []byte
	var req *http.Request
	var resp *http.Response
	var body_reader io.Reader
	var url = "https://" // can't include this when doing the `.Join()` below

	// Make sure we're still authenticated. Will refresh if not.
	if !reqData.isAuthReq {
		self.tokeninfo.mux.RLock()
		var isExpired = time.Now().After(self.tokeninfo.expiration)
		self.tokeninfo.mux.RUnlock()

		if isExpired {
			if err := self.authenticate(); err != nil {
				return err
			}
		}
	}

	// use sandbox url if requested
	if self.server == Server.Sandbox {
		url += path.Join("api.sandbox.paypal.com/v1", reqData.path)
	} else {
		url += path.Join("api.paypal.com/v1", reqData.path)
	}

	switch val := reqData.body.(type) {
	case string:
		body_reader = strings.NewReader(val)

	case []byte:
		body_reader = bytes.NewReader(val)

	case nil:
		body_reader = bytes.NewReader(nil)

	default:
		if result, err := json.Marshal(val); err != nil {
			return err
		} else {
			//fmt.Println("sending...", string(result))
			body_reader = bytes.NewReader(result)
		}
	}

	// TODO: Paypal docs mention a `nonce`. Research that.

	req, err = http.NewRequest(reqData.method.String(), url, body_reader)
	if err != nil {
		return err
	}

	if reqData.isAuthReq {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.SetBasicAuth(self.id, self.secret)

	} else {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", self.tokeninfo.auth_token())

		// TODO: How to include idempotent id
		// TODO: The UUID generation needs to be improved------v
		req.Header.Set(
			"PayPal-Request-Id", strconv.FormatInt(time.Now().UnixNano(), 36))
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en_US")

	if resp, err = self.client.Do(req); err != nil {
		return err
	}
	defer resp.Body.Close()

	//fmt.Println("received...", string(result))

	if resp.StatusCode == 401 { // Unauthorized (probably needs token reauth)
		if reqData.isAuthReq { // Unlikely... the auth request received a 401
			result, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				return fmt.Errorf("Auth request received 401: %s", err)
			}
			return fmt.Errorf("Auth request received 401: %s", result)
		}

		// Expired, yet somehow made it past the `expiration` check above.
		log.Println("Unexpected 401 response. Attempting re-auth.")
		if err = self.authenticate(); err != nil {
			log.Printf("Re-auth after 401 resulted in err: %s\n", err)
			return err
		}

		log.Println("Re-auth succeeded. Attempting recursive send()")
		return self.send(reqData)
	}

	if reqData.response == nil {
		return nil
	}

	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	reqData.responseData = bytes.TrimSpace(result)

	// If there was no Body, we can return
	if len(reqData.responseData) == 0 {
		return nil
	}

	if err = json.Unmarshal(result, reqData.response); err != nil {
		return err
	}

	return reqData.response.to_error()
}
