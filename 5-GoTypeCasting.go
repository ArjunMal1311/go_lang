package main

import "fmt"

func main() {

	// initialize integer variable to a floating-point number
	var number int = 4

	fmt.Printf("Number is %g", number)

	var number1 int = 20
	var number2 float32 = 5.7
	var sum float32

	// addition of different data types
	sum = float32(number1) + number2

	fmt.Printf("Sum is %g", sum)
}
