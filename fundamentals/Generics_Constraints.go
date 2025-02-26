package main

import "fmt"

type MyConstraint interface {
	int | float64
}

func getSumGenerics[T MyConstraint](x T, y T) T {
	return x + y
}

func GenericsConstraints() {
	fmt.Println(getSumGenerics(6.6, 44))
}
