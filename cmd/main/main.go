package main

import (
	context "context"
	"log"
	"net"

	pb "github.com/vitorsalgado/grpc-go/cmd/main/proto"
	grpc "google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedHelloWorldServiceServer
}

func (s *server) Greetings(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())

	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	pb.RegisterHelloWorldServiceServer(srv, &server{})

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
