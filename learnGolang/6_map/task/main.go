// Declare and make a map of integer values with a string as the key.
// Populate the map with five values and iterate over the map to display the key/value pairs.

package main

import "fmt"

func main() {
	departments := make(map[string]int)

	departments["IT"] = 42
	departments["Marketing"] = 7
	departments["Sales"] = 22
	departments["Security"] = 5
	departments["Executives"] = 10

	for key, value := range departments {
		fmt.Printf("Dept: %s People: %d\n", key, value)
	}
}
