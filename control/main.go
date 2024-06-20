package main

import "fmt"

func main() {

	// SIMPLE FOR LOOP

	// for i := 0; i <= 5; i++ {
	// 	fmt.Printf("i = %d \t go programming \n", i)
	// }

	// infinite loop is when you dont have the conditional statement in a loop

	// for {
	// 	fmt.Println("the language is go")
	// }

	// FOR LOOP AS WHILE LOOP
	// This is when the loop continues to execute as long as the condition remains true
	// Once the condition becomes false, it stops

	// x := 10

	// for x > 2 {
	// 	fmt.Println(x)

	// 	x--
	// }

	// USING BREAK STATEMENT

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)

	// 	if i < 7 {
	// 		break
	// 	}
	// }

	// USING CONTINUE STATEMENT

	// y := 1

	// for y < 9 {

	// 	if y == 5 {
	// 		y = y + 2
	// 		continue
	// 	}
	// 	fmt.Printf("the value of y is %d \n", y)
	// 	y++
	// }

	// GOTO STATEMENT

	x := 0

label1:

	for x < 10 {

		if x == 4 {

			x = x + 2
			goto label1
		}

		fmt.Printf("the value of x is %d \n", x)

		x++
	}

}
