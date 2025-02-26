package main

import "reflect"

func DataTypes() {
	// int, float64, bool, string, rune
	// 0 ,  0.0, 	 false, "" 				~ Default Type

	pl(reflect.TypeOf(22))
	pl(reflect.TypeOf(3.14))

	var typeOf = reflect.TypeOf

	pl(typeOf("yellow"))
	pl(typeOf(true))
}
