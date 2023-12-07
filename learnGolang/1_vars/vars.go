package main

import "fmt"

func main() {
	var v1 int // v1 == 0
	fmt.Println(v1)

	var v2 int = 100 // v2 == 100
	fmt.Println(v2)

	v3 := 42
	fmt.Println(v3)

	var v5, v6 = 1, 2
	fmt.Println(v5, v6)
	v5, v6 = 3, 4
	fmt.Println(v5, v6)
	v5, v6 = v6, v5
	fmt.Println(v5, v6)

	var v7 int
	v5, v6, v7 = 1, 2, 3
	fmt.Println(v5, v6, v7)

	var EarthRadiusInMeter int = 6371000
	_ = EarthRadiusInMeter

}
