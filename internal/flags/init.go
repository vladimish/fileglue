package flags

import (
	"flag"
	"os"
	"regexp"
)

func initHeaderFlag() {
	const (
		value = "@PATH:"
		usage = "Top is a text that will be printed at the top of any file." +
			"Use '@PATH' expression to insert file path." +
			"e.g. === @PATH ==="
	)
	flag.StringVar(&top, "top", value, usage)
}

func initFooterFlag() {
	const (
		value = "==="
		usage = "Footer is a text that will be printed at the bottom of any file." +
			"Use '@PATH' expression to insert file path." +
			"e.g. === @PATH ==="
	)
	flag.StringVar(&bottom, "bottom", value, usage)
}

func initPathFlag() {
	const (
		value = "." + string(os.PathSeparator)
		usage = "Path to the directory."
	)
	flag.StringVar(&path, "path", value, usage)
}

func initIgnoreFlag() {
	const (
		value = false
		usage = "By default, fileglue declines all the files which contains byte of value less then 40 except some escape sequences. Passing this flag will allow to read all files."
	)
	flag.BoolVar(&ignore, "ignore", value, usage)
}

func initExcludeFlag() {
	const (
		value = ""
		usage = "Specify a regex pattern to exclude file paths."
	)
	flag.StringVar(&exclude, "exclude", value, usage)
}

func Initialize() {
	initPathFlag()
	initHeaderFlag()
	initFooterFlag()
	initIgnoreFlag()
	initExcludeFlag()

	flag.Parse()

	if exclude != "" {
		excludeRegex = regexp.MustCompile(exclude)
	}
}
