package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// "fmt"
// "time"

func ForLoop() {
	// start := time.Now()

	for i := 1; i <= 500; i++ {
		// fmt.Println(i)
	}

	// elapsed := time.Since(start)
	// fmt.Printf("Time taken: %s\n", elapsed)

	// For like While loop
	var1 := 0
	for var1 < 5 {
		// fmt.Println(var1)
		var1++
	}

	// Guessing Number Game

	// 1. Generate a random Number
	// 2. use infinite loop
	// 3. get input from user
	// 4. trim any spaces (optional)
	// 5. Convert it into int ~ bcoz input gets string
	// 6. Check condition if match

	pl("GUESS GAME: ")
	randNum := rand.Intn(50)

	// for true - also works but below is suggested by go
	for {
		pl("Please enter a number: ")

		// INPUT
		reader := bufio.NewReader(os.Stdin)
		num1, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Rand Num is: ", randNum)
		// trim white space
		num1 = strings.TrimSpace(num1)

		// convert to int
		num2, err2 := strconv.Atoi(num1)
		if err2 != nil {
			log.Fatal(err2)
		}

		// check
		if num2 > randNum {
			fmt.Println("Pick Lower")
		} else if num2 < randNum {
			fmt.Println("Pick Higher")
		} else {
			pl("Gotcha")
			break
		}

	}

}
