package main

import (
	"fmt"
	"learnGolang/12_scope/example/users"
)

func main() {
	u := users.User{
		Name: "John",
		ID:   10,
		//password: "somePassword",
		// User field 'password' in struct literal

	}
	fmt.Printf("User: %#v\n", u)
}
