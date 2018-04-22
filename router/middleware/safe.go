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
