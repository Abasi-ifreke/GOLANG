package main

import "fmt"

// go is a statically typed language. it compiles the variables/values before they are run
// go can detect the type of variable (integer, float, string, boolea, pointers, etc) before compilation

func main() {
	var a = 4
	var b = 5.2

	// var y int = "5"   // this wont run because the type was stated as int, but it is assigned a string literal, not an int literal

	// a = b  //here, a is int, b is float. the assignment wont work.

	a = int(b) //this will work because the type is explicitly stated.

	fmt.Println(a, b)

	// the default zero values are
	// - numeric types = 0
	// - bool types = false
	// - string type = '' (empty string)
	// - pointer type = nil
}
