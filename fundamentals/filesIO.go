package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func FilesIO() {
	// create file
	file, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(file)
	defer file.Close()

	// create prime number and convert to string
	iPrime := []int{2, 3, 5, 7, 9}

	var sPrime []string

	for _, iPrimeNos := range iPrime {
		sPrime = append(sPrime, strconv.Itoa(iPrimeNos))
	}
	fmt.Println(sPrime)

	for _, sPrimeVal := range sPrime {
		_, err := file.WriteString(sPrimeVal + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	// Need this to open and read the file because above we have closed the file
	file, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Read File
	scan1 := bufio.NewScanner(file)
	for scan1.Scan() {
		fmt.Println("Prime: ", scan1.Text())
	}
	if scan1.Err(); err != nil {
		log.Fatal(err)
	}

	fileStat, err := os.Stat("data.txt")
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File doesnot exists")
	} else {
		fmt.Println(fileStat)
		// this one we can used to define what we can do with the file, just .Open() is associated with read only, and here we can also pass permission
		fileOpen, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		checkErr(err)
		fmt.Println(fileOpen)
		defer fileOpen.Close()

		if _, err := fileOpen.WriteString("14\n"); err != nil {
			log.Fatal(err)
		}
	}
}
