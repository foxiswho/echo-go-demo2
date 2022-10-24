package router

import (
	"github.com/zxysilent/blog/control/sysctl"

	"github.com/labstack/echo/v4"
)

// admRouter 登录访问
func admRouter(adm *echo.Group) {
	adm.GET("/auth/get", sysctl.AuthGet) // 登陆信息
	// adm.GET("/status/goinfo", sysctl.StatusGoinfo) // 环境信息
	adm.POST("/upload/file", sysctl.UploadFile)   // 文件上传
	adm.POST("/upload/image", sysctl.UploadImage) // 图片上传
	// adm.POST("/global/edit", sysctl.GlobalEdit)    // 配置修改
}
