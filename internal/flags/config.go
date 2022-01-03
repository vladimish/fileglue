package flags

import (
	"flag"
	"github.com/rsc/getopt"
)

var debug bool
var ignore bool

func init() {
	initDebugFlag()
	initIgnoreFlag()
}

func initDebugFlag() {
	const (
		value = false
		usage = "Print debug information."
	)
	flag.BoolVar(&debug, "debug", value, usage)
	getopt.Alias("d", "debug")
}

func initIgnoreFlag() {
	const (
		value = false
		usage = "By default, fileglue declines all the files which contains byte of value less then 40 except some escape sequences. Passing this flag will allow to read all files."
	)
	flag.BoolVar(&ignore, "ignore", value, usage)
	getopt.Alias("i", "ignore")
}

func GetDebug() bool {
	return debug
}

func GetIgnore() bool {
	return ignore
}
