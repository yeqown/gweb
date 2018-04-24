package controllers

import (
	. "gweb/constant"
	. "gweb/logger"
	S "gweb/services"
	"gweb/utils"
	"sync"
)

type RegUserForm struct {
	Mobile   string `schema:"mobile" valid:"Required" json:"mobile"`
	Password string `schema:"password" valid:"Required;MinSize(6)" json:"password"`
}

var PoolRegUserForm = &sync.Pool{New: func() interface{} { return &RegUserForm{} }}

type RegUserResp struct {
	CodeInfo
}

var PoolRegUserResp = &sync.Pool{New: func() interface{} { return &RegUserResp{} }}

func RegisterUserPost(req *RegUserForm) *RegUserResp {
	res := PoolRegUserResp.Get().(*RegUserResp)
	defer PoolRegUserResp.Put(res)

	if S.IsMobileReged(req.Mobile) {
		Response(res, NewCodeInfo(CodeDupMobile, ""))
		return res
	}

	if err := S.AddUser(&S.User{
		Mobile:   req.Mobile,
		Password: req.Password,
	}); err != nil {
		AppL.Error(err.Error())
		Response(res, NewCodeInfo(CodeSystemErr, err.Error()))
		return res
	}
	Response(res, NewCodeInfo(CodeOk, ""))
	return res
}

type LogUserForm struct {
	Mobile   string `schema:"mobile" valid:"Required" json:"mobile"`
	Password string `schema:"password" valid:"Required;MinSize(6)" json:"password"`
}

var PoolLogUserForm = &sync.Pool{New: func() interface{} { return &LogUserForm{} }}

type LogUserResp struct {
	CodeInfo
	U *S.User `json:"user"`
}

var PoolLogUserResp = &sync.Pool{New: func() interface{} { return &LogUserResp{} }}

func LoginUserPost(req *LogUserForm) *LogUserResp {
	res := PoolLogUserResp.Get().(*LogUserResp)
	defer PoolLogUserResp.Put(res)
	res.U = nil

	password, salt := "", ""

	if !S.IsMobileReged(req.Mobile) {
		Response(res, NewCodeInfo(CodeNoRegMobile, ""))
		return res
	}

	if u, err := S.FindUserWithMobile(req.Mobile); err != nil {
		AppL.Error(err.Error())
		Response(res, NewCodeInfo(CodeSystemErr, err.Error()))
		return res
	} else {
		res.U = u
		password = u.Password
		salt = u.Salt
	}

	if password != utils.GenPasswordHash(req.Password, salt) {
		res.U = nil
		Response(res, NewCodeInfo(CodeWrongPwd, ""))
		return res
	}

	Response(res, NewCodeInfo(CodeOk, ""))
	return res
}
