package types

import "fmt"

type Employee struct {
	Name string
	Age  int
}

// Implement Stringer interface
func (e Employee) String() string {
	return fmt.Sprintf("Employee{Name: %s, Age: %d}", e.Name, e.Age)
}

func updateName(e Employee) {
	e.Name = "updated"
	fmt.Println(e)
}

func addToArray(arr []int) {
	arr[2] = 9
}

func structValueTypeUsage() {
	fmt.Printf("\n\n== struct value type usage ==\n\n")

	shan := Employee{}
	shan.Name = "shan"
	shan.Age = 40
	updateName(shan)
	fmt.Printf("\nshan didn't get updated since its value type: \n%v\n", shan)

	arr := []int{1, 2, 3}
	fmt.Printf("\narr before update: \n%v", arr)
	addToArray(arr)
	fmt.Printf("\narr got updated since its reference type: \n%v\n", arr)
}
