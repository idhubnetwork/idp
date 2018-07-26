package db

import (
	"time"

	"github.com/gobuffalo/uuid"
)

// OrgUser struct
type OrgUser struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Identity    string    `db:"identity" json:"identity" form:"identity"`
	Roles       string    `db:"roles" json:"roles" form:"roles" validate:"required"`
	ProviderArn string    `db:"provider_arn" json:"providerArn" form:"providerArn" validate:"required"`
	RoleArn     string    `db:"role_arn" json:"roleArn" form:"roleArn" validate:"required"`
	AwsID       string    `db:"aws_id" json:"awsId" form:"awsId" validate:"required"`
	IsActive    bool      `db:"is_active" json:"isActive"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time `db:"updated_at" json:"updatedAt"`
}
