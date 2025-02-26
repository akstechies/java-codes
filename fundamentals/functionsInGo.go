package main

import "fmt"

func sayHello() {
	fmt.Println("Hello from Func")
}

func getIncrement(x int, y int) (int, int) {
	return x + 1, y + 1
}

func getSum(x int, y int) int {
	return x + y
}

func nParams(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func getArraySum(arr1 []int) int {
	sum := 0
	for _, arrNum := range arr1 {
		sum += arrNum
	}
	return sum
}

func FunctionsInGo() {
	sayHello()

	sum := getSum(10, 25)
	fmt.Println(sum)

	increX, increY := getIncrement(11, 23)
	fmt.Println(increX, increY)

	nParamsVal := nParams(10, 20, 30, 40, 50)
	fmt.Println(nParamsVal)

	arr1 := []int{1, 2, 3, 4, 5}
	arrSum := getArraySum(arr1)
	fmt.Println(arrSum)
}
