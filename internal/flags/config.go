package flags

import (
	"regexp"
)

var path string
var ignore bool
var exclude string
var excludeRegex *regexp.Regexp

func GetPath() string {
	return path
}

func GetIgnore() bool {
	return ignore
}

func GetExcludeRegex() *regexp.Regexp {
	return excludeRegex
}
