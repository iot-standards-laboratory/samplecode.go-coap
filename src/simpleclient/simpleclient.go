package simpleclient

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-ocf/go-coap"
)

// Serve : simple example for coap client
func Serve(ip string, port int) {
	fmt.Println("client")
	co, err := coap.Dial("udp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}
	path := "/packetCapture"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := co.GetWithContext(ctx, path)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	log.Printf("Response payload: %s", string(resp.Payload()))

	resp, err = co.PostWithContext(ctx, path, coap.TextPlain, strings.NewReader("Hello Server!!"))
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	log.Printf("Response payload: %s", string(resp.Payload()))

	resp, err = co.PutWithContext(ctx, path, coap.TextPlain, strings.NewReader("Hello Server!!"))
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	log.Printf("Response payload: %s", string(resp.Payload()))

	resp, err = co.DeleteWithContext(ctx, path)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	log.Printf("Response payload: %s", string(resp.Payload()))
}
