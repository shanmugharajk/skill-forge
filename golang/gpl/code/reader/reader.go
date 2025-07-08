package reader

import (
	"bufio"
	"fmt"
	"io"
)

type CustomType struct {
	source string
	offset int
}

func (sf *CustomType) Read(b []byte) (n int, err error) {
	if sf.offset >= len(sf.source) {
		return 0, io.EOF
	}

	lengthCopied := copy(b, sf.source[sf.offset:])
	sf.offset += lengthCopied
	return lengthCopied, nil
}

func usingReader() {
	fmt.Printf("\n== using io reader ==\n\n")

	reader := &CustomType{source: "hello world"}

	buffer := make([]byte, 5)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		fmt.Printf("Read %d bytes: %q\n", n, buffer[:n])
	}
}

func usingScanner() {
	fmt.Printf("\n== using bufio scanner ==\n\n")

	scanner := bufio.NewScanner(&CustomType{source: "hello world"})

	for scanner.Scan() {
		fmt.Printf("Scanned: %q\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning: %v\n", err)
	}
}

func Run() {
	usingScanner()
	usingReader()
}
