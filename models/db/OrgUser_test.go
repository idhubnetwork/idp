package db

import (
	"idp/dbconn"
	"testing"

	"github.com/gobuffalo/uuid"
)

func TestCreateOrgUser(t *testing.T) {
	db := dbconn.GetConnection()
	err := db.Create(&OrgUser{
		Identity:    "0x123",
		Roles:       "Jamie",
		ProviderArn: "arn:aws:iam::000000000000:saml-provider/IDHub",
		RoleArn:     "arn:aws:iam::000000000000:role/IDHUB_IDP",
		AwsID:       "000000000000",
		IsActive:    true,
	})

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateOrgUser(t *testing.T) {
	db := dbconn.GetConnection()
	err := db.Save(&OrgUser{
		Identity:    "0x123",
		Roles:       "Jamie is Rich",
		ProviderArn: "arn:aws:iam::000000000000:saml-provider/IDHub",
		RoleArn:     "arn:aws:iam::000000000000:role/IDHUB_IDP",
		AwsID:       "111111111111",
		IsActive:    false,
	})

	if err != nil {
		t.Error(err)
	}
}

func TestJoin(t *testing.T) {
	type orgUserJoinUser struct {
		ID       uuid.UUID `db:"id"`
		Identity string    `db:"identity"`
		Roles    string    `db:"roles"`
		AwsID    string    `db:"awsId"`
		IsActive bool      `db:"isActive"`
		ID2      uuid.UUID `db:"id2"`
		Msg      string    `db:"msg"`
	}

	result := &[]orgUserJoinUser{}

	db := dbconn.GetConnection()

	err := db.RawQuery(`
		select ou.id id, s.id id2, s.msg msg, ou.identity identity, ou.roles roles, ou.aws_id awsId, ou.is_active isActive
		from org_users ou, signins s
		where ou.identity = s.identity;
		`).All(result)

	if err != nil {
		t.Error(err)
	}

	if len(*result) == 0 {
		t.Error("no data")
	}
}

func TestLeftJoin(t *testing.T) {
	type orgUserJoinUser struct {
		ID       uuid.UUID     `db:"id"`
		Identity string        `db:"identity"`
		Roles    string        `db:"roles"`
		AwsID    string        `db:"awsId"`
		IsActive int           `db:"isActive"`
		ID2      uuid.NullUUID `db:"id2"`
	}

	result := &[]orgUserJoinUser{}

	db := dbconn.GetConnection()

	err := db.RawQuery(`
		select ou.id id, u.id id2, ou.identity identity, ou.roles roles, ou.aws_id awsId, ou.is_active isActive
		from org_users ou left join users u
		on ou.identity = u.org_user;
		`).All(result)

	if err != nil {
		t.Error(err)
	}

	for _, res := range *result {
		v, _ := res.ID2.Value()

		if v != nil {
			t.Error(v, "is not nil")
		}
	}
}
