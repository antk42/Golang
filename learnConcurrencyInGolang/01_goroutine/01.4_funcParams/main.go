package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second) // try to comment

		go func() {
			fmt.Printf("Job %d\n", i)
		}()
	}
	time.Sleep(time.Second)
}
