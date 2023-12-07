// Declare an interface named speaker with a method named speak.
// Declare a struct named english that represents a person who speaks english and declare a struct named
// chinese for someone who speaks chinese. Implement the speaker interface for each
// struct using a value receiver and these literal strings "Hello World" and "你好世界".
// Declare a variable of type speaker and assign the address of a value of type english
// and call the method. Do it again for a value of type chinese.
//
// Add a new function named sayHello that accepts a value of type speaker.
// Implement that function to call the speak method on the interface value. Then create
// new values of each type and use the function.

package main

import "fmt"

type speaker interface {
	speak()
}
type english struct{}

func (english) speak() {
	fmt.Println("Hello world")
}

type chinese struct{}

func (*chinese) speak() {
	fmt.Println("你好世界")
}

func main() {
	var sp speaker
	var eng english
	sp = eng
	sp.speak()

	var c chinese
	sp = &c
	sp.speak()

	sayHello(english{})
	sayHello(&english{})
	sayHello(&chinese{})
}

func sayHello(sp speaker) {
	sp.speak()
}
