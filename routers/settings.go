package routers

import (
	"go-chat/controller"
)

func (r RouterGroup) SettingRouter() {
	settingsController := controller.Controller.SettingsController

	r.GET("settings", settingsController.SettingsInfoView)
}
