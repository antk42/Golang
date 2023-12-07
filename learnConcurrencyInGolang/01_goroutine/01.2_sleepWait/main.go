package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("job %d\n", i)
		}
	}()
	fmt.Println("main job1 done?")
	time.Sleep(time.Second)
	fmt.Println("main job2 done?")

}
