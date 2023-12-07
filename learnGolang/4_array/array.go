package main

import "fmt"

func main() {
	var x1 [5]int
	var x2 [0]int
	var x3 [0]string
	fmt.Println(x1, x2, x3)

	var arr [3]int = [3]int{1, 2, 3}
	var arr2 = [3]int{1, 2, 3}
	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr, arr2, arr3)

	arr5 := [...]int{1, 2, 3, 4, 5}
	s := len(arr5)
	fmt.Println(arr5)
	fmt.Println(s)
}
