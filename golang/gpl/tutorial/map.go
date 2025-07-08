package tutorial

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(file *os.File, counts map[string]int) {
	// Scanner accepts the reader here. File has the reader interface
	input := bufio.NewScanner(file)

	// ctrl + D to exit the scan loop here
	for input.Scan() {
		counts[input.Text()]++
	}

	// print the map
	for line, index := range counts {
		fmt.Printf("%d \t %s\n", index, line)
	}
}

func Map() {
	counts := make(map[string]int)
	files := os.Args[2:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
		return
	}

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
			continue
		}

		countLines(f, counts)
		f.Close()
	}
}
