// Package middleware include param, reqlog, response, safe
// related functional modules
package middleware

import (
	"encoding/json"
	"errors"
	valid "github.com/astaxie/beego/validation"
	"github.com/gorilla/schema"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"reflect"
	"sync"

	. "github.com/yeqown/gweb/logger"
	"github.com/yeqown/gweb/utils"
	// "github.com/yeqown/gweb"
)

var (
	decoder   = schema.NewDecoder()
	poolValid = &sync.Pool{
		New: func() interface{} {
			return &valid.Validation{RequiredFirst: true}
		},
	}
)

func init() {
	// IgnoreUnknownKeys controls the behaviour when the decoder encounters unknown keys in the map.
	// If i is true and an unknown field is encountered, it is ignored.
	// If i is false then Decode will return an error.
	// To preserve backwards compatibility, the default value is false.
	decoder.IgnoreUnknownKeys(true)
}

// ParamFile to save file from request
type ParamFile struct {
	File       multipart.File
	FileHeader multipart.FileHeader
}

// ParamError include Field, Value, Message
type ParamError struct {
	Field   string      `json:"field"`
	Value   interface{} `json:"value"`
	Message string      `json:"message"`
}

// String of ParamError, to format as string
func (pe *ParamError) String() string {
	return utils.Fstring("field-[%s], invalid with value-[%s], tip: [%s]", pe.Field, pe.Value, pe.Message)
}

// Error stringfiy ParamError
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

// Errors saved errs in valid.Error type
type Errors []*valid.Error

// ParseParams, parse params into reqRes from req.Form, and support
// form-data, json-body
// TODO: support parse file
func ParseParams(w http.ResponseWriter, req *http.Request, reqRes interface{}) (errs []*ParamError) {
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

	// if should parse Json body
	// parse json into reqRes
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

	// if has FILES field,
	// so parese req to get attachment files
	if shouldParseFile(reqRes) {
		AppL.Info("should parse files")
		if req.MultipartForm == nil || req.MultipartForm.File == nil {
			errs = append(errs, NewParamError("FILES", "empty file param", ""))
			return
		}
		rv := reflect.ValueOf(reqRes).Elem().FieldByName("FILES")
		// typ := reflect.ValueOf(reqRes).Elem().FieldByName("FILES").Type()
		filesMap := reflect.MakeMap(rv.Type())

		// parse file loop
		for key, _ := range req.MultipartForm.File {
			file, file_header, err := req.FormFile(key)
			if err != nil {
				errs = append(errs, NewParamError(utils.Fstring("parse request.FormFile: %s", key),
					err.Error(), ""))
			}
			defer file.Close()

			filesMap.SetMapIndex(
				reflect.ValueOf(key),
				reflect.ValueOf(ParamFile{
					File:       file,
					FileHeader: *file_header,
				}),
			)
		} // loop end

		// set value to reqRes.Field `FILES`
		rv.Set(filesMap)

		if len(errs) != 0 {
			return
		}
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

// shouldParseFile check i has filed `FILE`
func shouldParseFile(i interface{}) bool {
	v := reflect.ValueOf(i).Elem()
	// field not ZeroValie means true
	if _, ok := v.Type().FieldByName("FILES"); !ok {
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
