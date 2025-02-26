package main

import (
	"fmt"
	"os"
)

func main() {

	withArgs := os.Args
	withOutPath := os.Args[1:]

	elem := withOutPath[2]

	fmt.Println(withArgs)
	fmt.Println(withOutPath)
	fmt.Println(elem)
}
