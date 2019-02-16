package main

func main() {
	singleton := newSingleton()
	singleton2 := newSingleton()

	singleton.printInt()
	singleton2.printInt()
}
