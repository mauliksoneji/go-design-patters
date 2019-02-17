package main

import "fmt"

type server struct {
	httpUrl string
	options *options
}

func (s *server) printArgs() {
	fmt.Println(fmt.Sprintf("[ServerArgs] httpURL: %s, options: %s", s.httpUrl, s.options.String()))
}

func New(httpUrl string, opts ...option) server {
	args := &options{
		port:    8080,
		timeout: 1000,
	}

	for _, opt := range opts {
		opt(args)
	}

	return server{
		options: args,
		httpUrl: httpUrl,
	}
}
