package main

import "fmt"

func main() {
	num := 42
	fmt.Printf("Input is %d.\n", num)

	if num < 50 {
		fmt.Println("That is less than 50")
	} else if num == 50 {
		fmt.Println("That is exactly 50")
	} else {
		fmt.Println("That is greater 50")
	}

	if num < 100 {
		fmt.Println("That is also less than 100.")
	}

	if dbl := num * 2; dbl < 100 {
		fmt.Println("Double that number is also less than 100.")
	}
}
