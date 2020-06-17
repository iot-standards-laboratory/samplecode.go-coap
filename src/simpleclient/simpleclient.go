package simpleclient

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
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

	// resp, err := co.GetWithContext(ctx, path)
	// if err != nil {
	// 	log.Fatalf("Error sending request: %v", err)
	// }
	// log.Printf("Response payload: %s", string(resp.Payload()))

	b := make([]byte, 800)

	for i := 0; i < 800; i++ {
		b[i] = 1
	}
	for i := 0; i < 10; i++ {
		msg, err := co.NewPutRequest(path, coap.TextPlain, bytes.NewReader(b))
		if err != nil {
			log.Fatalln(err)
		}
		co.WriteMsgWithContext(ctx, msg)
	}

	// resp, err = co.PutWithContext(ctx, path, coap.TextPlain, strings.NewReader("Hello Server!!"))
	// if err != nil {
	// 	log.Fatalf("Error sending request: %v", err)
	// }
	// log.Printf("Response payload: %s", string(resp.Payload()))

	// resp, err = co.DeleteWithContext(ctx, path)
	// if err != nil {
	// 	log.Fatalf("Error sending request: %v", err)
	// }
	// log.Printf("Response payload: %s", string(resp.Payload()))
}
