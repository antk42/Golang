package main

import "fmt"

type Sender interface {
	Send(msg string) error
}

type Email struct {
	Address string
}

func (e *Email) Send(msg string) error {
	fmt.Printf("Message \"%v\"was sent via email to %v\n", msg, e.Address)
	return nil
}

func Notify(s Sender) {
	err := s.Send("Notify message")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Success")
}

type Phone struct {
	Number  int
	Balance int
}

func (p *Phone) Send(msg string) error {
	fmt.Printf("Message \"%v\"was sent via phone to %v\n", msg, p.Number)
	return nil
}

func main() {
	email := &Email{"someEmail@gmail.com"}
	Notify(email)

	phone := &Phone{Number: 12345678, Balance: 1000}
	Notify(phone)
}
