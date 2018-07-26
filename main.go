package main

import (
	"idp/actions/auth"
	"idp/actions/authmgr"
	"idp/actions/orgusr"
	"idp/dbconn"
	"idp/models/util"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/gommon/log"

	"github.com/dgrijalva/jwt-go"

	"gopkg.in/go-playground/validator.v9"

	"golang.org/x/sync/errgroup"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var g errgroup.Group

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func logger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, host=${host}, method=${method}, uri=${uri}, status=${status}\n",
	})
}

func frontend(port int) http.Handler {
	e := echo.New()

	e.Use(logger())
	e.Use(middleware.Recover())

	e.Static("/", "public")
	e.File("/", "index.html")

	e.Logger.SetLevel(log.INFO)
	e.Logger.Infof("frontend start at port: %v", port)

	return e
}

func backend(port int) http.Handler {
	e := echo.New()

	e.Validator = &customValidator{validator.New()}
	key, _ := ioutil.ReadFile(util.GetAbsPath("/keystore/id_rsa.pub"))
	pub, _ := jwt.ParseRSAPublicKeyFromPEM(key)

	e.Use(logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://127.0.0.1:8080"},
		AllowCredentials: true,
	}))
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper: func(c echo.Context) bool {
			if strings.HasPrefix(c.Path(), "/auth/") {
				return true
			}

			return false
		},
		SigningKey:    pub,
		SigningMethod: jwt.SigningMethodRS256.Name,
		TokenLookup:   "cookie:IDHUB_JWT",
		ContextKey:    "IDHub",
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "後端…1323")
	})

	e.POST("/auth/verify", auth.Verify)
	e.POST("/auth/booking", auth.Booking)
	e.GET("/auth/logout", auth.Logout)

	e.GET("/org-user/pending-list", orgusr.PendingList)
	e.POST("/org-user/create", orgusr.Create)
	e.POST("/org-user/approval", orgusr.Approval)

	e.GET("/auth-manager/authorized-list", authmgr.AuthorizedList)
	e.GET("/auth-manager/apply-list", authmgr.ApplyList)
	e.POST("/auth-manager/apply", authmgr.Apply)
	e.GET("/auth-manager/get-saml-response/:id", authmgr.GetSAMLResponse)

	e.Logger.SetLevel(log.INFO)
	e.Logger.Infof("backend start at port: %v", port)

	return e
}

func main() {
	db := dbconn.GetConnection()
	defer db.Close()

	frontend := &http.Server{
		Addr:         ":8080",
		Handler:      frontend(8080),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	backend := &http.Server{
		Addr:         ":1312",
		Handler:      backend(1312),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return frontend.ListenAndServe()
	})

	g.Go(func() error {
		return backend.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
