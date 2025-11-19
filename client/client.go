package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "github.com/Horizon-School-of-Digital-Technologies/library/api"
)

func main() {
	// Connect to gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Connection error:", err)
	}
	defer conn.Close()

	// Create client stub
	client := pb.NewLibraryServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Call CreateBook with a simple book
	req := &pb.CreateBookRequest{
		Book: &pb.Book{
			Title:  "1984",
			Author: "George Orwell",
			Isbn:   "9780451524935",
			Genre:  "Dystopian",
			PublicationYear: 1949,
		},
	}

	res, err := client.CreateBook(ctx, req)
	if err != nil {
		log.Fatal("RPC error:", err)
	}

	log.Printf("✅ Book created: ID=%d, Title=%s, Author=%s\n", res.Book.Id, res.Book.Title, res.Book.Author)
}
