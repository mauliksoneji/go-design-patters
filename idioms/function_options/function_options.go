package main

import "fmt"

type options struct {
	port    int
	timeout int
}

func (opts *options) String() string {
	return fmt.Sprintf("port: %d, timeout: %d", opts.port, opts.timeout)
}

type option func(*options)

func port(port int) option {
	return func(args *options) {
		args.port = port
	}
}

func timeout(timeout int) option {
	return func(args *options) {
		args.timeout = timeout
	}
}
