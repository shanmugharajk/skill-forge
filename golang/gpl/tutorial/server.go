package tutorial

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

var count int64

var mu sync.Mutex

func Server() {
	http.HandleFunc("/", baseHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/count", countHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func baseHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("GET /")

	mu.Lock()
	count++
	mu.Unlock()

	// Write to response body
	res.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(res, "%s %s %s\n", req.Method, req.URL, req.Proto)

	for k, v := range req.Header {
		fmt.Fprintf(res, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(res, "Host = %q\n", req.Host)
	fmt.Fprintf(res, "RemoteAddr = %q\n", req.RemoteAddr)

	if err := req.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range req.Form {
		fmt.Fprintf(res, "Form[%q] = %q\n", k, v)
	}
}

func countHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("GET /count")

	mu.Lock()
	fmt.Fprintf(res, "%d\n", count)
	mu.Unlock()
}

func faviconHandler(res http.ResponseWriter, req *http.Request) {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Construct the full path to the favicon
	// Assuming we run from the folder 'gpl'. This is a temp fix just for now.
	faviconPath := filepath.Join(currentDir, "tutorial/resources/favicon.png")

	file, err := os.Open(faviconPath)
	if err != nil {
		// If favicon not found, return 404 instead of crashing
		http.Error(res, "Not Found", http.StatusNotFound)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	res.Header().Set("Content-Type", "image/png")
	res.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))
	io.Copy(res, file)
}
