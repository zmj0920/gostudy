package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello world")

	if len(os.Args) > 1 {
		fmt.Println("Hello World", os.Args[0])
	}

	fmt.Println()
	os.Exit(0)
}
