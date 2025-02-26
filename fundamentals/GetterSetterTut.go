package main

import (
	"basics/mypackage"
	"fmt"
)

func GetterSetterTut() {

	date := mypackage.Date{}

	err := date.SetDay(14)
	if err != nil {
		panic(err)
	}
	fmt.Println("Day is", date.Day())

}
