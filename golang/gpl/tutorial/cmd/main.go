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
	case "fetch":
		tutorial.Fetch()
	case "fetchall":
		tutorial.FetchAll([]string{
			"https://jsonplaceholder.typicode.com/todos",
			"https://jsonplaceholder.typicode.com/comments",
			"https://jsonplaceholder.typicode.com/posts",
			"https://jsonplaceholder.typicode.com/albums",
			"https://jsonplaceholder.typicode.com/photos",
			"https://jsonplaceholder.typicode.com/users",
		})
	case "server":
		tutorial.Server()
	default:
		fmt.Println("Invalid argument")
	}
}
