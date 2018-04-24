package services

import (
	. "gweb/logger"
	M "gweb/models"
	"gweb/utils"

	"time"
)

type Time struct {
}

type User struct {
	Time
	ID         int64     `json:"id"`
	Mobile     string    `json:"mobile"`
	Password   string    `json:"-"`
	Salt       string    `json:"-"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func NewUserFromModel(mu *M.UserInfo) *User {
	return &User{
		ID:         mu.ID,
		Mobile:     mu.Mobile,
		Password:   mu.Password,
		Salt:       mu.Salt,
		CreateTime: mu.CreateTime,
		UpdateTime: mu.UpdateTime,
	}
}

func AddUser(u *User) error {
	nuc := M.NewUserInfoColl()

	u.CreateTime = time.Now()
	u.CreateTime = time.Now()

	u.Salt = utils.RandStr(6)
	u.Password = utils.GenPasswordHash(u.Password, u.Salt)

	err := nuc.Create(&M.UserInfo{
		Mobile:     u.Mobile,
		Password:   u.Password,
		Salt:       u.Salt,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}).Error
	return err
}

func IsMobileReged(mobile string) bool {
	nuc := M.NewUserInfoColl()
	// mu := new(M.UserInfo)
	cnt := 0
	if err := nuc.Where("mobile = ?", mobile).Count(&cnt).Error; err != nil {
		AppL.Error(err.Error())
		return true
	}
	return cnt > 0
}

func FindUserWithMobile(mobile string) (*User, error) {
	nuc := M.NewUserInfoColl()
	mu := new(M.UserInfo)
	if err := nuc.Where("mobile = ?", mobile).
		First(mu).Error; err != nil {
		return nil, err
	}
	return NewUserFromModel(mu), nil
}
