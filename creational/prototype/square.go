package main

import "fmt"

type square struct{}

func (s square) Draw() {
	fmt.Println("Drawing square")
}

func (s square) clone() shape {
	return square{} // copy all properties from square
}


