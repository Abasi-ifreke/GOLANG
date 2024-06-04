package main

import "fmt"

func main() {
	var age int = 40
	fmt.Println("Age:", age)

	// another way of declaring variables in go is by using ":="
	name := "Johnson"
	fmt.Println("your name is: ", name)

	// you can declare and not use variables in go using "_" It will run without any errors
	position := "1st"
	_ = position

	// multiple declaration
	var (
		number    int
		girl      bool
		firstname string
	)
	fmt.Println(number, girl, firstname)

	open, man := true, false
	open, man = man, open
	close, woman := false, true
	_, _ = close, woman
	println(open, man)

	var a, b, c int
	fmt.Println(a, b, c)

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

	// this is for declaring multiple variables of different types

	var var1, var2, var3 = 2, "yes", 47.448756

	fmt.Printf("the value of var1 is : %d\n", var1)
	fmt.Printf("the type of var1 is : %T\n\n", var1)

	fmt.Printf("the value of var2 is : %s\n", var2)
	fmt.Printf("the type of var2 is : %T\n\n", var2)

	fmt.Printf("the value of var3 is : %f\n", var3)
	fmt.Printf("the type of var3 is : %T\n\n", var3)
}
