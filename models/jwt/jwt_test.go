package jwt

import (
	"testing"
	"time"
)

func TestSign(t *testing.T) {
	token, err := Sign(map[string]interface{}{
		"a": 123,
		"b": 456,
		"c": time.Now(),
	})

	if err != nil {
		t.Error(err)
	}

	if len(token) == 0 {
		t.Error(token, "should return signature")
	}
}

func TestParse(t *testing.T) {
	token, _ := Sign(map[string]interface{}{
		"a": 123,
		"b": 456,
		"c": time.Now(),
	})
	ok := Verify(token)

	if !ok {
		t.Error("signature invalid")
	}
}
