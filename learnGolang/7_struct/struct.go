package main

import "fmt"

func main() {
	type Point struct {
		X int
		Y int
		//Z string
		//T []int
	}

	p := Point{
		X: 5,
		Y: 19,
	}
	fmt.Println(p.X, p.Y)

	p = Point{5, 42}
	fmt.Println(p)

	p.X = 42
	fmt.Println(p)
}
