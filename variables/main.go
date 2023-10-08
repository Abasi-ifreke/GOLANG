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
}
