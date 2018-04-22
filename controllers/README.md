## gweb.controllers

Necessary Condition:

* 1. Fn prototype `Fn(req *reqType) *respType`
* 2. Add Fn to Route, like this:
```golang
# file: router.router.go, func: RegisterHandler
addRoute(&Route{"/hello", http.MethodPut, ctr.HelloPut, ctr.PoolHelloPutForm, ctr.PoolHelloPutResp})
```
* 3. var **`PoolOfReqType`** and **`PoolOfRespType`** ** in type **`Sync.Pool`**

### Get Method

```golang
type HelloGetForm struct {
	Name string `schema:"name" valid:"Required" json:"name"`
	Age  int    `schema:"age" valid:"Required;Min(18)" json:"age"`
}

var PoolHelloGetForm = &sync.Pool{New: func() interface{} { return &HelloGetForm{} }}

type HelloGetResp struct {
	CodeInfo
	Tip string `json:"tip"`
}

var PoolHelloGetResp = &sync.Pool{New: func() interface{} { return &HelloGetResp{} }}

func HelloGet(req *HelloGetForm) *HelloGetResp {
	resp := PoolHelloGetResp.Get().(*HelloGetResp)
	defer PoolHelloGetResp.Put(resp)

	resp.Tip = fmt.Sprintf("Get Hello, %s! your age[%d] is valid to access", req.Name, req.Age)

	Response(resp, NewCodeInfo(CodeOk, ""))
	return resp
}
```

![Get-Method](https://raw.githubusercontent.com/yeqown/gweb/master/screenshots/getmethod.png)

### Post Method

```golang
type HelloPostForm struct {
	Name string `schema:"name" valid:"Required" json:"name"`
	Age  int    `schema:"age" valid:"Required;Min(18)" json:"age"`
}

var PoolHelloPostForm = &sync.Pool{New: func() interface{} { return &HelloPostForm{} }}

type HelloPostResp struct {
	CodeInfo
	Tip string `json:"tip"`
}

var PoolHelloPostResp = &sync.Pool{New: func() interface{} { return &HelloPostResp{} }}

func HelloPost(req *HelloPostForm) *HelloPostResp {
	resp := PoolHelloPostResp.Get().(*HelloPostResp)
	defer PoolHelloPostResp.Put(resp)

	resp.Tip = fmt.Sprintf("POST Hello, %s! your age[%d] is valid to access", req.Name, req.Age)

	Response(resp, NewCodeInfo(CodeOk, ""))
	return resp
}
```
![POST-Method](https://raw.githubusercontent.com/yeqown/gweb/master/screenshots/postmethod.png)

### Put Method

```golang
type HelloPutForm struct {
	Name string `schema:"name" valid:"Required" json:"name"`
	Age  int    `schema:"age" valid:"Required;Min(18)" json:"age"`
}

var PoolHelloPutForm = &sync.Pool{New: func() interface{} { return &HelloPutForm{} }}

type HelloPutResp struct {
	CodeInfo
	Tip string `json:"tip"`
}

var PoolHelloPutResp = &sync.Pool{New: func() interface{} { return &HelloPutResp{} }}

func HelloPut(req *HelloPutForm) *HelloPutResp {
	resp := PoolHelloPutResp.Get().(*HelloPutResp)
	defer PoolHelloPutResp.Put(resp)

	resp.Tip = fmt.Sprintf("POST Hello, %s! your age[%d] is valid to access", req.Name, req.Age)

	Response(resp, NewCodeInfo(CodeOk, ""))
	return resp
}
```

![Put-Method](https://raw.githubusercontent.com/yeqown/gweb/master/screenshots/putmethod.png)