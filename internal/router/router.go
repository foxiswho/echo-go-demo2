package router

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/pangu-2/pangu-config/configs"
	"github.com/zxysilent/blog/internal/controller"
	"github.com/zxysilent/blog/internal/controller/web"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zxysilent/blog/internal/logs"
)

// RunApp 入口
func RunApp() {
	e := echo.New()                              // 实例化echo
	e.Renderer = initRender()                    // 初始渲染引擎
	e.Use(midRecover, midLogger)                 // 恢复 日志记录
	e.Use(middleware.CORSWithConfig(crosConfig)) // 跨域设置
	// engine.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey:  []byte("secret"),
	// 	TokenLookup: "query:token",
	// }))
	signingKey := []byte("secret")

	config := middleware.JWTConfig{
		TokenLookup: "query:token",
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			keyFunc := func(t *jwt.Token) (interface{}, error) {
				if t.Method.Alg() != "HS256" {
					return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
				}
				return signingKey, nil
			}

			// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
			token, err := jwt.Parse(auth, keyFunc)
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, errors.New("invalid token")
			}
			return token, nil
		},
	}

	e.Use(middleware.JWTWithConfig(config))
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
	// Configure middleware with the custom claims type
	// config := middleware.JWTConfig{
	// 	Claims:     &jwtCustomClaims{},
	// 	SigningKey: []byte("secret"),
	// }
	// r.Use(middleware.JWTWithConfig(config))
	r.GET("", web.Restricted)
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
