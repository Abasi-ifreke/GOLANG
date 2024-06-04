// Declaring multiple variables in a single block

package main

import "fmt"

func main() {
	var (
		firstName = "Kim"
		num1      = 5
		num2      = 4.5678
		eaten     = false
	)

	fmt.Println(firstName)
	fmt.Println(num1)
	fmt.Println(eaten)
	fmt.Printf("var num2 is : %f", num2)
}
