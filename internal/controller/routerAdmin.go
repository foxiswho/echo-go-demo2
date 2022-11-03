package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/pangu-2/go-echo-demo/internal/controller/admin"
)

// admRouter 登录访问
func AdminRouter(e *echo.Echo, mid echo.MiddlewareFunc) {
	api := e.Group("/api")                     // api/
	api.GET("/auth/vcode", admin.AuthVcode)    // 验证码
	api.POST("/auth/login", admin.AuthLogin)   // 用户登陆
	api.POST("/auth/logout", admin.UserLogout) // 注销登陆

	//
	adm := e.Group("/adm", mid)                  // adm/ 需要登陆才能访问
	adm.GET("/auth/get", admin.AuthGet)          // 登陆信息
	adm.POST("/upload/file", admin.UploadFile)   // 文件上传
	adm.POST("/upload/image", admin.UploadImage) // 图片上传
}
