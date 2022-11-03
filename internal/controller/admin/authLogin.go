package admin

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pangu-2/go-echo-demo/internal/consts"
	"github.com/pangu-2/go-echo-demo/middleware/baseContext"
	"github.com/pangu-2/go-echo-demo/middleware/jwtMid"
	"github.com/pangu-2/go-echo-demo/pkg/base/holder/jwtHolder"
)

type AuthLogin struct {
	baseContext.BaseContext
}

func (c *AuthLogin) Login(ctx echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throws unauthorized error
	if username != "admin" || password != "admin" {
		return echo.ErrUnauthorized
	}

	jwtPg := jwtHolder.JwtPg{}
	token, err := jwtMid.MakeJwt(jwtPg, consts.APP_ADMIN)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
