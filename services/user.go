package services

import (
	"go-chat/global"
	"go-chat/model"
)

type UserServices struct{}

func (u UserServices) GetUserList() []*model.UserBasicNotPrivate {

	global.Log.Infof("获取用户列表")

	// 指定返回的数据结构中，只有 UserBasicNotPrivate 定义的属性
	var res []*model.UserBasicNotPrivate

	// DB.Model 需要绑定要查询的表
	global.DB.Model(&model.UserBasic{}).Find(&res)
	return res
}

func (u UserServices) AddUser(user *model.UserBasic) error {
	res := global.DB.Create(&user)

	if res.Error != nil {
		global.Log.Errorf("新建用户失败: %v", res.Error)
		return res.Error
	}
	return nil
}

func (u UserServices) DeleteUser(userId string) error {
	res := global.DB.Where("Identity = ?", userId).Delete(&model.UserBasic{})

	if res.Error != nil {
		global.Log.Errorf("删除用户失败: %v", res.Error)
		return res.Error
	}
	return nil
}

func (u UserServices) UserDesc(userId string) (model.UserBasicDetail, error) {
	res := model.UserBasicDetail{}
	findRes := global.DB.Model(&model.UserBasic{}).Where("Identity = ?", userId).Limit(1).Find(&res)

	if findRes.Error != nil {
		global.Log.Errorf("查询用户详情失败: %v", findRes.Error)
		panic(findRes.Error)
	}
	return res, nil
}

var UserService = new(UserServices)
