package main

import (
	"fmt"
	"strings"
)

func StringsTut() {

	// replace a string
	word1 := "A Word"
	replacerA := strings.NewReplacer("A", "Another")

	replacedWord1 := replacerA.Replace(word1)

	fmt.Println(replacedWord1)

	// string functions
	fmt.Println("length: ", len(word1))
	fmt.Println("contains: ", strings.Contains("Hi Ayush", "Hi"))
	pl("o index", strings.Index(replacedWord1, "o"))
	pl("Replace o: ", strings.Replace("ok Bro", "o", "s", -1))
	pl("Trim space:", strings.TrimSpace("  Heyy boi"))
	pl("Split: ", strings.Split("a-b-c-d-e", "-"))
	pl("Upper: ", strings.ToUpper("Test upp"))
	pl("Lower: ", strings.ToLower("TEST LOW"))

	// --
	fmt.Println("Prefix: ", strings.HasPrefix("Hello World", "Hello"))
	fmt.Println("Suffix: ", strings.HasSuffix("Hello World", "Worlds"))
}
