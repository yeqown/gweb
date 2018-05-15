package middleware

import (
	"encoding/json"
	"net/http"
	"net/url"

	. "github.com/yeqown/gweb/logger"
	"github.com/yeqown/gweb/utils"
)

func logReq(req *http.Request) {
	reqlog := RequestLog(req)
	ReqL.Info(reqlog)
}

// RequestLog to log http.Request values
func RequestLog(req *http.Request) string {
	return utils.Fstring("Path: [%s], Method: [%s], Headers: [%s], Form: [%s]",
		req.URL.Path,
		req.Method,
		headerToString(req.Header),
		valuesToString(req.Form),
	)
}

func headerToString(header http.Header) string {
	bs, _ := json.Marshal(header)
	return string(bs)
}

func valuesToString(values url.Values) string {
	bs, _ := json.Marshal(values)
	return string(bs)
}
