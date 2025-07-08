package tutorial

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func FetchAll(urls []string) {
	fmt.Printf("\n== fetching urls ==\n\n")

	ch := make(chan string, len(urls))

	start := time.Now()

	for _, url := range urls {
		go fetch(url, ch)
	}

	for range urls {
		// It will read one by value from channel. This will block until a value is available.
		fmt.Println(<-ch)
	}

	secs := time.Since(start).Seconds()
	fmt.Printf("\n%.2fs elapsed\n", secs)
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%s %s", url, err)
		return
	}

	bytesCopied, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("%s %s", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%s %7d (bytes) 	%.2fs\n", url, bytesCopied, secs)
}
