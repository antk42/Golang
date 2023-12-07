// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and bool.
// Display the values of those variables.
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).

package main

import "fmt"

func main() {
	// declare vars
	var age int
	var name string
	var legal bool

	// Display the values of those variables
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(legal)

	// short declaretion and initialization
	month := 10
	dayOfWeek := "Tuesday"
	happy := true

	// display them
	fmt.Println(month)
	fmt.Println(dayOfWeek)
	fmt.Println(happy)

	// Perform a type conversion.
	pi := float32(3.14)

	// Display the value of that variable.
	fmt.Printf("%T [%v]\n", pi, pi)
}
