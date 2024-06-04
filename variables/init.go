package main

import "fmt"

func main() {
	var variable1 = 100
	var variable2 = "hello"
	var variable3 = 3.142

	variable4 := "yes"
	variable5, variable6 := 7, 8
	var variable7, variable8 int = 20, 30

	fmt.Printf("the value of variable1 : %d\n", variable1)  // %d is for printing integer using printf
	fmt.Printf("the type of variable1 : %T\n\n", variable1) // %T is for printing the Type

	fmt.Printf("value of variable2 : %s\n", variable2) // %s is for printing string variables
	fmt.Printf("type of variable2 : %T\n\n", variable2)

	fmt.Printf("value of variable3 : %f\n", variable3) // %f is for printing float type variable
	fmt.Printf("type of variable3 : %T\n\n", variable3)

	fmt.Printf("value of variable4 : %s\n", variable4)
	fmt.Printf("type of variable4 : %T\n\n", variable4)

	fmt.Printf("value of the variable5 : %d\n", variable5)
	fmt.Printf("value of variable6 : %d\n\n", variable6)

	fmt.Printf("the value of variable7 is : %d\n", variable7)
	fmt.Printf("the value of variable8 is : %d\n", variable8)
}
