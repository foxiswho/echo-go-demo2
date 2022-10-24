package web

import (
	"net/http"

	. "github.com/foxiswho/echo-go-demo2/config"
	"github.com/foxiswho/echo-go-demo2/router/base"
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
)

func Routers() *echo.Echo {
	// Echo instance
	e := echo.New()
	// Context自定义
	e.Use(base.NewBaseContext())
	// Customization
	if Config.ReleaseMode {
		e.Debug = false
	}
	//e.Use(session.MiddlewareWithConfig(session.Config{}))

	e.Static("/static", "assets")
	e.Use(mw.GzipWithConfig(mw.GzipConfig{
		Level: 5,
	}))

	// CSRF
	e.Use(mw.CSRFWithConfig(mw.CSRFConfig{
		ContextKey:  "_csrf",
		TokenLookup: "form:_csrf",
	}))

	//注册一个Get请求, 路由地址为: /hello  并且绑定一个控制器函数, 这里使用的是闭包函数。
	e.GET("/hello", func(c echo.Context) error {
		//控制器函数直接返回一个字符串，http响应状态为http.StatusOK，就是200状态。
		return c.String(http.StatusOK, "hello echo")
	})

	return e
}
