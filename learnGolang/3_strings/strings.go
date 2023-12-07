package main

import "fmt"

func main() {
	var myString string // ""
	fmt.Println(myString)

	var hello string = "\tHello\n"
	var title string = `\tHello\n`
	fmt.Println(hello, title)

	word := "my string"
	fmt.Println(word)

	str := "你好"
	fmt.Println(str)
	str = "Привет"
	fmt.Println(str)

	var b byte = 'c' // равно 99
	var r rune = '你' // равно 20320
	fmt.Println(b, r)
}
