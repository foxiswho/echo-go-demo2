package baseContext

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"gitea.com/lunny/log"
	"github.com/labstack/echo/v4"
	"github.com/pangu-2/go-echo-demo/middleware/baseAuth"
	"github.com/pangu-2/go-echo-demo/pkg/base/interfaces"
)

type BaseContext struct {
	echo.Context
}

func NewBaseContext() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := &BaseContext{Context: c}
			return h(ctx)
		}
	}
}

//func (ctx *BaseContext) Session() session.Session {
//	return session.Default(ctx)
//}
//
//func (ctx *BaseContext) Auth() auth.Auth {
//	return auth.Default(ctx)
//}

// func (ctx *BaseContext) OpenTracingSpan() opentracing.Span {
// 	return ot.Default(ctx)
// }

/**
 * 以接口参数或后缀返回数据
 * JSONP、JSON/XML
 */
func (c *BaseContext) AutoFMT(code int, i interface{}) (err error) {
	// JSONP
	callback := c.QueryParam("callback")
	log.Infof("callback=%#v", callback)
	if len(callback) > 0 {
		c.Logger().Infof("JSONP callback func:%v", callback)
		return c.JSONP(code, callback, i)
	} else {
		return c.JSON(code, i)
	}
}

// 自定义JSON解析
//func (c *BaseContext) CustomJSON(code int, i interface{}, f string) (err error) {
//	if c.Context.Echo().Debug {
//		return c.JSONPretty(code, i, "  ")
//	}
//	b, err := json.MarshalFilter(i, f)
//	if err != nil {
//		return
//	}
//	return c.JSONBlob(code, b)
//}

func (ctx *BaseContext) CookieGet(name string) (string, error) {
	cookie, err := ctx.Request().Cookie(name)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

// others are ordered as cookie's max age time, path,domain, secure and httponly.
func (ctx *BaseContext) CookieSet(name string, value string, others ...interface{}) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Path = "/"
	cookie.Value = value
	//fix cookie not work in IE
	if len(others) > 0 {
		var maxAge int

		switch v := others[0].(type) {
		case int:
			maxAge = v
		case int32:
			maxAge = int(v)
		case int64:
			maxAge = int(v)
		}
		switch {
		case maxAge > 0:
			cookie.MaxAge = maxAge
		case maxAge < 0:
			cookie.MaxAge = 0
		}
	}

	// the settings below
	// Path, Domain, Secure, HttpOnly
	// can use nil skip set

	// default "/"
	if len(others) > 1 {
		if v, ok := others[1].(string); ok && len(v) > 0 {
			cookie.Path = v
		}
	}

	// default empty
	if len(others) > 2 {
		if v, ok := others[2].(string); ok && len(v) > 0 {
			cookie.Domain = v
		}
	}

	// default empty
	if len(others) > 3 {
		var secure bool
		switch v := others[3].(type) {
		case bool:
			secure = v
		default:
			if others[3] != nil {
				secure = true
			}
		}
		if secure {
			cookie.Secure = true
		}
	}
	// default false. for session cookie default true
	if len(others) > 4 {
		if v, ok := others[4].(bool); ok && v {
			cookie.HttpOnly = true
		}
	}
	http.SetCookie(ctx.Response(), cookie)
}

func (ctx *BaseContext) CookieDel(name string) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.MaxAge = -1
	http.SetCookie(ctx.Response(), cookie)
}

// 获取登陆用户信息,得到的是副本,修改后数据,不会体现在原数据中
func (ctx *BaseContext) AuthHolder() interfaces.IHolderPg {
	return baseAuth.Get(ctx)
}

// 获取登陆用户信息,得到的原始数据,修改后,原始数据 即为修改
func (ctx *BaseContext) AuthHolderPointer() *interfaces.IHolderPg {
	get := baseAuth.Get(ctx)
	return &get
}

// 用户是否登陆
func (ctx *BaseContext) AuthHolderIs() bool {
	return baseAuth.GetIs(ctx)
}

// 绑定arr
func (ctx *BaseContext) BindArrString() (arr []string, err error) {
	var all []byte
	all, err = ioutil.ReadAll(ctx.Context.Request().Body)
	if err != nil {
		log.Errorf(err.Error())
	} else {
		log.Infof("format=%#v", string(all))
	}
	err = json.Unmarshal(all, &arr)
	return
}

// 绑定arr
func (ctx *BaseContext) BindArrInt64() (arr []int64, err error) {
	var all []byte
	all, err = ioutil.ReadAll(ctx.Context.Request().Body)
	if err != nil {
		log.Errorf(err.Error())
	} else {
		log.Infof("format=%#v", string(all))
	}
	err = json.Unmarshal(all, &arr)
	return
}

// 绑定arr
func (ctx *BaseContext) BindArrInt32() (arr []int32, err error) {
	var all []byte
	all, err = ioutil.ReadAll(ctx.Context.Request().Body)
	if err != nil {
		log.Errorf(err.Error())
	} else {
		log.Infof("format=%#v", string(all))
	}
	err = json.Unmarshal(all, &arr)
	return
}

// 绑定arr
func (ctx *BaseContext) BindArrInt() (arr []int, err error) {
	var all []byte
	all, err = ioutil.ReadAll(ctx.Context.Request().Body)
	if err != nil {
		log.Errorf(err.Error())
	} else {
		log.Infof("format=%#v", string(all))
	}
	err = json.Unmarshal(all, &arr)
	return
}
