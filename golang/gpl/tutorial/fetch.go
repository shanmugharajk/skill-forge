package tutorial

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Fetch() {
	urls := os.Args[2:]

	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("\nHttpStatus: %s\n", resp.Status)
		fmt.Printf("\nResponse body: ")

		bytesCopied, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("\n\nBytes Copied: %d\n", bytesCopied)
	}
}
