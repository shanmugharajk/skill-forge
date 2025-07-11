package types

import (
	"fmt"
	"os"
)

func Run() {
	arg := os.Args[2]

	switch arg {
	case "newkeyword":
		newKeywordUsage()
	case "set":
		setExample()
	case "async":
		executeAsyncTask()
	case "embedding":
		embeddingStructExample()
	case "structvaluetype":
		structValueTypeUsage()
	default:
		fmt.Println("Invalid argument")
	}
}
