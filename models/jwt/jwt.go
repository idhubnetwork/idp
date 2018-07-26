package jwt

import (
	"fmt"
	"idp/models/util"
	"io/ioutil"

	"github.com/labstack/echo"

	"github.com/dgrijalva/jwt-go"
)

// Sign returns jwt token string
func Sign(claims map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims(claims))

	key, err := ioutil.ReadFile(util.GetAbsPath("/keystore/id_rsa"))
	priv, _ := jwt.ParseRSAPrivateKeyFromPEM(key)

	tokenString, err := token.SignedString(priv)

	return tokenString, err
}

// Verify returns the payload and is valid
func Verify(tokenString string) bool {
	key, _ := ioutil.ReadFile(util.GetAbsPath("/keystore/id_rsa.pub"))
	pub, _ := jwt.ParseRSAPublicKeyFromPEM(key)

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("%v", token.Header["alg"])
		}

		return pub, nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return ok
	}

	return false
}

// DecodeContext will decode context and return the value by key
func DecodeContext(c echo.Context, key string) interface{} {
	return c.Get("IDHub").(*jwt.Token).Claims.(jwt.MapClaims)[key]
}
