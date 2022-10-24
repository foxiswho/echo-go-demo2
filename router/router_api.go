package router

import (
	"github.com/zxysilent/blog/control/sysctl"

	"github.com/labstack/echo/v4"
)

// apiRouter 通用访问
func apiRouter(api *echo.Group) {
	api.GET("/auth/vcode", sysctl.AuthVcode) // 验证码
	// api.GET("/global/get", sysctl.GlobalGet)    // 全局配置
	api.POST("/auth/login", sysctl.AuthLogin)   // 用户登陆
	api.POST("/auth/logout", sysctl.UserLogout) // 注销登陆
}
