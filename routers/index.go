package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"go-blog-server/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.CONFIG.System.Env)
	r := gin.Default()
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	ApiRouters := RouterGroup{r.Group("api")}

	//setting 相关接口
	ApiRouters.SettingRouter()

	return r
}
