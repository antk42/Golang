package main

import (
	"fmt"
)

func main() {
	a := 42
	b := &a
	fmt.Println(a, b)

	*b = 84 // a = 84
	fmt.Println(a, *b)

	c := &a
	fmt.Println(a, *b, c)

	d := new(int)
	*d = 12
	*c = *d // c = 12 -> a = 12
	fmt.Println(a, *b, *c, *d)
	*d = 24
	fmt.Println(a, *b, *c, *d)

	c = d
	fmt.Println(c, d)
	fmt.Println(*c, *d)

	*c = 100500
	fmt.Println(a, *b, *c, *d)
	fmt.Println(a, b, c, d)
}
