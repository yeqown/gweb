package utils

import (
	"bytes"
	"github.com/gorilla/schema"
	"io/ioutil"
	"net/http"
)

var decoder = schema.NewDecoder()

func init() {
	decoder.IgnoreUnknownKeys(true)
}

// copy a new http.Request
func CopyRequest(req *http.Request) *http.Request {
	body, _ := ioutil.ReadAll(req.Body)
	rd_only := ioutil.NopCloser(bytes.NewBuffer(body))

	new_req, err := http.NewRequest(req.Method, req.URL.String(), bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	new_req.Header = req.Header
	req.Body = rd_only
	return new_req
}

// parse Request from request
func ParseRequest(req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		req.ParseForm()
	case http.MethodPost, http.MethodPut:
		req.ParseMultipartForm(20 << 32)
	default:
		req.ParseForm()
	}
}
