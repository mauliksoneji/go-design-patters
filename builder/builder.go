/*
builder pattern is a chain type pattern
Builder is used when creating complex objects and it cannot be built in one step.
So we do builder.setThis().setThat().build()
*/
package main

type color string

const (
	RedColor  color = "RED"
	BlueColor       = "BLUE"
)

type speed float64

const (
	KM   speed = 1
	MILE       = 1.6
)

type builder interface {
	Color(color color) builder
	MaxSpeed(speed speed) builder
	Build() product
}

func (builder RealBuilder) Color(color color) builder {
	builder.color = color
	return builder
}

func (builder RealBuilder) MaxSpeed(speed speed) builder {
	builder.maxSpeed = speed
	return builder
}

func (builder RealBuilder) Build() product {
	if builder.maxSpeed > 100 {
		return sportsCar{
			color:    builder.color,
			maxSpeed: builder.maxSpeed,
		}
	} else {
		return normalCar{
			color:    builder.color,
			maxSpeed: builder.maxSpeed}
	}
}

type product interface {
	Drive()
	Stop()
}

type RealBuilder struct {
	color    color
	maxSpeed speed
}

func main() {
	realBuilder := RealBuilder{}
	sportsCar := realBuilder.Color(RedColor).MaxSpeed(250 * KM).Build()
	sportsCar.Drive()
	sportsCar.Stop()
}
