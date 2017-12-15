package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/vamsi199/examples/grpc/pb"
	"net"
	"log"
	"fmt"
)
type content struct {

	 s string
}
type val struct {
	s string
}

func main () {
	lis,err:=net.Listen("tcp","localhost:8080")
	if err != nil {
		log.Fatal("cannot listen to the post 8080",err)
	}

    server:=grpc.NewServer()
    fmt.Print(lis)
    pb.RegisterHelloServer(server,content{})
	fmt.Print("2")
    server.Serve(lis)
}

func (content) Sayhello (ctx context.Context, i *pb.Input) (*pb.Output,error) {


	fmt.Println(i.Name,i.Wish)
	o:= pb.Output{i.Name}

	return &o,nil



}

func (content) SayHelloAll(ctx context.Context, i *pb.Input1) (o *pb.Output,err error) {
	o=&pb.Output{i.Alias}


return o,nil
}

func (content) StreamAll(i2 *pb.Input2,stream pb.Hello_StreamAllServer)  (err error) {
	s := content{"stream1"}
	s1 := content{"stream2"}
     for i:=0;i<2;i++ {
     	if i==0 {

			err = stream.Send(s)
		}
		err = stream.Send(s1)
		return err
	 }

	 return nil


}
