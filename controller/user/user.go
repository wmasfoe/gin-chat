package user

import (
	"github.com/gin-gonic/gin"
	"go-chat/global"
	"go-chat/model"
	"go-chat/model/respModel"
	"go-chat/services"
)

type UsersController struct {
}

func (u UsersController) UserList(c *gin.Context) {
	res := services.UserService.GetUserList()
	respModel.SuccessWithData(res, c)
}

func (u UsersController) UserAdd(c *gin.Context) {
	tmpUser := model.UserBasic{}
	err := c.ShouldBindJSON(&tmpUser)
	if err != nil {
		global.Log.Errorf("add user Error: %v", err.Error())
		respModel.FailWithMsg("新建用户发生异常", c)
	}

	services.UserService.AddUser(&tmpUser)

	respModel.SuccessWithMsg("新建用户成功", c)
}
