package main

import (
	"context"
	"log"
	"net"

	grpcprotopb "github.com/OyePuru/grpc-proto/gen/go/proto/grpcproto"
	"google.golang.org/grpc"
)

type server struct {
	// Embedding the generated interface services
	grpcprotopb.UnimplementedExampleGetServiceServer
	grpcprotopb.UnimplementedExamplePostServiceServer
}

// Implementation of the ExampleGetHandler method for the ExampleGetService service
func (s *server) ExampleGetHandler(ctx context.Context, in *grpcprotopb.ExampleGetRequest) (*grpcprotopb.ExampleGetResponse, error) {
	return &grpcprotopb.ExampleGetResponse{Message: "Grpc Gateway is working."}, nil
}

// Implementation of the ExamplePostHandler method for the ExamplePostService service
func (s *server) ExamplePostHandler(ctx context.Context, in *grpcprotopb.ExamplePostRequest) (*grpcprotopb.ExamplePostResponse, error) {
	return &grpcprotopb.ExamplePostResponse{Message: "Hello, " + in.Name}, nil
}

// Register the grpc services with the server
func RegisterGrpcServicesWithServer(grpcServer grpc.ServiceRegistrar) {
	grpcprotopb.RegisterExampleGetServiceServer(grpcServer, &server{})
	grpcprotopb.RegisterExamplePostServiceServer(grpcServer, &server{})
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
