package main

import (
	"fmt"
	"gpl/tutorial"
	"os"
)

func main() {
	arg := os.Args[1]

	switch arg {
	case "args":
		tutorial.Args()
	case "map":
		tutorial.Map()
	default:
		fmt.Println("Invalid argument")
	}
}
