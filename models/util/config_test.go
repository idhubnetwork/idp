package util

import (
	"testing"
)

func TestGetConfig(t *testing.T) {
	c := &Cfg{}

	t.Log(c.GetConfig("SAMLAPI.url"))
	t.Log(c.GetConfig("CORSConfig.allowOrigins").List())
	t.Log(c.GetConfig("CORSConfig.allowOrigins"))
	t.Log(c.GetConfig("foo"))
	t.Log(c.GetConfig("foo").String())
}
