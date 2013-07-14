package gopal

import "net/http"
import "path"
import "io/ioutil"
import "fmt"
import "strings"


// TODO: What about `PayPal-Request-Id` mentioned in docs for idempotency

func (pp *PayPal) make_request(method, subdir string, body, idempotent_id string, auth_req bool) ([]byte, error) {
	var err error
	var result []byte
	var req *http.Request
	var resp *http.Response
    var url = "https://api"

	// use sandbox url if requested
    if pp.sandbox {
        url += ".sandbox"
    }
    url = url + ".paypal.com/" + path.Join("v1", subdir)

fmt.Printf("\nSending to PayPal: %s\n%s\n\n", url, body)

    req, err = http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
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
        return nil, err
    }
    defer resp.Body.Close()

    result, err = ioutil.ReadAll(resp.Body)

fmt.Printf("\nReceived from PayPal: \n%s\n\n", result)

	return result, err
}


