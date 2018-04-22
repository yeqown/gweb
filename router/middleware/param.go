package middleware

import (
	valid "github.com/astaxie/beego/validation"
	"github.com/gorilla/schema"
	. "gweb/utils"

	"errors"
	"io/ioutil"
	"net/http"
	"sync"
)

type ParamError struct {
	Field   string      `json:"filed"`
	Value   interface{} `json:"value"`
	Message string      `json:"message"`
}

type ParamErrors []*ParamError

func (pe *ParamError) String() string {
	return Fstring("filed-[%s], invalid with value-[%s], tip: [%s]", pe.Field, pe.Value, pe.Message)
}

func (pe *ParamError) Error() string {
	return pe.String()
}

func NewParamError(field, message string, value interface{}) *ParamError {
	return &ParamError{
		Field:   field,
		Value:   value,
		Message: message,
	}
}

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

func ParseParams(w http.ResponseWriter, req *http.Request, reqRes interface{}) (errs ParamErrors) {
	switch req.Method {
	case http.MethodGet:
		req.ParseForm()
	case http.MethodPost, http.MethodPut:
		req.ParseMultipartForm(20 << 32)
	default:
		req.ParseForm()
	}

	// decode
	if err := decoder.Decode(reqRes, req.Form); err != nil {
		errs = append(errs, NewParamError("decoder", err.Error(), ""))
		return
	}
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

// parse json body
func GetJsonData(req *http.Request) (body []byte, err error) {
	if body, err = ioutil.ReadAll(req.Body); err != nil {
		return
	}
	if len(string(body)) == 0 {
		err = errors.New("json body is empty")
	}
	return
}
