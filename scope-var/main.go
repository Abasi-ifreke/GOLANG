// scope of variables in Go
// scope of variable determined at compile time
// we have local and global variables (declared inside and outside a block/function)

package main

import "fmt"

var globalVar int = 100

func main() {
	var localVar int = 150

	_ = localVar

	fmt.Printf("the global variable is : %d\n", globalVar)
	fmt.Printf("the local var is : %d\n\n", localVar)

	exam()
	same()
}

func exam() {
	localVar := 50

	fmt.Printf("the local var is : %d\n", localVar)
	fmt.Printf("Local variable called globalVar in this function is : %d\n\n", globalVar)
}

// Here, global variables are called in both functions
// The local variable from exam function is the same name in the main function, yet the local variable from the exam function is the value printed.

func same() {
	globalVar := 20

	fmt.Printf("the new global variable is : %d", globalVar)
	// the local variable takes precedence over the global variable

}
