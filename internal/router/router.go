package router

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/pangu-2/pangu-config/configs"
	"github.com/zxysilent/blog/internal/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zxysilent/blog/internal/logs"
)

// RunApp 入口
func RunApp() {
	engine := echo.New()                              // 实例化echo
	engine.Renderer = initRender()                    // 初始渲染引擎
	engine.Use(midRecover, midLogger)                 // 恢复 日志记录
	engine.Use(middleware.CORSWithConfig(crosConfig)) // 跨域设置
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

	engine.Use(middleware.JWTWithConfig(config))
	engine.HideBanner = true                   // 不显示横幅
	engine.HTTPErrorHandler = HTTPErrorHandler // 自定义错误处理
	engine.Debug = true                        // 运行模式 - echo框架好像没怎么使用这个
	RegDocs(engine)                            // 注册文档
	engine.Static("/dist", "dist")             // 静态目录 - 后端专用
	//engine.StaticFS("/static", bi.StaticFS)           // 静态目录
	engine.Static("/static", "static")            // 静态目录
	engine.File("/favicon.ico", "favicon.ico")    // ico
	engine.File("/dashboard*", "dist/index.html") // 前后端分离页面
	// 后台登录
	engine.GET("/login.html", func(ctx echo.Context) error {
		// 302 临时重定向
		return ctx.Redirect(302, "/dashboard/")
	})
	// qq登录
	// engine.GET("/login/qq.html", sysctl.ViewLoginQq)
	// engine.GET("/auth/qq.html", sysctl.ViewAuthQq)
	//--- 页面 -- start
	// engine.GET("/", appctl.ViewIndex)              // 首页
	// engine.GET("/cate/:cate", appctl.ViewCatePost) // 分类
	//--- 页面 -- end
	controller.AdminRouter(engine, midAuth)
	err := engine.Start(":" + configs.GetApplication().PortToString())
	if err != nil {
		logs.Fatal("run error :", err.Error())
	}
}
