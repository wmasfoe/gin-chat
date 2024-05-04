package model

import (
	"gorm.io/gorm"
)

// TODO @dev 这里只对数据库的表结构进行抽象，建议其他逻辑在对应 services 处理
type UserBasic struct {
	gorm.Model
	Identity   string `json:"identity"`
	Name       string `json:"name"`
	NickName   string `json:"nickName"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	ClientIp   string `json:"clientIp"`
	ClientPort string `json:"clientPort"`
	LoginTime  string `json:"loginTime"`  // 登录的时间
	LogoutTime string `json:"logoutTime"` // 下线的时间
	IsLogout   bool   `json:"isLogout"`
	DeviceInfo string `json:"deviceInfo"`
}

// UserBasicNotPrivate 脱敏的user信息
type UserBasicNotPrivate struct {
	Identity   string `json:"identity"`
	Name       string `json:"name"`
	NickName   string `json:"nickName"`
	Email      string `json:"email"`
	LoginTime  string `json:"loginTime"`  // 登录的时间
	LogoutTime string `json:"logoutTime"` // 下线的时间
	IsLogout   bool   `json:"isLogout"`
}

var UserBasicNotPrivateList = []string{
	"Identity",
	"Name",
	"NickName",
	"Email",
	"LoginTime",
	"LogoutTime",
	"IsLogout",
}

// UserBasicDetail user详细信息，没有脱敏
type UserBasicDetail struct {
	Identity   string `json:"identity"`
	Name       string `json:"name"`
	NickName   string `json:"nickName"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	ClientIp   string `json:"clientIp"`
	ClientPort string `json:"clientPort"`
	LoginTime  string `json:"loginTime"`  // 登录的时间
	LogoutTime string `json:"logoutTime"` // 下线的时间
	IsLogout   bool   `json:"isLogout"`
	DeviceInfo string `json:"deviceInfo"`
}

func (u UserBasic) TableName() string {
	return "user_basic"
}
