package routers

import (
	"go-chat/controller"
)

func (r RouterGroup) UserRouter() {
	r.GET("list", controller.Controller.UserController.UserList)
}
