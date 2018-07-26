package db

import (
	"idp/dbconn"
	"time"

	"github.com/gobuffalo/uuid"
)

// Signin struct
type Signin struct {
	ID        uuid.UUID `db:"id"`
	Identity  string    `db:"identity"`
	Msg       string    `db:"msg"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// GetBookingMsg return the booking message from db
func GetBookingMsg(identity string) (string, error) {
	db := dbconn.GetConnection()
	signin := &Signin{}
	msg := "IDHub: 0x29F20242051AccDA50D52a7E272A5F23237e4696 @ " + time.Now().String()
	var err error

	err = db.Where("identity = ?", identity).First(signin)

	signin.Identity = identity
	signin.Msg = msg

	err = db.Save(signin)

	return msg, err
}

// GetVerifyMsg return what message you should use
func GetVerifyMsg(identity string) (string, error) {
	db := dbconn.GetConnection()
	signin := &Signin{}

	err := db.Where("identity = ?", identity).First(signin)

	return signin.Msg, err
}
