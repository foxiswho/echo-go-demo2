package jwtMid

import (
	"errors"
	"fmt"
	"time"

	jwtG "github.com/golang-jwt/jwt"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pangu-2/go-echo-demo/internal/consts"
	"github.com/pangu-2/go-echo-demo/pkg/base/const/constContext"
	"github.com/pangu-2/go-echo-demo/pkg/base/holder/jwtHolder"
	"github.com/pangu-2/go-echo-demo/pkg/base/interfaces"
	"github.com/pangu-2/pangu-config/configs"
)

// JwtPg 用户 会话信息 登录人信息
type Jwt struct {
	MultiTenant   interfaces.IMultiTenantPg `json:"mTenant,omitempty" multiTenant` //多租户
	No            string                    `json:"no"`                            //系统编号
	LoginUserName string                    `json:"loginUserName"`                 //登录用户名
	Name          string                    `json:"name"`                          //显示名称
	OrgName       string                    `json:"OrgName,omitempty"`             //组织名称
	TenantName    string                    `json:"tName,omitempty"`               //组织名称
	Type          int64                     `json:"type,omitempty"`                //类型
	Other         string                    `json:"other,omitempty"`               //其他信息
	Version       string                    `json:"version,omitempty"`             //版本
	Extra         interface{}               `json:"extra,omitempty"`               //额外的，扩展
	jwtG.StandardClaims

	// LoginNo       string                     `json:"loginNo"`             //登录用户No,随时可以修改变动
}

// 生成jwt
func MakeJwt(param jwtHolder.JwtPg, module consts.AppModule) (string, error) {
	jwtC := configs.GetPgJwt()
	// Set custom claims
	claims := &Jwt{
		MultiTenant: param.MultiTenant,
	}
	copier.Copy(&param, &claims)
	// 默认2天
	minute := 60 * 24 * 2
	if consts.APP_ADMIN == module {
		con := jwtC.Admin
		if con.Expire > 0 {
			minute = con.Expire
		}
		// jwt签发者
	} else if consts.APP_WEB == module {
		con := jwtC.Web
		if con.Expire > 0 {
			minute = con.Expire
		}
	}
	// jwt的过期时间，这个过期时间必须要大于签发时间
	claims.ExpiresAt = time.Now().Add(time.Minute * time.Duration(minute)).Unix()
	// jwt的签发时间
	claims.IssuedAt = time.Now().Unix()
	//jwt所面向的用户
	claims.Subject = param.LoginUserName
	//jwt的唯一身份标识，主要用来作为一次性token
	claims.Id = param.LoginNo
	// iss: jwt签发者
	// sub: jwt所面向的用户
	// aud: 接收jwt的一方
	// exp: jwt的过期时间，这个过期时间必须要大于签发时间
	// nbf: 定义在什么时间之前，该jwt都是不可用的.
	// iat: jwt的签发时间
	// jti: jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击。

	// Create token with claims
	token := jwtG.NewWithClaims(jwtG.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}

// 默认jwt 中间件
func UseMidJwt(module consts.AppModule) echo.MiddlewareFunc {
	jwtC := configs.GetPgJwt()
	signingKey := []byte("secret")
	if consts.APP_ADMIN == module {
		signingKey = []byte(jwtC.Admin.Secret)
	} else {
		signingKey = []byte(jwtC.Web.Secret)
	}
	config := middleware.JWTConfig{
		SigningMethod: jwtG.SigningMethodES512.Name,
		SigningKey:    signingKey,
		ContextKey:    constContext.AUTH_LOGIN,
		TokenLookup:   "header:" + echo.HeaderAuthorization,
		AuthScheme:    "Bearer",
		Claims:        jwtG.MapClaims{},
	}
	return middleware.JWTWithConfig(config)
}

// 自定义 未实现
func UseMidJwtCustom(module consts.AppModule) echo.MiddlewareFunc {
	jwtC := configs.GetPgJwt()
	signingKey := []byte("secret")
	if consts.APP_ADMIN == module {
		signingKey = []byte(jwtC.Admin.Secret)
	} else {
		signingKey = []byte(jwtC.Web.Secret)
	}
	config := middleware.JWTConfig{
		ContextKey:  constContext.AUTH_LOGIN,
		TokenLookup: "header:" + echo.HeaderAuthorization,
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			keyFunc := func(t *jwtG.Token) (interface{}, error) {
				if t.Method.Alg() != "HS256" {
					return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
				}
				return signingKey, nil
			}

			// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
			token, err := jwtG.Parse(auth, keyFunc)
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, errors.New("invalid token")
			}
			return token, nil
		},
	}
	return middleware.JWTWithConfig(config)
}
