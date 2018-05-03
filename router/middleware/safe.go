// Package middleware include param, reqlog, response, safe
// related functional modules
// current safe.go mainly recover while panic err hanppening
package middleware

import (
	. "gweb/logger"
	"net/http"
)

func SafeHandler(w http.ResponseWriter, req *http.Request) {
	if err, ok := recover().(error); ok {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		ReqL.Errorf("panic: %s", err.Error())
	}
}
