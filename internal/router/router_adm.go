package router

import (
	"github.com/zxysilent/blog/internal/controller/admin"

	"github.com/labstack/echo/v4"
)

// admRouter 登录访问
func admRouter(adm *echo.Group) {
	adm.GET("/auth/get", admin.AuthGet) // 登陆信息
	// adm.GET("/status/goinfo", sysctl.StatusGoinfo) // 环境信息
	adm.POST("/upload/file", admin.UploadFile)   // 文件上传
	adm.POST("/upload/image", admin.UploadImage) // 图片上传
	// adm.POST("/global/edit", sysctl.GlobalEdit)    // 配置修改
}
