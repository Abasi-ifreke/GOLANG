// showing how to iterate over an array

package main

import (
	"fmt"
	"strings"
)

func main() {
	// intArray := [5]string{"boy", "girl", "man", "woman", "this"}

	// for i := 0; i < len(intArray); i++ {
	// 	fmt.Println(intArray[i])
	// }

	//using range to iterate over an array
	// range is just a keyword used for iteration

	arrayVar := [5]int{1, 2, 3, 4, 5}

	fmt.Println(arrayVar)

	arrayVar[0], arrayVar[1], arrayVar[2] = 2, 4, 6
	arrayVar[3] = 8
	arrayVar[4] = 10

	for i, j := range arrayVar {
		fmt.Println("the index is:", i, "and the value is:", j)
	}

	fmt.Println(strings.Repeat("#", 20))

	// we can create a multidimensional array (3D array)
	arrayString := [2][3]string{
		{"okon", "uwem", "abas"}, // This creates a 2X3 array - 2 rows and 3 columns
		{"ifreke", "freke", "reke"},
	}

	fmt.Println(arrayString)

	// Equality of an array
	x := [3]int{2, 3, 5}

	y := x

	fmt.Println("array x is equal to array y:", x == y)
	fmt.Println(y)

	fmt.Println(strings.Repeat("#", 20))

	// ARRAYS WITH KEYED ELEMENTS
	// Creating and assigning elements with indexes

	// specifying index and elements for an array
	grades := [3]int{
		1: 15,
		2: 30,
		0: 0,
	}
	fmt.Println(grades)

	secondGrades := [...]int{3: 10}
	fmt.Println(secondGrades, len(secondGrades))

	fmt.Println(strings.Repeat("#", 30))

	// what happens when you dont specify the index??
	cities := [...]string{
		1: "Lagos",
		"port harcourt",
		3:        "uyo",
		"ibadan", // the element that isnt keyed (without an index) gets its index from the last keyed element (uyo, which has an index of 2) in the array. Its index becomes 3

	}
	fmt.Printf("%#v\n", cities)

	fmt.Println(strings.Repeat("#", 40))

	//Iterating through days of the week using range and for loop

	week := [...]string{
		0: "Monday",
		1: "Tuesday",
		2: "Wednesday",
		3: "Thursday",
		4: "Friday",
		5: "Saturday",
		6: "Sunday",
	}

	// for i := 0; i < len(week); i++ {
	// 	if i < 5 {
	// 		fmt.Println("today is", week[i], "still a weekday")
	// 	} else {
	// 		fmt.Println("today is", week[i], "weekend")
	// 	}
	// }

	// The i will print out the index
	// for i, j := range week {
	// 	if i < 5 {
	// 		fmt.Println("today is:", j, "and its still weekday")
	// 	} else {
	// 		fmt.Println("its weekend:", j)
	// 	}
	// }

	// we could also use the range keyword without printing the index
	// for _, j := range week {
	// 	fmt.Println("Today is", j)
	// }

	// another pattern we could use together with the range is using for loop as a while loop
	a := 0

	for range week {
		if a < 5 {
			fmt.Println("today", week[a], "is still a week day")
		} else {
			fmt.Println("today", week[a], "is weekend")
		}
		a++
	}
}
