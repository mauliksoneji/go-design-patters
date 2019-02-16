package main

import (
	"time"
	"fmt"
)

func startTimer(tt time.Time, functionName string) func() time.Duration{
	return func() time.Duration {
		d := time.Now().Sub(tt)
		fmt.Println(fmt.Sprintf("%s function took %f seconds", functionName, d.Seconds()))
		return d
	}
}

func main() {
	endTimer := startTimer(time.Now(), "main")
	defer endTimer()
	time.Sleep(time.Second * 2)
}