package gweb

import (
	"fmt"
	"net/http"
	"reflect"
	"sync"

	. "github.com/yeqown/gweb/logger"
	"github.com/yeqown/gweb/middleware"
	. "github.com/yeqown/gweb/utils"
)

func init() {
	// set other handler
	ApiHdl.NotFound = NfController
	ApiHdl.MethodNotAllowed = MnaController
}

type ApiHandler struct {
	NotFound         http.Handler
	MethodNotAllowed http.Handler
}

// JsonErr Includes `Errs` field which contains interface{} value
type JsonErr struct {
	CodeInfo
	Errs interface{} `json:"errs"`
}

var (
	ApiHdl       = &ApiHandler{}
	Routes       = map[string][]*Route{}
	assRoutesMap = map[string]bool{} // check method not allowed
	// FileHdl *FileHanlder
)

func (a *ApiHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	defer middleware.SafeHandler(w, req)

	path := req.URL.Path
	route, ok := foundRoute(path, req.Method)

	//// handle 404
	if !ok {
		if a.NotFound != nil {
			a.NotFound.ServeHTTP(w, req)
		} else {
			http.NotFound(w, req)
		}
		return
	}

	// not nil and to, ref to foundRoute
	if route != nil {
		goto Found
	}

	//// handle 405
	if !allowed(path, req.Method) {
		if a.MethodNotAllowed != nil {
			a.MethodNotAllowed.ServeHTTP(w, req)
		} else {
			http.Error(w,
				http.StatusText(http.StatusMethodNotAllowed),
				http.StatusMethodNotAllowed,
			)
		}
		return
	}

Found:
	//// normal handle
	reqRes := route.ReqPool.Get()
	defer route.ReqPool.Put(reqRes)

	// parse params
	if errs := middleware.ParseParams(w, req, reqRes); len(errs) != 0 {
		je := new(JsonErr)
		Response(je, NewCodeInfo(CodeParamInvalid, ""))
		je.Errs = errs
		middleware.ResponseErrorJson(w, je)
		return
	}
	in := make([]reflect.Value, 1)
	in[0] = reflect.ValueOf(reqRes)
	Fn := reflect.ValueOf(route.Fn)

	//// Call web server handle function
	out := Fn.Call(in)

	//// response to client
	resp := out[0].Interface()
	defer route.ResPool.Put(resp)

	middleware.ResponseJson(w, resp)
	return
}

// func (f *FileHanlder) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	http.FileServer(http.Dir("/User/yeqiang"))
// }

type Route struct {
	Path    string      // Route Path
	Method  string      // Route Method
	Fn      interface{} // Route Fn
	ReqPool *sync.Pool  // Route ReqPool
	ResPool *sync.Pool  // Route ResPool
}

func AddRoute(r *Route) {
	AppL.Infof("Adding route: %s %s", r.Method, r.Path)
	if _, ok := Routes[r.Path]; !ok {
		Routes[r.Path] = []*Route{}
	}
	Routes[r.Path] = append(Routes[r.Path], r)
	arkey := assMapKey(r.Path, r.Method)
	assRoutesMap[arkey] = true
}

func assMapKey(path, method string) string {
	return fmt.Sprintf("%s_%s", path, method)
}

func foundRoute(path, reqMethod string) (*Route, bool) {
	routes, ok := Routes[path]
	if !ok {
		// no path
		return nil, false
	}
	for idx, r := range routes {
		if r.Method == reqMethod {
			// eauql
			return routes[idx], true
		}
	}
	// means no method equal with path and method
	return nil, true
}

func allowed(path, reqMethod string) bool {
	arkey := assMapKey(path, reqMethod)
	if _, ok := assRoutesMap[arkey]; !ok {
		return false
	}
	return true
}
