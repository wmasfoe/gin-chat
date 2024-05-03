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

// UserList
// @Summary 获取所有用户列表
// @Tags 用户
// @Accept application/json
// @Success 200 {string} json{"code", "msg", "data"}
// @Router /api/user/list [get]
func (u UsersController) UserList(c *gin.Context) {
	res := services.UserService.GetUserList()
	respModel.SuccessWithData(res, c)
}

// UserAdd
// @Summary 新建用户接口
// @Tags 用户
// @Accept application/json
// @Param Name query string true "用户名"
// @Param Password query string true "密码"
// @Success 200 {string} json{"code", "msg", "data"}
// @Router /api/user [post]
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
