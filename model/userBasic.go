package model

import (
	"gorm.io/gorm"
)

// TODO @dev 这里只对数据库的表结构进行抽象，建议其他逻辑在对应 services 处理
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
