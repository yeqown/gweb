package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type UserInfo struct {
	ID         int64     `gorm:"primary_key;column:id"`
	Mobile     string    `gorm:"column:mobile"`
	Password   string    `gorm:"column:password"`
	Salt       string    `gorm:"column:salt"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

type UserInfoColl struct {
	*gorm.DB
}

func NewUserInfoColl() *UserInfoColl {
	return &UserInfoColl{
		DB: getMysqlDB().Model(&UserInfo{}),
	}
}
