package main

import (
	"flag"
	"fmt"
	"isl/samplecodes/go-coap/src/simpleclient"
	"isl/samplecodes/go-coap/src/simpleobserveclient"
	"isl/samplecodes/go-coap/src/simpleobserveserver"
	"isl/samplecodes/go-coap/src/simpleserver"
)

func main() {
	fmt.Println("Hello world!!")

	observeExample()
}

func simpleExample() {
	var server = flag.Bool("server", false, "")
	var ip = flag.String("ip", "127.0.0.1", "")
	var port = flag.Int("port", 5683, "")

	flag.Parse()

	if *server {
		simpleserver.Serve(*ip, *port)
	} else {
		simpleclient.Serve(*ip, *port)
	}
}

func observeExample() {
	var server = flag.Bool("server", false, "")

	flag.Parse()

	if *server {
		simpleobserveserver.Serve()
	} else {
		simpleobserveclient.Serve()
	}
}
