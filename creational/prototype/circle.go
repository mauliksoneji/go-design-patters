package main

import "fmt"

type circle struct{}

func (s circle) Draw() {
	fmt.Println("Drawing circle")
}

func (s circle) clone() shape {
	return circle{} // copy all properties from circle
}
