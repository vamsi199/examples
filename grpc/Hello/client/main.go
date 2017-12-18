package main

import (
	"fmt"
	pb "github.com/vamsi199/examples/grpc/Hello/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"strconv"
)

func main() {
	//var i pb.Input
	i := &pb.Input{"hello", "vamsi"}
	i1 := &pb.Input1{1, "karthik"+"added"+strconv.Itoa(1)}
	i2 := &pb.Input2{"stream"}
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("error dailing", err)
	}

	client := pb.NewHelloClient(conn)
	resp, err := client.Sayhello(context.Background(), i)
	if err != nil {
		log.Fatalln("cannot get the response", err)
	}
	fmt.Println(resp.Name)

	resp, err = client.SayHelloAll(context.Background(), i1)
	if err != nil {
		log.Fatalln("error dailing", err)
	}
	fmt.Println(resp)

	stream, err := client.StreamAll(context.Background(), i2)
	if err != nil {
		log.Fatalln("cannot get the response stream", err)

	}
	for {
		st, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.streams= _, %v", client, err)
		}
		log.Println(st)
	}

}
