package controller

import (
	"go-chat/controller/setting"
	"go-chat/controller/user"
)

type Controllers struct {
	SettingsController setting.SettingsController
	UserController     user.UsersController
}

// 下面两种写法是一样的
var Controller = new(Controllers)

//var Controller = Controllers{
//	SettingsController: setting.SettingsController{},
//}
