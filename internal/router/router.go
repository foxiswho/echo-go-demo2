package router

import (
	"github.com/pangu-2/go-echo-demo/internal/consts"
	"github.com/pangu-2/go-echo-demo/internal/controller"
	"github.com/pangu-2/go-echo-demo/internal/controller/web"
	"github.com/pangu-2/go-echo-demo/middleware/jwtMid"
	"github.com/pangu-2/pangu-config/configs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pangu-2/go-echo-demo/internal/logs"
)

// RunApp 入口
func RunApp() {
	e := echo.New()                              // 实例化echo
	e.Renderer = initRender()                    // 初始渲染引擎
	e.Use(midRecover, midLogger)                 // 恢复 日志记录
	e.Use(middleware.CORSWithConfig(crosConfig)) // 跨域设置
	//
	e.HideBanner = true                   // 不显示横幅
	e.HTTPErrorHandler = HTTPErrorHandler // 自定义错误处理
	e.Debug = true                        // 运行模式 - echo框架好像没怎么使用这个
	RegDocs(e)                            // 注册文档
	e.Static("/dist", "dist")             // 静态目录 - 后端专用
	//engine.StaticFS("/static", bi.StaticFS)           // 静态目录
	e.Static("/static", "static")            // 静态目录
	e.File("/favicon.ico", "favicon.ico")    // ico
	e.File("/dashboard*", "dist/index.html") // 前后端分离页面
	// 后台登录
	e.GET("/login.html", func(ctx echo.Context) error {
		// 302 临时重定向
		return ctx.Redirect(302, "/dashboard/")
	})
	e.POST("/login", web.Login)
	// Restricted group
	r := e.Group("/restricted")
	{
		r.Use(jwtMid.UseMidJwt(consts.APP_WEB))
		r.GET("", web.Restricted)
	}

	// qq登录
	// engine.GET("/login/qq.html", sysctl.ViewLoginQq)
	// engine.GET("/auth/qq.html", sysctl.ViewAuthQq)
	//--- 页面 -- start
	// engine.GET("/", appctl.ViewIndex)              // 首页
	// engine.GET("/cate/:cate", appctl.ViewCatePost) // 分类
	//--- 页面 -- end
	controller.AdminRouter(e, midAuth)
	err := e.Start(":" + configs.GetApplication().PortToString())
	if err != nil {
		logs.Fatal("run error :", err.Error())
	}
}
