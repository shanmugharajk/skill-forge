package types

import "fmt"

func scopeExample2() {
	x := "hello!"

	for i := 0; i < len(x); i++ {
		fmt.Printf("\n\nouter scope = %s\n", x)

		x := x[i]

		fmt.Printf("instead outer scope printing inner scope => %c\n", x)

		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("adding chars = %c", x) // "HELLO" (one letter per iteration)
		}
	}
}

func scopeExample1() {
	fmt.Println()

	x := "hello"

	for _, x := range x { // This X is implicit
		x := x + 'A' - 'a'
		fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
	}

	fmt.Println()
}

func scopeUsage() {
	scopeExample2()
	scopeExample1()
}
