package main

import "fmt"

type shapeCache struct {
	cache map[string]shape
}

func (s shapeCache) loadCache() {
	s.cache = make(map[string]shape)
	square := GetShape("square")
	circle := GetShape("circle")

	s.cache["square"] = square
	s.cache["circle"] = circle
	fmt.Println(s.cache)
}

func (s shapeCache) createShape(shapeType string) shape {
	fmt.Println(s.cache)
	currShape := s.cache[shapeType]
	return currShape.clone()
}
