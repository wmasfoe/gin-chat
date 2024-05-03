package controller

import "go-admin/controller/setting"

type Controllers struct {
	SettingsController setting.SettingsController
}

// 下面两种写法是一样的
var Controller = new(Controllers)

//var Controller = Controllers{
//	SettingsController: setting.SettingsController{},
//}
