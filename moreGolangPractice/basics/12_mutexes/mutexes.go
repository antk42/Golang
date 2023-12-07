package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Mutex

	account := Account{
		Balance: 1000,
		Mutex:   &m,
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go account.WithDraw(700, &wg)
	go account.Deposit(500, &wg)
	wg.Wait()

	fmt.Println("Account Balances Updated")
	fmt.Println(account.Balance)
}

type Account struct {
	Balance int
	Mutex   *sync.Mutex
}

func (a *Account) WithDraw(value int, wg *sync.WaitGroup) {
	a.Mutex.Lock()
	a.Balance -= value
	a.Mutex.Unlock()
	wg.Done()
}

func (a *Account) Deposit(value int, wg *sync.WaitGroup) {
	a.Mutex.Lock()
	a.Balance += value
	a.Mutex.Unlock()
	wg.Done()
}
