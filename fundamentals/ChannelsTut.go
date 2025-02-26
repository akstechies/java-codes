package main

import "fmt"

func num1(chan1 chan int) {
	chan1 <- 1 //Send a value into a channel using the channel <- syntax
	chan1 <- 2
	chan1 <- 3
}

func num2(chan2 chan int) {
	for i := 4; i < 10; i++ {
		chan2 <- i
	}
}

func ChannelsTut() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	go num1(channel1)
	go num2(channel2)

	/***
	* using variables will use single value only the first one but in python it is opposite ~ similar to next where in next if we directly call function multiple times it will print only first one so we define in a variable then use
	* But in go it is opposite if we use messages := <-channel1 then print messages then only one the first value will be printed
	 */
	// messages := <-channel1
	fmt.Println(<-channel1)
	fmt.Println(<-channel1)
	fmt.Println(<-channel1)
	// fmt.Println(<-channel1) /** Using extra will give ~ fatal error: all goroutines are asleep - deadlock!*/

	// messages2 := <-channel2
	fmt.Println(<-channel2)
	fmt.Println(<-channel2)
	fmt.Println(<-channel2)
}
