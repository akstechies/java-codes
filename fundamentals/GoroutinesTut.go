package main

import (
	"fmt"
	"time"
)

// In Goroutines we cannot define in which order it will be executed

func countTo10() {
	for i := 0; i <= 10; i++ {
		fmt.Println("Fun 1---", i)
	}
}

func countTo15() {
	for i := 0; i <= 15; i++ {
		fmt.Println("Fun 2---", i)
	}
}

func countTo20() {
	for i := 0; i <= 20; i++ {
		fmt.Println("Fun 3---", i)
	}
}

func GoroutinesTut() {
	countTo10()

	go countTo15()
	go countTo20()

	time.Sleep(time.Second) //If we dont wait then function will exit bcoz it is running parallel to main will exit without waiting for the output
	// goroutines order are not defined so for that we can use channels
	/**
	Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.
	*/

}
