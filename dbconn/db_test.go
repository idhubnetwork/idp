package dbconn

import (
	"testing"
)

func TestGetConnection(t *testing.T) {
	conn1 := GetConnection()
	conn2 := GetConnection()

	if conn1 == nil {
		t.Error("conn1 is", conn1)
	}

	if conn2 == nil {
		t.Error("conn2 is", conn2)
	}

	if conn1 != conn2 {
		t.Errorf("conn1(%p) and conn2(%p) are not the same instance", conn1, conn2)
	}
}
