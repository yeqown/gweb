// Package middleware include param, reqlog, response, safe
// related functional modules
// current safe.go mainly parse and valid `Request.Form`
package middleware

import (
	valid "github.com/astaxie/beego/validation"
	"github.com/gorilla/schema"
	. "gweb/logger"
	. "gweb/utils"

	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"sync"
)

// ParamError include Field, Value, Message
type ParamError struct {
	Field   string      `json:"field"`
	Value   interface{} `json:"value"`
	Message string      `json:"message"`
}

type ParamErrors []*ParamError

// String of ParamError, to format as string
func (pe *ParamError) String() string {
	return Fstring("field-[%s], invalid with value-[%s], tip: [%s]", pe.Field, pe.Value, pe.Message)
}

func (pe *ParamError) Error() string {
	return pe.String()
}

// new ParamError with (field, message string, value interface{})
func NewParamError(field, message string, value interface{}) *ParamError {
	return &ParamError{
		Field:   field,
		Value:   value,
		Message: message,
	}
}

// new ParamError from valid.Error
func NewParamErrorFromValidError(ve *valid.Error) *ParamError {
	return &ParamError{
		Field:   ve.Field,
		Value:   ve.Value,
		Message: ve.Message,
	}
}

type Errors []*valid.Error

var poolValid = &sync.Pool{
	New: func() interface{} {
		return &valid.Validation{
			RequiredFirst: true,
		}
	},
}

var decoder = schema.NewDecoder()

// ParseParams, parse params into reqRes from req.Form, and support
// form-data, json-body
// TODO: support parse file
func ParseParams(w http.ResponseWriter, req *http.Request, reqRes interface{}) (errs ParamErrors) {
	switch req.Method {
	case http.MethodGet:
		req.ParseForm()
	case http.MethodPost, http.MethodPut:
		req.ParseMultipartForm(20 << 32)
	default:
		req.ParseForm()
	}
	// log request
	logReq(req)

	if shouldParseJson(reqRes) {
		data, err := getJsonData(req)
		if err != nil {
			errs = append(errs, NewParamError("parse.json", err.Error(), ""))
			return
		}
		if err = json.Unmarshal(data, reqRes); err != nil {
			errs = append(errs, NewParamError("json.unmarshal", err.Error(), ""))
			return
		}
		bs, _ := json.Marshal(reqRes)
		ReqL.Info("pasing json body: " + string(bs))
		goto Valid
	}

	// decode
	if err := decoder.Decode(reqRes, req.Form); err != nil {
		errs = append(errs, NewParamError("decoder", err.Error(), ""))
		return
	}
Valid:
	// valid
	v := poolValid.Get().(*valid.Validation)
	if ok, err := v.Valid(reqRes); err != nil {
		errs = append(errs, NewParamError("validation", err.Error(), ""))
	} else if !ok {
		for _, err := range v.Errors {
			errs = append(errs, NewParamErrorFromValidError(err))
		}
	}
	return
}

// shouldParseJson check `i` has field `JSON`
func shouldParseJson(i interface{}) bool {
	v := reflect.ValueOf(i).Elem()
	// field not ZeroValie means true
	if _, ok := v.Type().FieldByName("JSON"); !ok {
		return false
	}
	return true
}

// getJsonData parse json body from request
func getJsonData(req *http.Request) (body []byte, err error) {
	if body, err = ioutil.ReadAll(req.Body); err != nil {
		return
	}
	if len(string(body)) == 0 {
		err = errors.New("json body is empty")
	}
	return
}
