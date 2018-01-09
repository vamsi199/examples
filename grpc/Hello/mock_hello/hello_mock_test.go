package mock_pb

import (
	"fmt"
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	pb "github.com/vamsi199/examples/grpc/Hello/pb"
)

var msg = &pb.DuplexOut{[]byte("test")}
var resp=&pb.DuplexOut1{Response:"received"}

func TestHello(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock for the stream returned by RouteChat
	stream:=NewMockHello_DuplexstreamClient(ctrl)

	// set expectation on sending.
	stream.EXPECT().Send(
		gomock.Any(),
	).Return(nil)
	// Set expectation on receiving.
	stream.EXPECT().Recv().Return(msg, nil)
	stream.EXPECT().CloseSend().Return(nil)
	// Create mock for the client interface.
	rgclient := NewMockHelloClient(ctrl)
	// Set expectation on RouteChat
	rgclient.EXPECT().Duplexstream(
		gomock.Any(),
	).Return(stream, resp)
	if err := testHello(rgclient); err != nil {
		t.Fatalf("Test failed: %v", err)
	}

}

func testHello(client pb.HelloClient) error{
	stream,err:= client.Duplexstream(context.Background())
	if err != nil {
		return err
	}
	if err := stream.Send(msg); err != nil {
		return err
	}
	if err := stream.CloseSend(); err != nil {
		return err
	}
	got, err := stream.Recv()
	if err != nil {
		return err
	}
	if !proto.Equal(got, msg) {
		return fmt.Errorf("stream.Recv() = %v, want %v", got, msg)
	}
	return nil

}

