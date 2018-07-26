package util

import (
	"strings"
	"testing"
)

func TestGetAbsPath(t *testing.T) {
	if !strings.HasSuffix(GetAbsPath(), "idp") {
		t.Error("invalid path")
	}

	path := GetAbsPath("/foo", "/bar")
	if strings.Index(path, "/foo") < 0 || strings.Index(path, "/bar") < 0 {
		t.Error("invalid path")
	}
}
