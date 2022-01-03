package main

import (
	"fmt"
	"github.com/vladimish/fileglue/internal/core"
)

func main() {
	err := core.Start()
	if err != nil {
		fmt.Println(err)
	}
}
