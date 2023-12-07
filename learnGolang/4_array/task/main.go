// Declare an array of 5 strings with each element initialized to its zero value.
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
// Display the string value and address of each element.
package main

import "fmt"

func main() {
	var names [5]string

	friends := [5]string{"Adam", "John", "Jane", "Andy", "Thomas"}

	names = friends

	for i, name := range names {
		fmt.Println(name, &names[i])
	}
}
