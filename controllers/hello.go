package controllers

import (
	. "gweb/constant"
	. "gweb/logger"
	mw "gweb/router/middleware"

	"bufio"
	"fmt"
	"sync"
	"time"
)

/*
 * Get Demo
 */
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

	// TODO: sleep over 10 *time.Second, test Response TimeOut
	time.Sleep(10 * time.Second)

	Response(resp, NewCodeInfo(CodeOk, ""))
	return resp
}

/*
 * POST Demo
 */
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

/*
 * PUT Demo
 */
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

	resp.Tip = fmt.Sprintf("PUT Hello, %s! your age[%d] is valid to access", req.Name, req.Age)

	Response(resp, NewCodeInfo(CodeOk, ""))
	return resp
}

/*
 * JSON-Body Demo
 */
type HelloJsonBodyForm struct {
	JSON bool   `schema:"-" json:"-"`
	Name string `schema:"name" valid:"Required" json:"name"`
	Age  int    `schema:"age" valid:"Required;Min(0)" json:"age"`
}

var PoolHelloJsonBodyForm = &sync.Pool{New: func() interface{} { return &HelloJsonBodyForm{} }}

type HelloJsonBodyResp struct {
	CodeInfo
	Tip string `json:"tip"`
}

var PoolHelloJsonBodyResp = &sync.Pool{New: func() interface{} { return &HelloJsonBodyResp{} }}

func HelloJsonBody(req *HelloJsonBodyForm) *HelloJsonBodyResp {
	resp := PoolHelloJsonBodyResp.Get().(*HelloJsonBodyResp)
	defer PoolHelloJsonBodyResp.Put(resp)

	resp.Tip = fmt.Sprintf("JSON-Body Hello, %s! your age[%d] is valid to access", req.Name, req.Age)

	Response(resp, NewCodeInfo(CodeOk, ""))
	return resp
}

/*
 * File Hanlder demo
 */

type HelloFileForm struct {
	FILES map[string]mw.ParamFile `schema:"-" json:"-"`
	Name  string                  `schema:"name" valid:"Required"`
	Age   int                     `schema:"age" valid:"Required"`
}

var PoolHelloFileForm = &sync.Pool{New: func() interface{} { return &HelloFileForm{} }}

type HelloFileResp struct {
	CodeInfo
	Data struct {
		Tip  string `json:"tip"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	} `json:"data"`
}

var PoolHelloFileResp = &sync.Pool{New: func() interface{} { return &HelloFileResp{} }}

func HelloFile(req *HelloFileForm) *HelloFileResp {
	resp := PoolHelloFileResp.Get().(*HelloFileResp)
	defer PoolHelloFileResp.Put(resp)

	resp.Data.Tip = "foo"
	for key, paramFile := range req.FILES {
		AppL.Infof("%s:%s\n", key, paramFile.FileHeader.Filename)
		s, _ := bufio.NewReader(paramFile.File).ReadString(0)
		resp.Data.Tip += s
	}

	resp.Data.Name = req.Name
	resp.Data.Age = req.Age

	Response(resp, NewCodeInfo(CodeOk, ""))
	return resp
}
