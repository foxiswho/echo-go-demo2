package jwt

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func TestXest(t *testing.T) {
	mySigningKey := []byte("AllYourBase")

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.RegisteredClaims
	}

	// Create the claims
	claims := MyCustomClaims{
		"bar",
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}

	// Create claims while leaving out some of the optional fields
	claims = MyCustomClaims{
		"bar",
		jwt.RegisteredClaims{
			// Also fixed dates can be used for the NumericDate
			ExpiresAt: jwt.NewNumericDate(time.Unix(1516239022, 0)),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", tokenString, err)

	token2, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if claims, ok := token2.Claims.(*MyCustomClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.Foo, claims.RegisteredClaims.Issuer)
	} else {
		fmt.Println(err)
	}

}

func TestMake(t *testing.T) {
	x, err := MakeJwt("xxxx", true)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(x)
	}
}
