package flags

import (
	"strings"
)

// top contains top separator flag.
var top string

// bottom contains bottom separator flag.
var bottom string

// GenerateTop returns top built by the pattern defined by the flag.
func GenerateTop(path string) string {
	var res string
	res = strings.Replace(top, "@PATH", path, -1)
	res = strings.Replace(res, "\\n", "\n", -1)
	res = strings.Replace(res, "\\t", "\t", -1)
	return res
}

// GenerateBottom returns top built by the pattern defined by the flag.
func GenerateBottom(path string) string {
	var res string
	res = strings.Replace(bottom, "@PATH", path, -1)
	res = strings.Replace(res, "\\n", "\n", -1)
	res = strings.Replace(res, "\\t", "\t", -1)
	return res
}
