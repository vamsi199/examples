package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/kit/endpoint"
	"context"
	"github.com/prometheus/common/log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type HelloRequest struct {
	S string `json:"s"`
}

type HelloResponse struct {
	V string `json:"v"`
	E string `json:"e,omitempty"`
}

type Server struct {}

type Service interface {
	HelloWorld(context.Context, string) (string,error)
}

func (Server) HelloWorld (ctx context.Context,s string) (string,error) {
	s1 := s+"from request"
	return s1,nil
}

func HelloEndpoint(service Service) endpoint.Endpoint {
	return func (ctx context.Context,request interface{}) (interface{},error){
		r:=request.(HelloRequest)
		v,err := service.HelloWorld(ctx,r.S)
		if err != nil {
			log.Fatal("error in getting the HelloWorld service", err)

			return HelloResponse{v, err.Error()}, nil
		}
		return HelloResponse{v,""},nil
	}
}

func main () {
	s := Server{}
	router:= mux.NewRouter()
	handler:=httptransport.NewServer(HelloEndpoint(s),DecodeReq,EncodeResp)
	router.Handle("/msg",handler).Methods("POST")
	log.Fatalln(http.ListenAndServe(":8080",router))
}

func DecodeReq (ctx context.Context,r *http.Request)(interface{},error) {
	var ri HelloRequest
	err:=json.NewDecoder(r.Body).Decode(&ri)
	if err!= nil {
		log.Fatalln("error in decode request",err)
	}
	return ri,err
}

func EncodeResp (ctx context.Context,w http.ResponseWriter,resp interface{}) error {
	err:=json.NewEncoder(w).Encode(resp)
	if err!= nil {
		log.Fatalln("error in encode response",err)
	}
	return nil
}



