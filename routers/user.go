package routers

import (
	"go-chat/controller"
)

func (r RouterGroup) UserRouter() {
	// TODO 示例声明了 /api/user/list 接口
	r.GET("", controller.Controller.UserController.UserList)
	r.POST("", controller.Controller.UserController.UserAdd)
	r.DELETE("/:id", controller.Controller.UserController.UserDel)
	r.GET("/:id", controller.Controller.UserController.UserDesc)
	r.PUT("/:id", controller.Controller.UserController.UserUpdate)
}
