package main

import (
	"reflect"
	"strconv"
)

func CastingInGo() {

	var1 := 10.11
	var2 := int(var1)
	pl(var1)
	pl(var2)

	// string to int
	var3 := "50000"
	var4, _ := strconv.Atoi(var3) // it returns error also
	pl(var3, reflect.TypeOf(var3))
	pl(var4, reflect.TypeOf(var4))

	// int to string
	var5 := 7744122
	var6 := strconv.Itoa(var5)
	pl(var5, reflect.TypeOf(var5))
	pl(var6, reflect.TypeOf(var6))

	// string to float
	var7 := "3.14"

	if var8, err := strconv.ParseFloat(var7, 64); err == nil {
		pl(var8, reflect.TypeOf(var8))
	}

	var9, _ := strconv.ParseFloat(var7, 64)
	pl(var9, reflect.TypeOf(var9))
}
