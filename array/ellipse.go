package main

import "fmt"

// using ellipse and not specifying the length of the array is the same thing
// go compiler will iddentify the length of the array based on the elements specfied in the array declaration.

func main() {
	// array3 := [...]int{2, 3, 4, 5, 6, 7}

	// fmt.Println("the length of the array is,", len(array3))

	// you can also initialise elements in an array using the syntax
	// i.e you get to choose what each element will be and in what position in the array

	array4 := [6]int{0: 5, 1: 3, 5: 2, 3: 88}

	array5 := [...]string{0: "okon", 1: "uwem", 2: "guy", 3: "boy", 4: "charles", 5: "yes"}

	fmt.Println(array4)
	// fmt.Println(array5)

	// you can also filter elements in the array using colon (:)

	fmt.Println(array5[:])   // prints out all elements in the array
	fmt.Println(array5[:2])  // prints out the first two elements in the array excluding the 2nd index element
	fmt.Println(array5[2:])  // prints out elements starting from the 2nd index
	fmt.Println(array5[1:4]) // prints out elements starting from the 1st index to the 4th index but not including the 4th index
}
