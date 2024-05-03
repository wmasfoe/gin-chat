package model

import (
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Identity   string
	Name       string
	Password   string
	Phone      string
	Email      string
	ClientIp   string
	ClientPort string
	LoginTime  string // 登录的时间
	LogoutTime string // 下线的时间
	IsLogout   bool
	DeviceInfo string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
