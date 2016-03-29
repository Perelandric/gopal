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

	fmt.Printf("Creating Paypal %q connection.\n", s)

	if err = c.authenticate(); err != nil {
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
		method:    method.Post,
		path:      "/oauth2/token",
		body:      "grant_type=client_credentials",
		response:  &self.tokeninfo,
		isAuthReq: true,
	}); err != nil {
		return err
	}

	fmt.Printf("%#v\n", self.tokeninfo)

	// Set the duration to expire 3 minutes early to avoid expiration during a request cycle
	duration = time.Duration(self.tokeninfo.ExpiresIn)*time.Second - 3*time.Minute
	self.tokeninfo.expiration = time.Now().Add(duration)

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
	/* TODO: How do I know if the authentication is still valid before sending?

	if !reqData.isAuthReq {
		if err := self.authenticate(); err != nil {
			return err
		}
	}
	*/

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
			fmt.Println("sending...", string(result))
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

	resp, err = self.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	result, err = ioutil.ReadAll(resp.Body)

	fmt.Println("received...", string(result))

	if err != nil {
		fmt.Println("Ready Body error")
		return err
	}

	reqData.responseData = result

	// If there was no Body, we can return
	if len(bytes.TrimSpace(result)) == 0 || reqData.response == nil {
		return nil
	}

	if err = json.Unmarshal(result, reqData.response); err != nil {
		fmt.Println("Unmarshaling error")
		return err
	}

	var e = reqData.response.to_error()
	if err != nil {
		log.Printf("Paypal response error: %q\n", e.Error())
	} else {
		fmt.Println("was nil")
	}

	return e
}
