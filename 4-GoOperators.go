// Go lang - +, -, *, /
// '/' give 2 when we divide 11 by 4

// However if we want to have decimal shit
// we can simple declare the variables in the form of float
//  num1: = 11.0

package main

import "fmt"

func main() {

	num1 := 11.0
	num2 := 4.0

	// Declare both float otherwise it will give error -- float64 / int not possible

	// / divide two floating point variables
	result := num1 / num2
	fmt.Printf(" %g / %g = %g\n", num1, num2, result)

	// Similarly we have modulo operator

	num := 5
	num++
	num--
}

// Operator
// += (addition assignment)
// -= (subtraction assignment)
// *= (multiplication assignment)
// /= (division assignment)
// %= (modulo assignment)
