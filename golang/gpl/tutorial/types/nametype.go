package types

import "fmt"

type Celsius float64

type Fahrenheit float64

func init() {
	fmt.Printf("\n== type init package ==\n\n")
}

func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func namedTypeUsage() {
	fmt.Printf("\n\n== named type usage ==\n")

	var c Celsius = 200
	f := c.ToFahrenheit()

	fmt.Printf("%v°C is %v°F\n", c, f)
}
