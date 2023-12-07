// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
// Declare and initialize an anonymous struct type with the same three fields. Display the value.

package main

import "fmt"

type user struct {
	name  string
	email string
	age   int
}

func main() {
	john := user{
		name:  "John",
		email: "john@gmail.com",
		age:   42,
	}
	fmt.Println("Name: ", john.name)
	fmt.Println("Email: ", john.email)
	fmt.Println("Age: ", john.age)

	// Declare a variable using an anonymous struct
	mrSmith := struct {
		name  string
		email string
		age   int
	}{
		name:  "MrSmith",
		email: "mrSmith@gmail.com",
		age:   50,
	}
	fmt.Println("Name: ", mrSmith.name)
	fmt.Println("Email: ", mrSmith.email)
	fmt.Println("Age: ", mrSmith.age)
}
