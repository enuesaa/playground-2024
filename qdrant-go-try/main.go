package main

import (
	"context"
	"log"

	pb "github.com/qdrant/go-client/qdrant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:6334", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCollectionsClient(conn)

	r, err := client.List(context.Background(), &pb.ListCollectionsRequest{})
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	log.Printf("%s", r.GetCollections())
}
