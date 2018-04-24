package router

import (
	. "gweb/constant"
	ctr "gweb/controllers"
	. "gweb/logger"
	"gweb/router/middleware"

	// "io"
	"fmt"
	"net/http"
	"reflect"
	"sync"
)

type ApiHandler struct {
	NotFound         http.Handler
	MethodNotAllowed http.Handler
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
		je := new(middleware.JsonErr)
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
	Path    string
	Method  string
	Fn      interface{}
	ReqPool *sync.Pool
	ResPool *sync.Pool
}

func RegisterHandler() {
	AppL.Info("RegisterHandler start")
	addRoute(&Route{"/hello", http.MethodGet, ctr.HelloGet, ctr.PoolHelloGetForm, ctr.PoolHelloGetResp})
	addRoute(&Route{"/hello", http.MethodPost, ctr.HelloPost, ctr.PoolHelloPostForm, ctr.PoolHelloPostResp})
	addRoute(&Route{"/hello", http.MethodPut, ctr.HelloPut, ctr.PoolHelloPutForm, ctr.PoolHelloPutResp})
	addRoute(&Route{"/hello/json", http.MethodPost, ctr.HelloJsonBody, ctr.PoolHelloJsonBodyForm, ctr.PoolHelloJsonBodyResp})

	addRoute(&Route{"/user/register", http.MethodPost, ctr.RegisterUserPost, ctr.PoolRegUserForm, ctr.PoolRegUserResp})
	addRoute(&Route{"/user/login", http.MethodPost, ctr.LoginUserPost, ctr.PoolLogUserForm, ctr.PoolLogUserResp})

	// set other handler
	ApiHdl.NotFound = ctr.NfController
	ApiHdl.MethodNotAllowed = ctr.MnaController
}

func addRoute(r *Route) {
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
