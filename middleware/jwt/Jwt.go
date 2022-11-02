package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func MakeJwt(name string, isAdmin bool) (string, error) {
	// Set custom claims
	claims := &jwtCustomClaims{
		name,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}

func UseMidJwt() echo.MiddlewareFunc {
	signingKey := []byte("secret")
	config := middleware.JWTConfig{
		ContextKey: "user",
		// TokenLookup: "query:token",
		TokenLookup: "header:" + echo.HeaderAuthorization,
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
	return middleware.JWTWithConfig(config)
}
