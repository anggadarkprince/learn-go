package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// var mutex sync.Mutex
// mutex.Lock()
// code... that accesses shared resource
// mutex.Unlock()
func TestRaceCondition(t *testing.T) {
	x := 0
	var mutex sync.Mutex // Mutex to protect shared resource

	for range 1000 {
		go func() { // run 1000 goroutines (each goroutine will run in parallel)
			for range 100 {
				mutex.Lock() // Lock the mutex before accessing shared resource
				x = x + 1 // each go routine can run run in same time
				mutex.Unlock() // Unlock the mutex after accessing shared resource	
			}
		}()
	}

	time.Sleep(5 * time.Second) // Wait for goroutines to finish
	fmt.Println("Final value of x:", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) Deposit(amount int) {
	account.RWMutex.Lock() // Write lock
	defer account.RWMutex.Unlock()
	account.Balance += amount
	// account.RWMutex.Unlock()
}
func (account *BankAccount) Withdraw(amount int) int {
	account.RWMutex.RLock() // Read lock
	defer account.RWMutex.RUnlock()
	account.Balance -= amount
	return account.Balance
	// account.RWMutex.RUnlock()
}

func TestReadWriteMutex(t *testing.T) {
	account := BankAccount{}

	for range 100 {
		go func() {
			for range 100 {
				account.Deposit(1) // Total should be 10000
			}
			go func() {
				account.Withdraw(1)
			}()
		}()
	}

	time.Sleep(5 * time.Second) // Wait for goroutines to finish
	fmt.Println("Final balance:", account.Balance) // Should be 9900
}

type CreditBalance struct {
	sync.Mutex
	Name string
	Balance int
}

func (user *CreditBalance) Lock() {
	user.Mutex.Lock()
}
func (user *CreditBalance) Unlock() {
	user.Mutex.Unlock()
}
func (user *CreditBalance) Deposit(amount int) {
	user.Lock()
	defer user.Unlock()
	user.Balance += amount
}
func (user *CreditBalance) Transfer(to *CreditBalance, amount int) {
	user.Lock()
	defer user.Unlock()
	fmt.Println("Lock user1", user.Name)
	user.Deposit(-amount) // Withdraw from user 1

	time.Sleep(1 * time.Second) // Simulate some work

	to.Lock()
	defer to.Unlock()
	fmt.Println("Lock useer2", user.Name)
	to.Deposit(amount) // Deposit to user 2

	time.Sleep(1 * time.Second) // Simulate some work
}
func TestDeadlock(t *testing.T) {
	angga := CreditBalance{Name: "Angga", Balance: 1000}
	ari := CreditBalance{Name: "Ari", Balance: 1000}

	go func() {
		angga.Transfer(&ari, 100) // angga lock ari
	}()
	go func() {
		ari.Transfer(&angga, 200) // ari lock angga
	}()

	time.Sleep(3 * time.Second) // Wait for goroutines to finish
	// Note: In a real-world scenario, you would use sync.WaitGroup or channels to wait for goroutines.
	fmt.Println("Final balance Angga:", angga.Balance) // no transfer being locked (DEADLOCK, it's waiting for each other)
	fmt.Println("Final balance Ari:", ari.Balance)
}

