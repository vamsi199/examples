package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	pb "github.com/vamsi199/examples/grpc/Hello/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

type gobvar struct {
	Name string
	Id   int
	Date time.Time
}

func main() {
	var gvar gobvar
	b, err := Marshal(&gvar)
	if err != nil {
		log.Println("error marshaling", err)
	}
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("error dailing", err)
	}
	defer conn.Close()

	client := pb.NewHelloClient(conn)
	Bistream, err := client.Duplexstream(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	ch := make(chan byte)

	// receive from server
	go func() {
		for {
			o, err := Bistream.Recv()
			if err == io.EOF {
				close(ch)
				break
			}
			fmt.Println("receiving", o.Response)
		}
	}()

	o := &pb.DuplexOut{b}
	err = Bistream.Send(o)
	if err != nil {
		fmt.Println("error in sending", err)
	}

	err = Bistream.CloseSend()
	if err != nil {
		fmt.Println("error in closing", err)
	}
	<-ch // or else sleep for some time to give the go routine to complete execution

}

func Marshal(v *gobvar) ([]byte, error) {
	*v = gobvar{"vamsi", 12, time.Now()}
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(v)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
