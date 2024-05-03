package routers

import (
	"go-admin/controller"
)

func (r RouterGroup) SettingRouter() {
	settingsController := controller.Controller.SettingsController

	r.GET("settings", settingsController.SettingsInfoView)
}
