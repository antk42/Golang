package main

import "fmt"

func main() {
	var list []int

	l := len(list)
	c := cap(list)
	fmt.Println(l, c)

	list = []int{1, 2, 3, 4, 5}
	l = len(list)
	c = cap(list)
	fmt.Println(l, c)

	list = make([]int, 0, 5)
	l = len(list)
	c = cap(list)
	fmt.Println(l, c)

	list = make([]int, 5) // 5, 5
	l = len(list)         // 5
	c = cap(list)         // 5
	fmt.Println(l, c)

	list = []int{1, 2, 3, 4}
	list = append(list, 5)
	fmt.Println(list)

	list = make([]int, 0, 3)
	l = len(list)
	c = cap(list)
	fmt.Println(l, c, list)

	list = append(list, 1)
	l = len(list)
	c = cap(list)
	fmt.Println(l, c, list)
	list = append(list, 2) // [1,2] len:2 cap:3
	list = append(list, 3) // [1,2,3] len:3 cap:3
	list = append(list, 4) // [1,2,3,4] len:4 cap:6
	l = len(list)
	c = cap(list)
	fmt.Println(l, c, list)

	fmt.Println("============================")

	list2 := list[1:3] // list [1,2,3,4] [1:3](2,3)-> list2 [2,3]
	fmt.Println(list2)
	list2[0] = 7
	fmt.Println(list2)
	fmt.Println(list)
}
