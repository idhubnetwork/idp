package util

import (
	"os"
	"strings"
)

// GetAbsPath return the root path
func GetAbsPath(dirs ...string) string {
	root := "/idp"
	dir, _ := os.Getwd()
	path := dir[:strings.Index(dir, root)] + root

	for _, dir := range dirs {
		path += dir
	}

	return path
}
