// The way fmt.Println() works is similar to how fmt.Print() works with a couple of differences.
// fmt.Println() prints a new line at the end by default.
// If we print multiple values and variables at once, a space is added between the values by default.

// Data Type	Format Specifier
// integer		%d
// float		%g
// string		%s
// bool			%t

package main

import "fmt"

// the Scan() function takes input from the user. However, it can only take input values up to a space.

func main() {
	var name string
	var age int
	var sunny bool

	// takes input value for name
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Printf("Name: %s", name)

	// Taking multiple inputs from the user
	// take name and age input
	fmt.Println("Enter your name and age:")
	fmt.Scan(&name, &age)

	// take boolean input
	fmt.Println("Is the day sunny?")
	fmt.Scanf("%t", &sunny)

}

// The fmt.Scanf() function is similar to Scanln() function. The only difference is that Scanf() takes inputs using format specifiers. For example,
