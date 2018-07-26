package orgusr

import (
	"idp/dbconn"
	"idp/models/db"
	"idp/models/jwt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// Create will create an OrgUser record
func Create(c echo.Context) error {
	ou := &db.OrgUser{}
	c.Bind(ou)

	err := c.Validate(ou)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	orgUser := &db.OrgUser{
		Identity:    jwt.DecodeContext(c, "identity").(string),
		Roles:       ou.Roles,
		ProviderArn: ou.ProviderArn,
		RoleArn:     ou.RoleArn,
		AwsID:       ou.AwsID,
		IsActive:    true,
	}

	db := dbconn.GetConnection()
	err = db.Create(orgUser)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// PendingList will returns the pending list by org. user
func PendingList(c echo.Context) error {
	users := &[]db.User{}

	db := dbconn.GetConnection()
	err := db.Where("org_user = ?", jwt.DecodeContext(c, "identity")).Where("is_active = false").All(users)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

// Approval will approval the pending list
func Approval(c echo.Context) error {
	type user struct {
		ID        string    `db:"id" json:"id" form:"id" validate:"required"`
		IsActive  int       `db:"is_active" json:"isActive" form:"isActive" validate:"required"`
		UpdatedAt time.Time `db:"updated_at"`
	}

	rp := &user{}
	c.Bind(rp)

	err := c.Validate(rp)

	rp.UpdatedAt = time.Now()

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	db := dbconn.GetConnection()
	err = db.Update(rp)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
