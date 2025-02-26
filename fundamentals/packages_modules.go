package main

import (
	"basics/mypackage"
	"fmt"
)

func PackagesModules() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	strArr := mypackage.IntArrToString(arr1)
	fmt.Println(strArr)
}
