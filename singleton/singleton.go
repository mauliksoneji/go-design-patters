package main

import (
	"math/rand"
	"fmt"
)

var mySingeleton *singleton

type singleton struct {
	myInt int
}

func newSingleton() *singleton {
	if mySingeleton == nil {
		mySingeleton = &singleton{
			myInt: rand.Intn(100),
		}
	}
	return mySingeleton
}

func (s singleton) printInt() {
	fmt.Printf("My int is %d\n", s.myInt)
}
