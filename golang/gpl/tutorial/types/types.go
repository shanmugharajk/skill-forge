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
	case "nametype":
		namedTypeUsage()
	case "scope":
		scopeUsage()
	default:

		// Floating point preciion example
		// TODO: read more about it
		var f float32 = 16777216
		fmt.Println(f == f+1)

		/*	== In js this is the same example ==
			let f_js = 9007199254740992; // 2^53
			console.log(f_js === f_js + 1); // true
		*/

		fmt.Println("Invalid argument")
	}
}
