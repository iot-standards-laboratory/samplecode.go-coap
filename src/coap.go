package main

import (
	"flag"
	"isl/samplecodes/go-coap/src/simpleclient"
	"isl/samplecodes/go-coap/src/simpleserver"
)

func main() {
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
