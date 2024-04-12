import "fmt"

func main() {

	// Integer
	var integer1 int
	var integer2 int

	integer1 = 5
	integer2 = 10

	fmt.Println(integer1)
	fmt.Print(integer1)

	// Float
	var salary1 float32
	var salary2 float64

	salary1 = 50000.503882901
	salary2 = 50000.503882901

	fmt.Println(salary1)
	fmt.Println(salary2)

	// String
	var message string
	message = "Welcome to Programiz"

	fmt.Println(message)

	// Boolean
	var boolValue bool
	boolValue = false

	fmt.Println(boolValue)
}

// Data Types		Description													  Examples
// int				Integer numbers.											  7123, 0, -5, 7023
// float			Numbers with decimal points. 								  20.2, 500.123456, -34.23
// complex			Complex numbers. 											  2+4i, -9.5+18.3i
// string			Sequence of characters.										  "Hello World!", "1 is less than 2"
// bool				Either true or false.	 									  true, false
// byte				A byte (8 bits) of non-negative integers.					  2, 115, 97
// rune				Used for characters. Internally used as 32-bit integers.	 'a', '7', '<'
