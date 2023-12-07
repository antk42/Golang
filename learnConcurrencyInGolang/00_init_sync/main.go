package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Job %d succes \n", i)
	}
	fmt.Println("Main job success")
}
