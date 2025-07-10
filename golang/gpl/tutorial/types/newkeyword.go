package types

import "fmt"

func newKeywordUsage() {
	// Case 1: Standard types (non-zero size)
	// Each call to new for a non-zero size type returns a distinct address.
	fmt.Printf("\n\n== new(int) == new(int)\n")
	pInt := new(int)
	qInt := new(int)
	fmt.Printf("pInt: %p, qInt: %p\n", pInt, qInt)
	fmt.Println("pInt == qInt:", pInt == qInt) // Output: false (always)

	// Case 2: Zero-size type - struct{}
	// new(struct{}) might return the same address depending on implementation.
	fmt.Printf("\n\n== new(struct{}) == new(struct{})\n")
	pStruct := new(struct{})
	qStruct := new(struct{})
	fmt.Printf("pStruct: %p, qStruct: %p\n", pStruct, qStruct)
	fmt.Println("pStruct == qStruct:", pStruct == qStruct) // Output: true or false (implementation-dependent)
}
