package main

import "fmt"

func main() {
	var a = 5
	var b = 10

	if a+b == 15 {
		fmt.Println("Answer is 15")
	} else {
		fmt.Println("Answer is not 15")
	}

	if a == 5 && b == 10 {
		fmt.Println("AND statement passed")
	} else if a == 5 || b == 10 {
		fmt.Println("OR statemtn passed")
	}

	// Variables declated in conditional are only locally accessible inside conditional block
	if d := 3; d > 5 {
		fmt.Println("D gt 5. d=", d)
	} else {
		fmt.Println("D lte 5. d=", d)
	}
}
