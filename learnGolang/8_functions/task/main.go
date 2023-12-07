// Declare a struct type to maintain information about a user. Declare a function
// that creates value of and returns pointers of this type and an error value.
// Call this function from main and display the value.
// Make a second call to your function but this time ignore the value and just test
// the error value.

package main

import "fmt"

type user struct {
	name  string
	email string
}

func newUser() (*user, error) {
	return &user{"Johny", "johny@gmail.com"}, nil
}
func main() {
	u, err := newUser()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*u)

	// Call the function and just check the error on the return.
	_, err = newUser()
	if err != nil {
		fmt.Println(err)
		return
	}
}
