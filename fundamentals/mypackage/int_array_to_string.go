package mypackage

import (
	"fmt"
	"strconv"
)

func IntArrToString(intArr []int) []string {
	var strArr []string
	for _, i := range intArr {
		strArr = append(strArr, strconv.Itoa(i))
	}
	fmt.Println("My Package Called")
	return strArr
}
