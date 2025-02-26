/**
Go supports anonymous functions, which can form closures. Anonymous functions are useful when you want to define a function inline without having to name it.
*/

package main

import "fmt"

// suppose we defined 2 functions and passing functions inside functions
func closureAddFunc(f func(int, int) int, x, y int) {
	fmt.Println("Answer: ", f(x, y))
}

func addFunc(a, b int) int {
	return a + b
}

func ClosuresInGo() {
	intSum := func(x, y int) int { return x + y }
	fmt.Println(intSum(10, 15))

	closureAddFunc(addFunc, 10, 30)
}
