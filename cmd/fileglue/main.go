package main

import (
	"fmt"

	"github.com/vladimish/fileglue/internal/core"
	"github.com/vladimish/fileglue/internal/flags"
)

func main() {
	flags.Initialize()

	err := core.Start()
	if err != nil {
		fmt.Println(err)
	}
}
