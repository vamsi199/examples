package main

import (
	"context"
	pb "github.com/vamsi199/examples/grpc/Hello/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type healthClient struct {
	client pb.HealthClient
	conn   *grpc.ClientConn
}

// NewGrpcHealthClient returns a new grpc Client.
func NewGrpcHealthClient(conn *grpc.ClientConn) *healthClient {
	client := new(healthClient)
	client.client = pb.NewHealthClient(conn)
	client.conn = conn
	return client
}

func (c *healthClient) Close() error {
	return c.conn.Close()
}

func check(ctx context.Context, conn *grpc.ClientConn) (bool, error) {
	healthclient := NewGrpcHealthClient(conn)

	var res *pb.HealthCheckResponse
	var err error
	req := new(pb.HealthCheckRequest)

	res, err = healthclient.client.Check(ctx, req)
	if err == nil {
		if res.GetStatus() == pb.HealthCheckResponse_SERVING {
			return true, nil
		}
		return false, nil
	}
	switch grpc.Code(err) {
	case
		codes.Aborted,
		codes.DataLoss,
		codes.DeadlineExceeded,
		codes.Internal,
		codes.Unavailable:
		// non-fatal errors
	default:
		return false, err
	}

	return false, err
}
