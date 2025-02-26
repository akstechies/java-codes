package main

import "fmt"

func changeVal(myPtr *int) {
	*myPtr = 15
}

func doubleArr(myArrPtr *[5]int) {
	for i := 0; i < 5; i++ {
		myArrPtr[i] *= 2
	}
}

func PointersinGo() {

	var1 := 53
	fmt.Println("Var1 first:", var1)

	changeVal(&var1) // The &i syntax gives the memory address of variable
	fmt.Println("Var1 Next:", var1)

	f4 := 73
	var f4Ptr *int = &f4
	fmt.Println("address f4:", f4Ptr)
	fmt.Println("value f4:", *f4Ptr)

	// change f4 value using pointer
	*f4Ptr = 55
	fmt.Println("value f4 2:", f4)
	fmt.Println("address f4 2:", &f4)

	// array pointers
	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr1:", arr1)
	doubleArr(&arr1)
	fmt.Println("arr1 doubled:", arr1)
}
