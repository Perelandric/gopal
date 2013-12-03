package gopal

import "fmt"
import "net/http"


type PathGroup struct {
    connection *Connection
    return_url string
    cancel_url string
    pending map[string]*Payment
}

func (self *PathGroup) GetPayment(req *http.Request) (*Payment, error) {
    var query = req.URL.Query()
    var uuid = query.Get("uuid")
    var pymt, _ = self.pending[uuid]

    if pymt == nil || pymt.uuid != uuid {
        return nil, fmt.Errorf("Unknown payment")
    }   
    pymt.url_values = query
    return pymt, nil 
}

