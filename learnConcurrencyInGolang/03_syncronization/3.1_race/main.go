package main

import (
	"fmt"
	"time"
)

type bill struct {
	amount int
}

func (b *bill) subAmout(amount int) {
	if b.amount > 0 {
		time.Sleep(100 * time.Millisecond)
		b.amount -= amount
	}
}

func main() {
	userBill := &bill{amount: 1000}

	go userBill.subAmout(200)
	go userBill.subAmout(400)

	time.Sleep(time.Second)
	fmt.Println(userBill.amount)
}
