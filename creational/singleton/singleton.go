/*
Creational Design Pattern
Used to have only one instance of the class at all times.

We store reference to the object so that we have to create it only once
 */
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

func main() {
	singleton := newSingleton()
	singleton2 := newSingleton()

	singleton.printInt()
	singleton2.printInt()
}
