package simpleserver

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-ocf/go-coap"
	"github.com/go-ocf/go-coap/codes"
)

// Serve : simple example for coap server
func Serve(ip string, port int) {
	fmt.Println("server")
	mux := coap.NewServeMux()
	mux.Handle("/packetCapture", coap.HandlerFunc(handler))
	log.Fatal(coap.ListenAndServe("udp", fmt.Sprintf("%s:%d", ip, port), mux))
}

func handler(w coap.ResponseWriter, req *coap.Request) {
	// log.Printf("Got message in handleB: path=%q: %#v from %v", req.Msg.Path(), req.Msg, req.Client.RemoteAddr())

	fmt.Println("a " + req.Msg.Code().String() + " message received")
	if req.Msg.Code() == codes.PUT || req.Msg.Code() == codes.POST {
		fmt.Println("Payload : " + string(req.Msg.Payload()))
	}
	fmt.Println("")
	resp := w.NewResponse(codes.Content)
	resp.SetOption(coap.ContentFormat, coap.TextPlain)
	resp.SetPayload([]byte("Hello Client!!"))
	ctx, cancel := context.WithTimeout(req.Ctx, time.Second)
	defer cancel()

	if err := w.WriteMsgWithContext(ctx, resp); err != nil {
		log.Printf("Cannot send response: %v", err)
	}
}
