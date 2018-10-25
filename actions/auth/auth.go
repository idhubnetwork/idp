package auth

import (
	"errors"
	"fmt"
	"idp/models/crypto"
	"idp/models/db"
	"idp/models/jwt"
	"idp/models/resolver"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
)

type verify struct {
	Msg  string `json:"msg" form:"msg"`
	Sig  string `json:"sig" form:"sig" validate:"required"`
	Addr string `json:"addr" form:"addr"`
}

func verifyAuth(msg, id, sig string) (info string, err error) {
	addr, err := crypto.EcRecover(msg, sig)
	fmt.Println("签名地址：" + addr)
	r, err := resolver.NewResolver("infuraRopsten", "0x1DbF8e4B47EA53a2b932850F7FEC8585C6A9c1d2")
	owner, err := r.IdentityOwner(id)
	fmt.Println("Owner地址：" + owner)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	publickey, err := crypto.SigPublicKey(msg, sig)
	fmt.Println("公钥：" + publickey)
	ok, err := r.ValidAuthentication(id, "sigAuth", publickey)
	fmt.Println("公钥验证结果如下：")
	fmt.Println(ok)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if strings.ToLower(addr) == strings.ToLower(id) {
		return "sign by self", nil
	} else if strings.ToLower(owner) == strings.ToLower(addr) {
		return "sign by did owner", nil
	} else if ok {
		return "sign with authKey", nil
	}
	return "", errors.New("verify failed")
}

// Verify response signed jwt token, if pass
func Verify(c echo.Context) (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			c.SetCookie(&http.Cookie{
				Name:     "IDHUB_JWT",
				HttpOnly: true,
				Path:     "/",
				MaxAge:   -1,
			})

			c.SetCookie(&http.Cookie{
				Name:     "IDHUB_IDENTITY",
				HttpOnly: false,
				Path:     "/",
				MaxAge:   -1,
			})

			err = c.String(http.StatusUnauthorized, r.(error).Error())
		}
	}()

	v := new(verify)
	fmt.Println(v)
	c.Bind(v)

	err = c.Validate(v)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	msg, err := db.GetVerifyMsg(v.Addr)

	info, err := verifyAuth(msg, v.Addr, v.Sig)
	fmt.Println(info)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	tokenString, err := jwt.Sign(map[string]interface{}{
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(30 * time.Minute).Unix(),
		"iss":      "IDHub IdP",
		"sub":      "IDHub identity is all your life",
		"identity": v.Addr,
	})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	c.SetCookie(&http.Cookie{
		Name:     "IDHUB_JWT",
		Value:    tokenString,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(30 * time.Minute),
	})

	c.SetCookie(&http.Cookie{
		Name:     "IDHUB_IDENTITY",
		Value:    v.Addr,
		HttpOnly: false,
		Path:     "/",
		Expires:  time.Now().Add(30 * time.Minute),
	})

	// return c.String(http.StatusOK, addr)
	return c.NoContent(http.StatusOK)
}

// Booking response the register message
func Booking(c echo.Context) error {
	v := &verify{}
	c.Bind(v)

	msg, err := db.GetBookingMsg(v.Addr)

	if err != nil {
		return c.String(http.StatusNotAcceptable, err.Error())
	}

	return c.String(http.StatusOK, msg)
}

// Logout will delete the cookie
func Logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:     "IDHUB_JWT",
		HttpOnly: true,
		Path:     "/",
		MaxAge:   -1,
	})

	c.SetCookie(&http.Cookie{
		Name:     "IDHUB_IDENTITY",
		HttpOnly: false,
		Path:     "/",
		MaxAge:   -1,
	})

	return c.NoContent(http.StatusOK)
}
