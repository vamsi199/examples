# example based on justforfunc #12: a Text to Speech server with gRPC and Kubernetes
# https://www.youtube.com/watch?v=XaMr--wAuSI


# instructions:
#
# generate the pb.go:
# cd $GOPATH/src/github.com/muly/howto/golang/web/grpc/echo-example/pb
# protoc -I . echo.proto --go_out=plugins=grpc:.
#
# run the server:
# cd $GOPATH/src/github.com/muly/howto/golang/web/grpc/echo-example/server
# go run main.go
#
# run the client:
# cd $GOPATH/src/github.com/muly/howto/golang/web/grpc/echo-example/client
# go run main.go

