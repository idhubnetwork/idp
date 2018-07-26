package db

import (
	"time"

	"github.com/gobuffalo/uuid"
)

// User struct
type User struct {
	ID        uuid.UUID `db:"id" json:"id"`
	OrgUser   string    `db:"org_user" json:"orgUser" form:"orgUser" validate:"required"`
	Roles     string    `db:"roles" json:"roles" form:"roles" validate:"required"`
	Identity  string    `db:"identity" json:"identity" form:"identity"`
	IsActive  int       `db:"is_active" json:"isActive"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}
