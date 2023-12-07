package main

import (
	"fmt"
)

func main() {

	name := "Tyrion"

	switch name {
	case "Cersei", "Jaime", "Tyrion":
		fmt.Println("A Lannister always pays their debts.")
	case "Robb", "Sansa", "Bran", "Arya", "Rickon":
		fmt.Println("Winter is coming.")
	default:
		fmt.Println("Who are you?")
	}

	// A switch with no expression is equivalent to "switch true".
	age := 42
	switch {
	case age < 18:
		fmt.Println("Kids rate is $5.00")
	case age >= 18 && age < 55:
		fmt.Println("Adults rate is $7.00")
	case age >= 55:
		fmt.Println("Seniors rate is $6.00")
	}
}
