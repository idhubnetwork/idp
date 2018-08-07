package util

import (
	"io/ioutil"
	"strings"

	"github.com/olebedev/config"
)

// Cfg type is *config.Config
type Cfg struct {
	config *config.Config
}

var cfg *config.Config

func init() {
	f, err := ioutil.ReadFile(GetAbsPath("/config.yml"))

	if err == nil {
		cfg, _ = config.ParseYaml(string(f))
	}
}

// GetConfig function
func (c *Cfg) GetConfig(key string) *Cfg {
	if cfg != nil {
		config, err := cfg.Get(key)

		if err != nil {
			c.config = nil
		}

		c.config = config

		return c
	}

	return c
}

func (c *Cfg) String() string {
	if c.config == nil {
		return "<nil>"
	}

	switch c.config.Root.(type) {
	case string:
		return c.config.Root.(string)
	case []interface{}:
		return strings.Join(c.List(), " ")
	default:
		return "<nil>"
	}
}

// List function
func (c *Cfg) List() []string {
	res := []string{}

	cfgs, ok := c.config.Root.([]interface{})

	if ok {
		for _, cfg := range cfgs {
			res = append(res, cfg.(string))
		}
	}

	return res
}
