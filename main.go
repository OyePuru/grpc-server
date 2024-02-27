package main

import (
	"context"
	"log"
	"net"

	// Importing the auto-generated gRPC client package from grpcproto
	grpcprotopb "github.com/OyePuru/grpc-proto/gen/go/client/proto/grpcproto"
	"google.golang.org/grpc"
)

type server struct {
	// Embedding the generated interface services
	grpcprotopb.UnimplementedGreeterServer
	grpcprotopb.UnimplementedGreeter2Server
}

// Implementation of the SayHello method for the Greeter service
func (s *server) SayHello(ctx context.Context, in *grpcprotopb.HelloRequest) (*grpcprotopb.HelloReply, error) {
	// Concatenate the received name with " world" and return
	return &grpcprotopb.HelloReply{Message: in.Name + " world"}, nil
}

// Implementation of the SayHello2 method for the Greeter2 service
func (s *server) SayHello2(ctx context.Context, in *grpcprotopb.HelloRequest2) (*grpcprotopb.HelloReply2, error) {
	// Return a predefined message
	return &grpcprotopb.HelloReply2{Message: "GET HELLO WORLD"}, nil
}

// Register the Greeter service with the server
func RegisterGrpcServicesWithServer(grpcServer grpc.ServiceRegistrar) {
	grpcprotopb.RegisterGreeterServer(grpcServer, &server{})
	grpcprotopb.RegisterGreeter2Server(grpcServer, &server{})
}

func main() {
	// Create a listener on TCP port 9000
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	RegisterGrpcServicesWithServer(grpcServer)

	log.Println("Serving gRPC on 0.0.0.0:9000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
