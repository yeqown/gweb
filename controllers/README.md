## gweb.controllers

### Get Method

Code:
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

Code:
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