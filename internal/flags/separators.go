package flags

import (
	"flag"
	"github.com/rsc/getopt"
	"strings"
)

// top contains top flag.
var top string

// bottom contains bottom flag.
var bottom string

func init() {
	initHeaderFlag()
	initFooterFlag()

	// Need to be called after all flags being defined.
	getopt.Parse()
}

func initHeaderFlag() {
	const (
		value = "@PATH:"
		usage = "Top is a text that will be printed at the top of any file." +
			"Use '@PATH' expression to insert file path." +
			"e.g. === @PATH ==="
	)
	flag.StringVar(&top, "top", value, usage)
	getopt.Alias("t", "top")
}

func initFooterFlag() {
	const (
		value = "==="
		usage = "Footer is a text that will be printed at the bottom of any file." +
			"Use '@PATH' expression to insert file path." +
			"e.g. === @PATH ==="
	)
	flag.StringVar(&bottom, "bottom", value, usage)
	getopt.Alias("b", "bottom")
}

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
