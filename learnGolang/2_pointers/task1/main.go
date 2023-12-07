// 1. Declare and initialize a variable of type int with the value of 20.
// 2 .Display the _address of_ and _value of_ the variable.
//
// 3. Declare and initialize a pointer variable of type int that points to the last
// variable you just created.
// 4. Display the _address of_ , _value of_ and the value that the pointer points to_.

package main

import "fmt"

func main() {
	// 1
	value := 20

	// 2. display the address of and value of the variable
	fmt.Println("Adress of: ", &value, "Value of: ", value)

	// 3  Declare and initialize a pointer variable of type int that points to the last
	// variable you just created
	p := &value

	// 4 Display the address of, value of and the value the pointer
	fmt.Println("Adress of: ", &p, "Value of: ", p, "Points to: ", *p)

}
