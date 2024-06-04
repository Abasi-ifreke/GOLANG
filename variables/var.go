package main

import "fmt"

// this is for declaring multiple variables of different types

func main() {
	var variable1, variable2, variable3 = 2, "yes", 47.448756

	fmt.Printf("the value of var1 is : %d\n", variable1)
	fmt.Printf("the type of var1 is : %T\n\n", variable1)

	fmt.Printf("the value of var2 is : %s\n", variable2)
	fmt.Printf("the type of var2 is : %T\n\n", variable2)

	fmt.Printf("the value of var3 is : %f\n", variable3)
	fmt.Printf("the type of var3 is : %T\n\n", variable3)

}
