package main

import "fmt"

func main() {

	// explicitly declare the data type
	var number1 int = 10
	fmt.Println(number1)

	// assign a value without declaring the data type
	var number2 = 20
	fmt.Println(number2)

	// shorthand notation to define variable
	number3 := 30
	fmt.Println(number3)

	// Creating multiple variables at once
	var name, age = "Arjun", 22
	fmt.Println(name, age)

	const lightspeed = 222882
	// lightspeed = 3434343 Error we cant change the value of const

}
