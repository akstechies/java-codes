package main

import "fmt"

func factorial(num int) int {
	fmt.Println("num--", num)
	if num == 0 {
		return 1
	}
	// else
	return num * factorial(num-1)
}

func RecursionInGo() {
	fact := factorial(5)
	fmt.Println(fact)
}
