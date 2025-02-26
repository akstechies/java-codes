package main

import (
	"math/rand"
	"time"
)

func MathInGo() {
	// get random nums
	seedSecs := time.Now().Unix() //This can be used to get unique
	pl(seedSecs)
	randNum := rand.Intn(50)
	pl(randNum)

}
