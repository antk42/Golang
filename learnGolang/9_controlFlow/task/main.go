// Write a program that inspects a user's name and greets them in a certain way
// if they are on a list or in a different way if they are not.
// Also look at the user's age and tell them some special secret if they are old enough to know it.

package main

import "fmt"

func main() {
	name := "John"
	age := 22

	switch name {
	case "Anna", "Jacob", "Kell", "Carter", "Rory":
		fmt.Println("Hello group number one!")
	case "Seth", "Julia", "Tanner", "Kenton", "Britten":
		fmt.Println("Hello group number two")
	default:
		fmt.Println("Who are you?")
	}
	if age > 30 {
		fmt.Println("Are you already senior?")
	} else {
		fmt.Println("What's up??")
	}

}
