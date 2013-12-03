package gopal

import "fmt"
import "net/http"
import "path"
import "io"
import "io/ioutil"
import "strings"
import "bytes"
import "encoding/json"


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
		if idempotent_id != "" {
			req.Header.Set("PayPal-Request-Id", idempotent_id)
		}
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


