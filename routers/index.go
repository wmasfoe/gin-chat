package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"go-chat/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.CONFIG.System.Env)
	r := gin.Default()
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	// 默认后端接口以 api 开头
	ApiRouters := RouterGroup{r.Group("api")}
	// /api/user 下的路由
	userRouters := RouterGroup{ApiRouters.Group("user")}
	//setting 相关接口
	ApiRouters.SettingRouter()
	userRouters.UserRouter()

	return r
}
