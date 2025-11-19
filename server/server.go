package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	pb "github.com/Horizon-School-of-Digital-Technologies/library/api"

	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Implementation of LibraryService
type LibraryServer struct {
	pb.UnimplementedLibraryServiceServer
}

// Implement the CreateBook RPC (dummy)
func (s *LibraryServer) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
	// simply echo back the book with an ID
	book := req.GetBook()
	book.Id = 1
	return &pb.CreateBookResponse{Book: book}, nil
}

func main() {
	// Create Prometheus metrics collector for gRPC
	serverMetrics := grpcprom.NewServerMetrics()

	// Create gRPC server with Prometheus interceptors
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(serverMetrics.UnaryServerInterceptor()),
		grpc.ChainStreamInterceptor(serverMetrics.StreamServerInterceptor()),
	)

	// Register the service
	pb.RegisterLibraryServiceServer(grpcServer, &LibraryServer{})
	serverMetrics.InitializeMetrics(grpcServer)

	// Serve Prometheus metrics on :2112
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Println("Prometheus metrics on :2112/metrics")
		log.Fatal(http.ListenAndServe(":2112", nil))
	}()

	// Start gRPC server on :50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("gRPC server on :50051")
	log.Fatal(grpcServer.Serve(lis))
}
