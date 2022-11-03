package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims2 struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throws unauthorized error
	if username != "admin" || password != "admin" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &JwtCustomClaims2{
		"Jon Snow",
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	fmt.Println("Restricted=", user)
	fmt.Println("Restricted.Header=", user.Header)
	fmt.Println("Restricted.Method=", user.Method)
	fmt.Println("Restricted.Raw=", user.Raw)
	fmt.Println("Restricted.Signature=", user.Signature)
	fmt.Println("Restricted.Valid=", user.Valid)
	claims := user.Claims

	fmt.Println("xxxxxxx")
	fmt.Println("xxxxxxx")
	fmt.Println(claims)
	xxx := user.Claims.(*JwtCustomClaims2)
	fmt.Println(xxx)
	name := "xxx"
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
