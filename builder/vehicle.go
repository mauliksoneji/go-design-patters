package main

import "fmt"

type sportsCar struct {
	color    color
	maxSpeed speed
}

func (product sportsCar) Drive() {
	fmt.Printf("I am driving a bike of color %s\n", product.color)
}

func (product sportsCar) Stop() {
	fmt.Printf("I am stopping  bikeafter driving at max speed of %f\n", product.maxSpeed)
}

type normalCar struct {
	color    color
	maxSpeed speed
}

func (product normalCar) Drive() {
	fmt.Printf("I am driving a car of color %s\n", product.color)
}

func (product normalCar) Stop() {
	fmt.Printf("I am stopping car after driving at max speed of %f\n", product.maxSpeed)
}
