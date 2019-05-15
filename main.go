package main

import (
	"net/http"

	"./controllers"
	"./core"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func main() {
	//配置初始化
	core.Config.Init()
	//日志初始化
	core.Config.Logger.Init()

	var router = httprouter.New()

	index := new(controllers.IndexController)
	router.GET("/", index.Hello)
	router.GET("/login", index.Login)

	log.Infof("Server started %s ...", core.Config.Listen)
	log.Fatal(http.ListenAndServe(core.Config.Listen, router))
}
