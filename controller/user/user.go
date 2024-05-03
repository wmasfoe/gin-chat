package user

import (
	"github.com/gin-gonic/gin"
	"go-chat/model/respModel"
	"go-chat/services"
)

type UsersController struct {
}

func (u UsersController) UserList(c *gin.Context) {
	res := services.UserService.GetUserList()
	respModel.SuccessWithData(res, c)
}
