package middleware

import (
	"encoding/json"
	"io"
	"net/http"

	. "github.com/yeqown/gweb/logger"
	. "github.com/yeqown/gweb/utils"
)

// response s to client
func response(w http.ResponseWriter, s string) {
	ReqL.Info(s)
	_, err := io.WriteString(w, s)
	if err != nil {
		ReqL.Errorf("response err: %s", err.Error())
	}
}

// ResponseJson response interface{} as json string to client
func ResponseJson(w http.ResponseWriter, i interface{}) {
	ReqL.Info("ResponseJson ~ ")

	bs, err := json.Marshal(i)
	if err != nil {
		bs, _ = json.Marshal(NewCodeInfo(CodeSystemErr, err.Error()))
		ReqL.Errorf("get an err: %s", err.Error())
	}
	// set ContentType
	w.Header().Set("Content-Type", "application/json")

	response(w, string(bs))
}

// ResponseErrorJson response JsonErr data to client
func ResponseErrorJson(w http.ResponseWriter, je interface{}) {
	bs, err := json.Marshal(je)
	if err != nil {
		ReqL.Errorf("get an err: %s", err.Error())
		ResponseJson(w, NewCodeInfo(CodeSystemErr, err.Error()))
		return
	}
	// set ContentType
	w.Header().Set("Content-Type", "application/json")

	response(w, string(bs))
}
