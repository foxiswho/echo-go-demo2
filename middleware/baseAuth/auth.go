package baseAuth

import (
	"encoding/json"
	"net/http"
	"net/url"

	"gitea.com/lunny/log"
	"github.com/labstack/echo/v4"
	"github.com/zxysilent/blog/kit/base/const/constContext"
	"github.com/zxysilent/blog/pkg/base/holder/session"
)

func NewDefault() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			//获取登陆用户数据
			_, b := processAuthSession(c)
			if b {
				return next(c)
			}
			// header := c.Request().Header
			//判断是不是微服务,如果是那么跳过
			// get := header.Get(constContext.HEADER_AUTH_MICRO)
			// if len(get) > 0 && get == constContext.HEADER_AUTH_MICRO_VALUE {
			// 	return next(c)
			// }
			//白名单
			//TODO 后期添加 FIXME
			return echo.NewHTTPError(http.StatusUnauthorized, "jwt auth 没有任何信息")
		}
	}
}

// shortcut to get Auth
func Get(c echo.Context) session.ISessionHolderPg {
	return c.Get(AUTH_LOGIN).(session.SessionHolder)
}

// shortcut to get Auth
func GetIs(c echo.Context) bool {
	get := c.Get(AUTH_LOGIN_IS)
	if get != nil {
		return get.(bool)
	}
	return false
}

// 处理 登陆用户信息
func processAuthSession(c echo.Context) (session.ISessionHolderPg, bool) {
	header := c.Request().Header
	log.Debugf("header=%#v", header)
	log.Infof("header=%#v", header)
	get := header.Get(constContext.HEADER_AUTH)
	auth := session.SessionHolder{}
	covAuth := false
	if len(get) > 0 {
		unescape, err := url.QueryUnescape(get)
		if err == nil && len(unescape) > 0 {
			err := json.Unmarshal([]byte(unescape), &auth)
			if err == nil {
				covAuth = true
			}
		}
	}
	c.Set(AUTH_LOGIN, auth)
	c.Set(AUTH_LOGIN_IS, covAuth)
	return auth, covAuth
}
