package main

import (
	"context"
	"log"
	"net"

	helloworldpb "github.com/OyePuru/grpc-proto/gen/go/client/proto/helloworld"
	"google.golang.org/grpc"
)

type server struct {
	helloworldpb.UnimplementedGreeterServer
	helloworldpb.UnimplementedGreeter2Server
}

func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	return &helloworldpb.HelloReply{Message: in.Name + " world"}, nil
}

func (s *server) SayHello2(ctx context.Context, in *helloworldpb.HelloRequest2) (*helloworldpb.HelloReply2, error) {
	return &helloworldpb.HelloReply2{Message: "GET HELLO WORLD"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	grpcServer := grpc.NewServer()
	helloworldpb.RegisterGreeterServer(grpcServer, &server{})
	helloworldpb.RegisterGreeter2Server(grpcServer, &server{})

	log.Println("Serving gRPC on 0.0.0.0:9000")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
