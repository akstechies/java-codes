package main

import (
	"bufio"
	"log"
	"os"
)

func GetInput() {

	// Lets take user input
	pl("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n') // means name variable to take the name and err if any error occurs - now here \n means read the input until new line comes or say enter is pressed and enter will give new line - also if I don't want err variable I can do name, _ (underscore)
	if err == nil {                      // got final fragment
		pl("Name is " + name)
	} else {
		log.Fatal(err)
	}

}
