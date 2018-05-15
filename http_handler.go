package gweb

import (
	"net/http"
)

type (
	_404Handler int
	_405Handler int
)

var (
	nfController  *_404Handler // not found handler
	mnaController *_405Handler // method not allowed
)

func (n *_404Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// ci := NewCodeInfo(CodeNotFound, "")
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	// middleware.ResponseJson(w)
	// WriteErrResp(w, http.StatusNotFound, ci.Code, ci.Message)
}

func (m *_405Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// ci := NewCodeInfo(CodeMethodNotAllowed, "")
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	// WriteErrResp(w, http.StatusMethodNotAllowed, ci.Code, ci.Message)
}
