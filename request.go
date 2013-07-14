package gopal

import "net/http"
import "path"
import "io"
import "io/ioutil"
import "fmt"


// TODO: What about `PayPal-Request-Id` mentioned in docs for idempotency

func (pp *PayPal) make_request(method, subdir string, body io.Reader, auth_req bool) ([]byte, error) {
	var err error
	var req *http.Request
	var resp *http.Response
    var url = "https://api"

	// use sandbox url if requested
    if pp.sandbox {
        url += ".sandbox"
    }
    url = url + ".paypal.com/" + path.Join("v1", subdir)

fmt.Println(url)
    req, err = http.NewRequest(method, url, body)
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
	}

	resp, err = pp.client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    return ioutil.ReadAll(resp.Body)
}


