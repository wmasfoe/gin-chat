package main

import (
	"go-admin/core"
	"go-admin/global"
	"go-admin/routers"
)

func main() {
	core.InitConf()
	core.InitGorm()
	core.InitLogger()
	router := routers.InitRouter()

	listenAddr := global.CONFIG.System.Addr()
	global.Log.Infof("go-admin 程序启动：%s", listenAddr)
	router.Run(listenAddr)
}
