package main

import (
	"fmt"
	"time"
)

// func main() {

// 	// a := 1

// 	// b := 5

// 	// if a < b {
// 	// 	fmt.Println("the value of a is greater than b")
// 	// } else {
// 	// 	fmt.Println("the value of b is greater")
// 	// }

// 	// if a == 10 {

// 	// 	if b == 5 {

// 	// 		fmt.Printf("the values of a and b are %d and %d respectively", a, b)
// 	// 	}
// 	// }

// 	// 	if a == b {

// 	// 		fmt.Println("the value of a is same with the value of b")
// 	// 	} else if a > b {

// 	// 		fmt.Println("the value of a is greater than b")
// 	// 	} else {
// 	// 		fmt.Println("the value of a is less than b")
// 	// 	}

// 	//	Expression SWITCH STATEMENT

// 	// switch {						theres a true (boolean) expression automatically attached to the switch statement for comparison
// 	// case b == 5 && a == 1:
// 	// 	fmt.Println("its a good day")

// 	// case a == 1:
// 	// 	fmt.Println("today is monday")

// 	// case a == 2:
// 	// 	fmt.Println("today is tuesday")

// 	// case a == 3:
// 	// 	fmt.Println("today is wednesday")

// 	// default:
// 	// 	fmt.Println("the value does not match")

// 	// }

// 	var value interface{} = "Go programming" // the variable "value" is an integer converted to a string

// 	switch a := value.(type) {
// 	case int64:
// 		fmt.Println("this is an integer type", a)

// 	case float64:
// 		fmt.Println("this is a float type", a)

// 	case string:
// 		fmt.Println("this is a string type", a)

// 	}

// }

func main() {
	language := "rust"

	switch language {
	case "python":
		fmt.Println("you are learning python")

	case "go", "golang":
		fmt.Println("you are learning go programming language")

	default:
		fmt.Println("any other programming language is okay")

	}

	x := 6

	switch {
	case x%2 == 0:
		fmt.Println("this is an even number")

	case x%2 != 0:
		fmt.Println("this is an odd number")
	}

	hour := time.Now().Hour()
	// fmt.Println(hour)

	switch {
	case hour < 12:
		fmt.Println("good morning")

	case hour < 17:
		fmt.Println("good afternoon")

	case hour < 24:
		fmt.Println("good evening")
	}
}
