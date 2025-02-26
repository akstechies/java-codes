package main

import "fmt"

func SlicesInGo() {
	/**
	Slices are similar to arrays, but are more powerful and flexible.

	Like arrays, slices are also used to store multiple values of the same type in a single variable.

	However, unlike arrays, the length of a slice can grow and shrink as you see fit.
	*/

	/**
	1. Using the []datatype{values} format
	2. Create a slice from an array
	3. Using the make() function
	*/

	slice1 := []int{1, 2, 4, 5}
	fmt.Println(slice1)

	arr1 := [5]int{12, 62, 45, 15, 73}
	slice2 := arr1[:3]
	fmt.Println("slice2", slice2)

	// slice_name := make([]type, length, capacity)
	// Note: If the capacity parameter is not defined, it will be equal to length.
	// cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)
	slice3 := make([]int, 6)
	fmt.Println(slice3)

	slice4 := arr1[2:]
	fmt.Println(slice4)

	slice1[3] = 55
	fmt.Println(slice1)

	// apeend 2 slices
	slice5 := append(slice1, slice2...) //Note: The '...' after slice2 is necessary when appending the elements of one slice to another.
	fmt.Println("slice5", slice5)

	// append in a slice
	slice5 = append(slice5, 7)
	fmt.Println("slice5", slice5)
}
