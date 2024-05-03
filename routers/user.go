package routers

import (
	"go-chat/controller"
)

func (r RouterGroup) UserRouter() {
	// TODO 示例声明了 /api/user/list 接口
	r.GET("list", controller.Controller.UserController.UserList)
	r.POST("", controller.Controller.UserController.UserAdd)
}
