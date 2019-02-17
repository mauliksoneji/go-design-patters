package main

func main() {
	serv := New("with-default-args")
	serv.printArgs()
	servWithOverridenArgs := New("with-overriden-args", port(9000), timeout(2000))
	servWithOverridenArgs.printArgs()
}
