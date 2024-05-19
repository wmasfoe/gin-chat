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
// @Router /api/user [get]
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

	defer func() {
		if err := recover(); err != nil {
			global.Log.Errorf("add user Error: %v", err)
			respModel.FailWithMsg("新建用户发生异常", c)
		}
	}()

	tmpUser := model.UserBasic{}
	c.ShouldBindJSON(&tmpUser)
	services.UserService.AddUser(&tmpUser)

	respModel.SuccessWithMsg("新建用户成功", c)
}

// UserDel
// @Summary 删除用户
// @Tags 用户
// @Accept application/json
// @Success 200 {string} json{"code", "msg", "data"}
// @Router /api/user/:id [delete]
func (u UsersController) UserDel(c *gin.Context) {
	userID := c.Param("id")
	err := services.UserService.DeleteUser(userID)

	if err != nil {
		global.Log.Errorf("add user Error: %v", err.Error())
		respModel.FailWithMsg("删除用户发生异常", c)
	}

	respModel.SuccessWithMsg("删除用户成功", c)
}

// UserDesc
// @Summary 用户详情
// @Tags 用户
// @Accept application/json
// @Success 200 {string} json{"code", "msg", "data"}
// @Router /api/user/:id [get]
func (u UsersController) UserDesc(c *gin.Context) {

	defer func() {
		if err := recover(); err != nil {
			global.Log.Errorf("user desc Error: %v", err)
			respModel.FailWithMsg("获取用户详情失败", c)
		}
	}()

	userID := c.Param("id")
	userDesc, _ := services.UserService.UserDesc(userID)

	respModel.SuccessWithData(userDesc, c)
}

// UserUpdate
// @Summary 更新用户信息
// @Tags 用户
// @Accept application/json
// @Params name query string false "用户名"
// @Params passWord query string false "密码"
// @Params NickName query string false "昵称"
// @Params Email query string false "邮箱"
// @Success 200 {string} json{"code", "msg", "data"}
// @Router /api/user/:id [put]
func (u UsersController) UserUpdate(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			global.Log.Errorf("user update Error: %v", err)
			respModel.FailWithMsg("更新用户信息失败", c)
		}
	}()

	userID := c.Param("id")

	var tmpUser model.UserBasicNotPrivate
	c.ShouldBindJSON(&tmpUser)

	services.UserService.UserUpdate(userID, tmpUser)

	respModel.SuccessWithData(map[string]any{}, c)
}
