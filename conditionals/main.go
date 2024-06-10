package main

import "fmt"

func main() {

	a := 1

	b := 5

	// if a < b {
	// 	fmt.Println("the value of a is greater than b")
	// } else {
	// 	fmt.Println("the value of b is greater")
	// }

	// if a == 10 {

	// 	if b == 5 {

	// 		fmt.Printf("the values of a and b are %d and %d respectively", a, b)
	// 	}
	// }

	if a == b {

		fmt.Println("the value of a is same with the value of b")
	} else if a > b {

		fmt.Println("the value of a is greater than b")
	} else {
		fmt.Println("the value of a is less than b")
	}
}
