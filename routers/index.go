package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	_ "go-chat/docs"
	"go-chat/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.CONFIG.System.Env)
	r := gin.Default()
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	// TODO @dev 默认后端接口以 api 开头
	ApiRouters := RouterGroup{r.Group("api")}
	// TODO @dev 不进行前置分组示例 setting 相关接口
	ApiRouters.SettingRouter()

	// TODO @dev 进行前置分组示例 /api/user 下的路由
	userRouters := RouterGroup{ApiRouters.Group("user")}

	userRouters.UserRouter()

	return r
}
