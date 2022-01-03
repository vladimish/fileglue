package core

import (
	"flag"
	"fmt"
	"github.com/vladimish/fileglue/internal/flags"
	"github.com/vladimish/fileglue/pkg/utils"
	"io/fs"
	"io/ioutil"
	"os"
)

var path string

func Start() error {
	checkPath()
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	err = printRecursively(files, path)
	if err != nil {
		return err
	}
	return nil
}

// checkPath checks if path is empty and making it './' if so.
func checkPath() {
	path = flag.Arg(0)
	if path == "" {
		path = "." + string(os.PathSeparator)
	}
	if flags.GetDebug() {
		fmt.Println(path)
	}
}

// printRecursively searches for matching files and prints them to stdout.
func printRecursively(files []fs.FileInfo, path string) error {
	// Iterate over all files and directories inside current directory.
	for _, f := range files {
		if f.IsDir() {
			newFiles, err := ioutil.ReadDir(path + f.Name())
			if err != nil {
				return err
			}
			err = printRecursively(newFiles, path+f.Name()+string(os.PathSeparator))
			if err != nil {
				return err
			}
		} else {
			// Open and read whole file.
			file, err := os.Open(path + f.Name())
			if err != nil {
				return err
			}
			data, err := ioutil.ReadAll(file)

			if utils.IsText(data) || flags.GetIgnore() {
				// Print header.
				fmt.Println(flags.GenerateTop(path + f.Name()))

				fmt.Println(string(data))

				// Print footer.
				fmt.Println(flags.GenerateBottom(path + f.Name()))
			} else if flags.GetDebug() {
				fmt.Println("File", path+f.Name(), "skipped due to not corresponding to UTF-8 format.")
			}
		}
	}

	return nil
}
