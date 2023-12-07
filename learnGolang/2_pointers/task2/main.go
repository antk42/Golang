// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.

package main

import "fmt"

type user struct {
	name        string
	email       string
	accessLevel int
}

func main() {
	john := user{
		name:        "John",
		email:       "johny@gmail.com",
		accessLevel: 1,
	}

	fmt.Println("access: ", john.accessLevel)

	accessLevel(&john, 10)
	fmt.Println("access: ", john.accessLevel)
}

func accessLevel(u *user, accessLevel int) {
	u.accessLevel = accessLevel
}
