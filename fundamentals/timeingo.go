package main

import (
	"fmt"
	"time"
)

func TimeInGo() {
	timeInGo := time.Now()
	fmt.Println(timeInGo)

	pl(timeInGo.Year(), timeInGo.Day(), timeInGo.Month())
	pl(timeInGo.Hour(), timeInGo.Minute(), timeInGo.Second())
}
