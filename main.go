package main

import (
	"go-chat/core"
	"go-chat/global"
	"go-chat/routers"
)

func main() {
	core.InitConf()
	core.InitGorm()
	core.InitLogger()
	router := routers.InitRouter()

	listenAddr := global.CONFIG.System.Addr()
	global.Log.Infof("go-chat 程序启动：%s", listenAddr)
	router.Run(listenAddr)
}
