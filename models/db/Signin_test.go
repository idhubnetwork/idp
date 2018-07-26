package db

import (
	"idp/dbconn"
	"testing"
)

func Test_select(t *testing.T) {
	db := dbconn.GetConnection()

	signin := &[]Signin{}
	err := db.All(signin)

	if err != nil {
		t.Error(err)
	}
}
func Test_insert(t *testing.T) {
	db := dbconn.GetConnection()

	signin := &Signin{}
	signin.Identity = "0x9487"
	signin.Msg = "ㄎㄎ"

	err := db.Create(signin)

	if err != nil {
		t.Error(err)
	}
}

func Test_update(t *testing.T) {
	db := dbconn.GetConnection()

	signin := &Signin{}

	err := db.Where("identity = ?", "0x9487").First(signin)

	if err != nil {
		t.Error(err)
	}

	signin.Msg = "嘖嘖"

	err = db.Update(signin)

	if err != nil {
		t.Error(err)
	}
}

func Test_delete(t *testing.T) {
	db := dbconn.GetConnection()

	signin := &Signin{}

	err := db.Where("identity = ?", "0x9487").First(signin)

	if err != nil {
		t.Error(err)
	}

	err = db.Destroy(signin)

	if err != nil {
		t.Error(err)
	}
}

func TestGetBookingMsg(t *testing.T) {
	msg1, err := GetBookingMsg("0x9487")

	if err != nil {
		t.Error(err)
	}

	msg2, err := GetBookingMsg("0x9487")

	if err != nil {
		t.Error(err)
	}

	if msg1 == msg2 {
		t.Error(msg1, "and", msg2, "are the same")
	}
}

func TestGetVerifyMsg(t *testing.T) {
	_, err := GetVerifyMsg("0x9487")

	if err != nil {
		t.Error(err)
	}
}
