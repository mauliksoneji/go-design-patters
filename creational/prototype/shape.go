package main

type shape interface {
	Draw()
	clone() shape
}

func GetShape(shapeType string) shape {
	if shapeType == "square" {
		return square{}
	}
	return circle{}
}
