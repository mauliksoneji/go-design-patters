/*
As creating new objects is expensive, we store a prototype of all object types in cache
when we require to create a new object, we just clone the object in cache and return the cloned object
 */

package main

func main() {
	cache := shapeCache{}
	cache.loadCache()

	square := cache.createShape("square")
	square.Draw()

	circle := cache.createShape("circle")
	circle.Draw()
}
