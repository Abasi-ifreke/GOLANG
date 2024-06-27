package main

import "fmt"

func main() {
	// var array1 [3]string

	// array1[0] = "okon"

	// array1[1] = "uwem"

	// array1[2] = "emmanuel"

	// fmt.Println(array1)

	// fmt.Println("the first, second and third objects in the array are:", array1[0], array1[1], array1[2])

	// array2 := [4]string{"me", "him", "them", "that"}

	// fmt.Println(array2)

	// fmt.Println("The first and second elements in the array is:", array2[0], "and", array2[1])

	// for i := 0; i < len(array2); i++ {
	// 	fmt.Println("element no.", i, "is", array2[i])
	// }

	var numbers [4]int

	fmt.Printf("%v\n", numbers)
	fmt.Printf("%#v\n", numbers)

	changeArray := [4]int{1, 3, 4, 5}

	fmt.Println(changeArray)

	changeArray[0] = 7

	fmt.Println(changeArray)
}
