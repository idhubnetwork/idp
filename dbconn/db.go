package dbconn

import (
	"github.com/gobuffalo/pop"
)

// GetConnection returns the singleton db connection
func GetConnection() *pop.Connection {
	pop.Debug = true

	conn, err := pop.Connect("development")

	if err != nil {
		panic(err)
	}

	return conn
}
