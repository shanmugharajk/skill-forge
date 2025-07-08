package main

import (
	"fmt"
	"gpl/code/reader"
	"os"
)

func main() {
	arg := os.Args[1]

	switch arg {
	case "reader":
		reader.Run()
	default:
		fmt.Println("Invalid argument")
	}
}
