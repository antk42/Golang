package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(1 * time.Second)
			fmt.Printf("job %d\n", i)
		}(i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("main job1 done?")
	time.Sleep(1 * time.Second)
	fmt.Println("main job2 done?")
}
