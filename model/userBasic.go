package model

import (
	"go-chat/global"
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

func GetUserList() []*UserBasic {
	global.Log.Infof("获取用户列表")
	var res []*UserBasic

	global.DB.Find(&res)
	return res
}

func AddUserBasic(user *UserBasic) error {
	global.DB.Create(user)
	return nil
}
