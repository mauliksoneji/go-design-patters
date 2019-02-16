/*
Abstract Factory is a factory of factories.
It is just a function which takes a few parameters and creates the proper factory.
Then it uses the created factory to produce the final object
 */
package main

import "fmt"

type shape interface {
	Draw()
}

func GetShape(shapeType string, isCircular bool) shape {
	if isCircular {
		cShape := circularShape{}
		return cShape.GetShape(shapeType)
	} else {
		nShape := normalShape{}
		return nShape.GetShape(shapeType)
	}
}

type normalShape struct {}

func (shape normalShape) GetShape(shapeType string) shape {
	if shapeType == "square" {
		return square{}
	} else {
		return square{} //default shape
	}
}

type circularShape struct {}

func (shape circularShape) GetShape(shapeType string) shape {
	if shapeType == "circle" {
		return circle{}
	} else {
		return circle{} //default shape
	}
}

type square struct {}

func (s square) Draw() {
	fmt.Println("Drawing square")
}

type circle struct {}

func (s circle) Draw() {
	fmt.Println("Drawing circle")
}

func main() {
	circle := GetShape("circle", true)
	circle.Draw()

	square := GetShape("square", false)
	square.Draw()
}
