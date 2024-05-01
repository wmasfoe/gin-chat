package main

import (
	"go-blog-server/core"
	"go-blog-server/global"
	"go-blog-server/routers"
)

func main() {
	core.InitConf()
	core.InitGorm()
	core.InitLogger()
	router := routers.InitRouter()

	listenAddr := global.CONFIG.System.Addr()
	global.Log.Infof("go-blog-server 程序启动：%s", listenAddr)
	router.Run(listenAddr)
}
