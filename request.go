package gopal

import "net/http"
import "path"
import "io"
import "io/ioutil"
import "fmt"
import "strings"
import "encoding/json"


func (pp *PayPal) make_request(method, subdir string, body interface{}, idempotent_id string, jsn interface{}, auth_req bool) error {
	var err error
	var result []byte
	var req *http.Request
	var resp *http.Response
	var body_reader io.Reader
    var url = "https://api"

	// use sandbox url if requested
    if pp.sandbox {
        url += ".sandbox"
    }
    url = url + ".paypal.com/" + path.Join("v1", subdir)

	if str, ok := body.(string); ok {
		body_reader = strings.NewReader(str)
	} else {
		result, err = json.Marshal(body)
		if err != nil {
			return err
		}
		body_reader = strings.NewReader(string(result))
		result = nil
	}

fmt.Printf("\nSending to PayPal: %s\n%s\n\n", url, body_reader)

    req, err = http.NewRequest(method, url, body_reader)
	if err != nil {
		return err
	}
    req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en_US")

	// Should only be not-authenticated when we are sending credentials for authentication
	if auth_req {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.SetBasicAuth(pp.id, pp.secret)
	} else {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", pp.tokeninfo.auth_token())
		if idempotent_id != "" {
fmt.Println("Sending idempotent_id ", idempotent_id)
			req.Header.Set("PayPal-Request-Id", idempotent_id)
		}
	}

	resp, err = pp.client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

fmt.Printf("\nReceived from PayPal: \n%s\n\n", result)

	return json.Unmarshal(result, jsn)
}


