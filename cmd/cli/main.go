package main

import (
	"bulbasaur/internal/container"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		panic("invalid parameter")
	}

	fmt.Println("starting Bulbasaur")
	fmt.Println("-------------------")
	fmt.Println("")
	container.Init(args[1])
}
