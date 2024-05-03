package services

import (
	"go-chat/global"
	"go-chat/model"
)

type UserServices struct{}

func (u UserServices) GetUserList() []*model.UserBasic {
	global.Log.Infof("获取用户列表")
	var res []*model.UserBasic

	global.DB.Find(&res)
	return res
}

func (u UserServices) AddUser(user *model.UserBasic) error {
	global.DB.Create(user)
	return nil
}

var UserService = new(UserServices)
