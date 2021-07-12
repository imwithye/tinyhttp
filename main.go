package main

import (
	"tinyhttp/args"
	"tinyhttp/log"
	"tinyhttp/net"
)

func main() {
	args.Parse()
	server, err := net.NewServer(*args.Host, *args.Port, *args.Dir)
	if err != nil {
		log.Fatal(err.Error())
	}
	if *args.Open {
		net.OpenBrowser(server.URL())
	}
	err = server.Serve()
	if err != nil {
		log.Fatal(err.Error())
	}
}
