package main

import "fmt"

func ArraysInGO() {

	arr1 := []int{1, 2, 3, 4}

	fmt.Println(arr1)

	for _, arrVal := range arr1 {
		pl(arrVal)
	}

	pl("1st elem: ", arr1[0])
	fmt.Println("len is: ", len(arr1))
	// print elem using for i loop
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	// multi dim array
	arr2 := [3][2]int{
		{12, 44},
		{72, 11},
		{55, 6},
	} // see we have defined 3 but if we give 2 elems then it will give zero in that remaining space
	fmt.Println(arr2)

	fmt.Println(arr2[2][1])

}
