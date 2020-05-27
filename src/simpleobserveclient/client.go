package simpleobserveclient

import (
	"log"

	coap "github.com/go-ocf/go-coap"
)

func observe(w coap.ResponseWriter, req *coap.Request) {
	log.Printf("Got %s", req.Msg.Payload())
}

// Serve : Run Simple Client for Observe
func Serve() {
	sync := make(chan bool)

	co, err := coap.Dial("tcp", "localhost:5683")
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}
	num := 0
	obs, err := co.Observe("/some/path", func(req *coap.Request) {
		log.Printf("Got %s", req.Msg.Payload())
		num++
		if num >= 10 {
			sync <- true
		}
	})
	if err != nil {
		log.Fatalf("Unexpected error '%v'", err)
	}
	<-sync
	obs.Cancel()

}
