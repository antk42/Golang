package main

import "fmt"

func main() {

	sl := []int{1, 2, 3, 7, 6, 5, 8, 9, 0}
	for key, value := range sl {
		fmt.Printf("%v, %v \n", key, value)
	}

	for _, value := range sl {
		fmt.Print(value)
	}
	fmt.Println()

	ages := map[string]int{
		"Andy":   30,
		"Jane":   22,
		"Thomas": 35,
	}
	for k, v := range ages {
		fmt.Println(k, v)
	}

	//word := "Hello world"
	//for i := 0; i < len(word); i++ {
	//	fmt.Println(word[i])
	//	fmt.Printf("%T", word[i])
	//}

}
