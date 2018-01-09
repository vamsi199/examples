package main

import (
	"fmt"
	pb "github.com/vamsi199/examples/grpc/Hello/pb"
	//"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	//"strconv"
	//"golang.org/x/net/html/atom"
	"bytes"
	"encoding/gob"
	"io"
	"time"
)

type gobvar struct {
	Name string
	Id   int
	Date time.Time
}
type content struct {
	s string
}
type val struct {
	s string
}

func main() {
	fmt.Println("listening...")
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("could not listen:",err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, content{})
	err = s.Serve(lis)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (content) Duplexstream(in pb.Hello_DuplexstreamServer) error {
	var g gobvar
	for {
		o, err := in.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println("receiving")
		err = g.Unmarshal(o.Event)
		if err != nil {
			log.Println("error in unmarshal")
			return err
		}
		fmt.Println("Name:", g)
	}

	o := &pb.DuplexOut1{Response: "E2P event has received"} //NOTE: not a requirement to send response
	err := in.Send(o)
	if err != nil {
		fmt.Printf("error sending message %v: %v\n", o.Response, err)
	}
	return nil

}

func (v *gobvar) Unmarshal(b []byte) error {
	fmt.Println("Unmarshal: I am here", v)
	r := bytes.NewReader(b)
	dec := gob.NewDecoder(r)
	return dec.Decode(&v)
}
