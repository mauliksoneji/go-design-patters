/*
Creational Design Pattern
Here we want to hide the creation details from the user,
so we have a build function which takes in type of object and builds the object

Factory is just a function which takes in the type and other details to create the object
 */
package main

import "fmt"

type vehicleType string

const (
	vehicleType_Car vehicleType = "CAR"
	vehicleType_Bike vehicleType = "BIKE"
)

type Vehicle interface {
	Drive()
}

type car struct {}

func (c car) Drive() {
	fmt.Println("Car is driving")
}

type bike struct {}

func (b bike) Drive() {
	fmt.Println("Bike is driving")
}

func GetVehicle(vType vehicleType) Vehicle {
	if vType == vehicleType_Car {
		return car{}
	} else {
		return bike{}
	}
}
