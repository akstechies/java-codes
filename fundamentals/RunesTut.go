package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
)

func RunesTut() {

	rune1 := "abcdefgh"

	countRune := utf8.RuneCountInString(rune1)
	pl("countRune", countRune)

	rune2 := 'a'
	pl(rune2, reflect.TypeOf(rune2)) // here you can see rune2 is printed in int32 format so we will use formated print , i.e., Printf

	// for loop to print all runes in rune1
	for index, runeVal := range rune1 {
		fmt.Printf("%d : %#U - %c\n", index, runeVal, runeVal)
	}
}
