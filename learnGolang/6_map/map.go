package main

import "fmt"

func main() {
	var m1 map[int]bool
	var m2 map[string]string
	fmt.Println(m1, m2)

	m3 := make(map[int]int)
	fmt.Println(m3)

	ages := map[string]int{
		"Andy": 30,
		"Jane": 25,
	}
	age := ages["Andy"] // 30
	fmt.Println(age)
	ages["Ivan"] = 21
	fmt.Println(ages)
	fmt.Println(ages["Michael"])
	ages["Michael"] = ages["Michael"] + 1
	fmt.Println(ages)

}
