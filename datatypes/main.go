package main

import "fmt"

func main() {
	var a uint8 = 50
	fmt.Println(a)

	var b int16 = 309
	fmt.Println(b)

	c := 75.0
	d := 65.564

	e := c + d

	fmt.Printf("c is %f\n", c)
	fmt.Printf("d is %f\n\n", d)

	fmt.Printf("%f + %f = %f\n", c, d, e)
	fmt.Printf("the value of e = %f\n", e)
	fmt.Printf("the type of e is %T\n\n", e)

	var f = complex(c, d)
	fmt.Println(f)
	fmt.Printf("the data type of f is:  %T", f)
}
