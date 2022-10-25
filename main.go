package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/zxysilent/blog/conf"
	"github.com/zxysilent/blog/internal/model"
	"github.com/zxysilent/blog/internal/router"

	"github.com/zxysilent/blog/logs"
)

// @Title Blog’s 接口文档
// @Version 1.0
// @Description token传递方式包括[get/post]token、[header]Authorization
// @Description /api/* 公共访问
// @Description /adm/* 必须传入 token
// @Host 127.0.0.1:8085
// @BasePath /
func main() {
	logs.Info("app initializing")
	conf.Init()
	model.Init()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	logs.Info("app running")
	go router.RunApp()
	<-quit
	model.Close()
	logs.Info("app quitted")
	logs.Flush()
}