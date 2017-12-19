package lorem

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

type Request struct {
	Min int
	Max int
	Type string
}

type Response struct {
	Error string `json:"error,omitempty"`
	Output string `json:"output"`
}

type Endpoints struct {
	MyEndpoint endpoint.Endpoint
}

func CreateMyEndpiont(ser Service) endpoint.Endpoint {
	return func( ctx context.Context,request interface{}) (interface{},error) {
		req:=request.
	}
}


