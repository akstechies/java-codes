package main

import "fmt"

func Conditional() {

	var age int = 18

	if age > 1 && age <= 18 {
		pl("Child")
	} else if age > 18 && age < 50 {
		if age == 25 || age == 30 {
			pl("Married")
		} else {
			pl("Adult")
		}
	} else {
		fmt.Println("Old")
	}

}
