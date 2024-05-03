package setting

import (
	"github.com/gin-gonic/gin"
	"go-chat/model/respModel"
)

type SettingsController struct {
}

func (s SettingsController) SettingsInfoView(c *gin.Context) {
	respModel.SuccessWithData(map[string]interface{}{
		"id": 1,
	}, c)
	//respModel.FailWithCode(respModel.SystemError, c)
}
