package baseAuth

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/pangu-2/go-echo-demo/middleware/jwtMid"
	"github.com/pangu-2/go-echo-demo/pkg/base/const/constContext"
	"github.com/pangu-2/go-echo-demo/pkg/base/holder"
	"github.com/pangu-2/go-echo-demo/pkg/base/holder/jwtHolder"
	"github.com/pangu-2/go-echo-demo/pkg/base/interfaces"
)

// @loginVerify true: 验证登陆
func NewAuthDefault(loginVerify bool) echo.MiddlewareFunc {
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
			//
			if loginVerify {
				return echo.NewHTTPError(http.StatusUnauthorized, "jwt auth 没有任何信息")
			}
			return next(c)
		}
	}
}

// shortcut to get Auth
func Get(c echo.Context) interfaces.IHolderPg {
	return c.Get(constContext.AUTH_LOGIN).(holder.HolderSimple)
}

// shortcut to get Auth
func GetIs(c echo.Context) bool {
	get := c.Get(constContext.AUTH_LOGIN_IS)
	if get != nil {
		return get.(bool)
	}
	return false
}

// 处理 登陆用户信息
func processAuthSession(c echo.Context) (interfaces.IHolderPg, bool) {
	/////////////////////////////
	login := c.Get(constContext.HOLDER)
	hold := interfaces.NewStandardHolder()
	loginIs := false
	if login == nil {
		auth := c.Get(constContext.AUTH_LOGIN)
		if nil != auth {
			user := c.Get(constContext.AUTH_LOGIN).(*jwt.Token)
			if nil != user {
				claimsJwt := user.Claims.(*jwtMid.Jwt)
				jwtData := jwtHolder.NewJwtPg()
				var jwtDataHolder interfaces.IHolderJwtPg
				jwtDataHolder = jwtData
				//复制对象
				copier.Copy(&claimsJwt, &jwtDataHolder)
				hold.Jwt = jwtDataHolder.(*interfaces.IHolderJwtPg)
				if nil != jwtData.MultiTenant {
					hold.MultiTenant = &jwtData.MultiTenant
				}
				c.Set(constContext.HOLDER, hold)
				loginIs = true
			}
		}
	}
	// 设置会话 已有
	c.Set(constContext.HOLDER_IS, true)
	return hold, loginIs
}
