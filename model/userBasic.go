package model

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"go-chat/global"
	"gorm.io/gorm"
)

// TODO @dev 这里只对数据库的表结构进行抽象，建议其他逻辑在对应 services 处理
type UserBasic struct {
	gorm.Model
	UserBasicDetail
}

// UserBasicNotPrivate 脱敏的user信息，所有人都能看到
type UserBasicNotPrivate struct {
	Identity string `json:"identity" structs:"identity"`
	Name     string `json:"name" structs:"name"`
	// TODO 想要更新，这里采用了第一种方法，需要：1.把小驼峰转成下划线命名，或者2.列名称改成小驼峰，因为更新时转成了map，直接用map进行更新了; 或者 3.更新时map再转成userInfo结构体
	NickName   string `json:"nickName" structs:"nick_name"`
	Email      string `json:"email" structs:"email" valid:"email"`
	LoginTime  string `json:"loginTime" structs:"login_time"`   // 登录的时间
	LogoutTime string `json:"logoutTime" structs:"logout_time"` // 下线的时间
	IsLogout   bool   `json:"isLogout" structs:"is_logout"`
}

// UserBasicDetail user详细信息，没有脱敏，可供用户更改的信息
type UserBasicDetail struct {
	UserBasicNotPrivate
	Password   string `json:"password" structs:"password"`
	Phone      string `json:"phone" structs:"phone" valid:"matches(1[3-9]{1}\\d{9}$)"`
	ClientIp   string `json:"clientIp" structs:"client_ip"`
	ClientPort string `json:"clientPort" structs:"client_port"`
	DeviceInfo string `json:"deviceInfo" structs:"device_info"`
}

func (u UserBasic) TableName() string {
	return "user_basic"
}

func (u UserBasicDetail) findDataByName() (bool, error) {
	res := global.DB.Model(&UserBasic{}).Where("name = ?", u.Name).First(&u)
	hasValue := false
	if res != nil {
		hasValue = true
	}
	return hasValue, nil
}
func (u UserBasicDetail) findDataByPhone() (bool, error) {
	res := global.DB.Model(&UserBasic{}).Where("phone = ?", u.Phone).First(&u)
	hasValue := false
	if res != nil {
		hasValue = true
	}
	return hasValue, nil
}
func (u UserBasicDetail) findDataByEmail() (bool, error) {
	res := global.DB.Model(&UserBasic{}).Where("email = ?", u.Email).First(&u)
	hasValue := false
	if res != nil {
		hasValue = true
	}
	return hasValue, nil
}

func (u UserBasicDetail) ValidateData() (bool, error) {
	// 校验结构体
	structRes, structErr := govalidator.ValidateStruct(&u)
	// 校验数据是否重复
	nameRepeat, _ := u.findDataByName()
	phoneRepeat, _ := u.findDataByPhone()
	emailRepeat, _ := u.findDataByEmail()

	res := structRes || nameRepeat || phoneRepeat || emailRepeat
	e := structErr
	if nameRepeat || phoneRepeat || emailRepeat {
		e = errors.New("用户信息已存在")
	}

	return res, e
}
