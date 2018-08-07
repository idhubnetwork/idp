package cfg

import (
	"idp/models/util"
	"net/http"

	"github.com/labstack/echo"
)

// GetConfig function
func GetConfig(c echo.Context) error {
	cfg := &util.Cfg{}

	return c.String(http.StatusOK, cfg.GetConfig(c.Param("key")).String())
}
