package middleware

import (
	. "github.com/yeqown/gweb/logger"
	"net/http"
	"runtime/debug"
)

// SafeHandler recover from panic
// will be used in `Handler.ServeHTTP`
func SafeHandler(w http.ResponseWriter, req *http.Request) {
	if err, ok := recover().(error); ok {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		ReqL.Errorf("panic: %s", err.Error())
		ReqL.Error(string(debug.Stack()))
	}
}
