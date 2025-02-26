package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// let us create a mutex to print cubes

// first define
type Number struct {
	lock_mu sync.Mutex
	counter int
}

func (number *Number) cubes(num int) {
	number.lock_mu.Lock()
	defer number.lock_mu.Unlock()
	number.counter++
	fmt.Println("Num is ", math.Pow(float64(num), 3))
}

func MutexInGo() {
	var number Number
	for i := 0; i < 10; i++ {
		go number.cubes(i)
	}

	time.Sleep(time.Second)

}

// ----- BANK ACCOUNT EXAMPLE -----
// Here I'll simulate customers accessing a
// bank account and lock out customers to
// allow for individual access
type Account struct {
	balance int
	lock_mu sync.Mutex // Mutual exclusion
}

func (a *Account) GetBalance() int {
	a.lock_mu.Lock()
	defer a.lock_mu.Unlock()
	return a.balance
}

func (a *Account) Withdraw(v int) {
	a.lock_mu.Lock()
	defer a.lock_mu.Unlock()
	if v > a.balance {
		pl("Not enough money in account")
	} else {
		fmt.Printf("%d withdrawn : Balance : %d\n",
			v, a.balance)
		a.balance -= v
	}
}

func AccountMutex() {
	var acct Account
	acct.balance = 100
	pl("Balance :", acct.GetBalance())

	for i := 0; i < 12; i++ {
		go acct.Withdraw(10)
	}
	time.Sleep(2 * time.Second)

}

// NOTES

/**

How does a Mutex work?
- Lock: Before accessing the shared resource, the goroutine acquires the lock. If another goroutine already holds the lock, the current goroutine waits until the lock is released.
- Unlock: After completing the operation on the shared resource, the goroutine releases the lock so others can proceed.

-> A Mutex (short for "mutual exclusion") in Go is used to prevent multiple goroutines from accessing a shared resource (like a variable, map, or file) at the same time. It ensures only one goroutine can access the critical section of code at any given time, preventing data races.

Why use a Mutex?
- Concurrency: When multiple goroutines try to access or modify shared data, a race condition may occur, leading to unpredictable results.
- Synchronization: Mutex locks ensure that only one goroutine at a time can access a critical section, making data modifications safe.

When to use a Mutex?
- Use a Mutex when multiple goroutines need to:
	-  Write to shared data (e.g., a map or slice).
- Read and write data simultaneously.
	- If only reads are being performed, you can use a Read/Write Mutex (sync.RWMutex) for better performance.

How does a Mutex work?
	- Lock: Before accessing the shared resource, the goroutine acquires the lock. If another goroutine already holds the lock, the current goroutine waits until the lock is released.
	- Unlock: After completing the operation on the shared resource, the goroutine releases the lock so others can proceed.
*/
