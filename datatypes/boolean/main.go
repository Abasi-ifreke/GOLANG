package main

import "fmt"

func main() {
	var1 := "Go programming"
	var2 := "GOLANG"
	var3 := "Go programming"

	result1 := var1 == var2
	result2 := var1 == var3
	result3 := var2 == var3

	fmt.Println("Result1 is: ", result1)
	fmt.Println("Result2 is: ", result2)
	fmt.Println("Result3 is: ", result3)

	fmt.Printf("the type of result1 is: %T\n", result1)
	fmt.Printf("the type of result2 is: %T\n", result2)
	fmt.Printf("the type of result3 is: %T\n", result3)
}
