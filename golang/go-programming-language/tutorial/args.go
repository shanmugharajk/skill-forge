package tutorial

import (
	"fmt"
	"os"
)

/*
for initialization; condition; post {
	// zero or more statements
}

// While statement
for condition {
	// zero or more statements
}

// Infinite loop
for {
	// zero or more statements
}
*/

func Args() {
	for index, arg := range os.Args[1:] {
		fmt.Println(index, arg)
	}
}
