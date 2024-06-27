package main

import "fmt"

func main() {

	oneArray := [...]int{3, 5, 7, 8}

	twoArray := oneArray // data is passed by value

	fmt.Printf("%v\n", oneArray)
	fmt.Printf("%v\n", twoArray)

	oneArray[0] = 1

	threeArray := &oneArray // data is passed by reference using the "&" symbol

	fmt.Printf("%v\n", oneArray)
	fmt.Printf("%v\n", twoArray)
	fmt.Printf("%v\n", *threeArray) // the "*" is used to display the pointer variable
}
