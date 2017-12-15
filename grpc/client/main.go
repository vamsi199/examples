package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/vamsi199/examples/grpc/pb"
	"log"
	"fmt"
	"io"

)

func main () {
	//var i pb.Input
	i := &pb.Input{"hello","vamsi"}
	i1 := &pb.Input1{1,"karthik"}
	i2 := &pb.Input2{"stream"}
	conn,err:=grpc.Dial("localhost:8080",grpc.WithInsecure())
	if err!=nil {
		log.Fatalln("error dailing",err)
	}

	client:=pb.NewHelloClient(conn)
	resp,err:=client.Sayhello(context.Background(),i)
	if err != nil {
		log.Fatalln("cannot get the response",err)
	}
	fmt.Println(resp.Name)

	resp,err=client.SayHelloAll(context.Background(),i1)
	if err!=nil {
		log.Fatalln("error dailing",err)
	}
	fmt.Println(resp)

	stream,err:=client.StreamAll(context.Background(),i2)
	if err != nil {
		log.Fatalln("cannot get the response stream",err)

	}
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		log.Println(feature)
	}

}
