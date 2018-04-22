package middleware

import (
	"encoding/json"
	. "gweb/constant"
	. "gweb/logger"
	"io"
	"net/http"
)

func response(w http.ResponseWriter, s string) {
	ReqL.Info(s)
	_, err := io.WriteString(w, s)
	if err != nil {
		ReqL.Errorf("response err: %s", err.Error())
	}
}

func ResponseJson(w http.ResponseWriter, i interface{}) {
	bs, err := json.Marshal(i)
	if err != nil {
		bs, _ = json.Marshal(NewCodeInfo(CodeSystemErr, err.Error()))
		ReqL.Errorf("get an err: %s", err.Error())
	}
	response(w, string(bs))
}

type JsonErr struct {
	CodeInfo
	Errs interface{} `json:"errs"`
}

func ResponseErrorJson(w http.ResponseWriter, je *JsonErr) {
	bs, err := json.Marshal(je)
	if err != nil {
		ReqL.Errorf("get an err: %s", err.Error())
		ResponseJson(w, NewCodeInfo(CodeSystemErr, err.Error()))
		return
	}
	response(w, string(bs))
}
