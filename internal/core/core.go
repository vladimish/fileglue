package core

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"

	"github.com/vladimish/fileglue/internal/flags"
	"github.com/vladimish/fileglue/pkg/utils"
)

var path string

func Start() error {
	flag.Usage = func() {
		fmt.Println("Usage: fileglue [path] [flags]")
		fmt.Println("Flags:")
		flag.PrintDefaults()
	}

	path = flags.GetPath()
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

// printRecursively searches for matching files and prints them to stdout.
func printRecursively(files []fs.FileInfo, path string) error {
	excludeRegex := flags.GetExcludeRegex()

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
			filePath := path + f.Name()
			if excludeRegex != nil && excludeRegex.MatchString(filePath) {
				continue
			}
			// Open and read whole file.
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			data, err := ioutil.ReadAll(file)

			if utils.IsText(data) || flags.GetIgnore() {
				// Print header.
				fmt.Println(flags.GenerateTop(filePath))

				fmt.Println(string(data))

				// Print footer.
				fmt.Println(flags.GenerateBottom(filePath))
			}
			//else if flags.GetDebug() {
			//	fmt.Println("File", filePath, "skipped due to not corresponding to UTF-8 format.")
			//}
		}
	}

	return nil
}
