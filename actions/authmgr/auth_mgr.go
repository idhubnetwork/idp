package authmgr

import (
	"fmt"
	"idp/models/util"
	"idp/dbconn"
	"idp/models/db"
	"idp/models/jwt"
	"net/http"
	"time"

	"github.com/parnurzeal/gorequest"

	"github.com/gobuffalo/uuid"
	"github.com/labstack/echo"
)

type orgUserJoinUser struct {
	ID           uuid.UUID   `db:"id" json:"id"`
	Identity     string      `db:"identity" json:"identity"`
	Roles        string      `db:"roles" json:"roles"`
	AwsID        string      `db:"awsId" json:"awsId"`
	CreatedAt    time.Time   `db:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time   `db:"updatedAt" json:"updatedAt"`
	UserIdentity interface{} `db:"userIdentity" json:"userIdentity"`
	IsActive     *int        `db:"isActive" json:"isActive"`
}

// ApplyList will list the apply list
func ApplyList(c echo.Context) error {
	results := &[]orgUserJoinUser{}

	db := dbconn.GetConnection()

	err := db.RawQuery(`
		select
			ou.id id
		,	ou.identity identity
		,	ou.roles roles
		,	ou.aws_id awsId
		,	ou.created_at createdAt
		,	ou.updated_at updatedAt
		,	u.identity userIdentity
		,	u.is_active isActive
		from org_users ou left join users u
		on ou.identity = u.org_user
		and ou.roles = u.roles
		and u.identity = ?
		where ou.is_active = true;
	`, jwt.DecodeContext(c, "identity")).All(results)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, results)
}

// Apply will execute the apply action
func Apply(c echo.Context) error {
	user := &db.User{}
	c.Bind(user)
	err := c.Validate(user)
	user.Identity = jwt.DecodeContext(c, "identity").(string)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	db := dbconn.GetConnection()
	err = db.Create(user)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// AuthorizedList will returns the authorized list
func AuthorizedList(c echo.Context) error {
	results := &[]db.User{}

	db := dbconn.GetConnection()

	err := db.RawQuery(`
		select u.*
		from users u, org_users ou
		where u.org_user = ou.identity
		and u.roles = ou.roles
		and u.identity = ?
		and u.is_active = true;
	`, jwt.DecodeContext(c, "identity")).All(results)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, results)
}

// GetSAMLResponse returns the SAML response
func GetSAMLResponse(c echo.Context) error {
	results := &db.OrgUser{}

	db := dbconn.GetConnection()

	err := db.RawQuery(`
		select ou.*
		from users u, org_users ou
		where u.org_user = ou.identity
		and u.roles = ou.roles
		and u.id = ?;
	`, c.Param("id")).First(results)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	pl := map[string]string{
		"org_id":       results.Identity,
		"issuer":       "https://saml.idhub.network",
		"user_id":      c.Param("id"),
		"provider_arn": results.ProviderArn,
		"role_arn":     results.RoleArn,
		"org_aws_id":   results.AwsID,
	}

	fmt.Println(pl)

	cfg := &util.Cfg{}

	_, body, errs := gorequest.New().Post(cfg.GetConfig("SAMLAPI.url").String()+"/getSamlResponse").Query(pl).End()

	if errs != nil {
		return c.JSON(http.StatusBadRequest, errs)
	}

	return c.JSON(http.StatusOK, body)
}
