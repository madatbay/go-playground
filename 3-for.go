package main

import "fmt"

func main() {

	var i = 0
	fmt.Println("Basic for loop")
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("=========")

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
	fmt.Println("=========")

	for i := range 3 {
		fmt.Println(i)
	}
	fmt.Println("=========")

	for n := range 10 {
		if n == 5 {
			continue
		}

		fmt.Println(n)
	}
	fmt.Println("=========")
}
