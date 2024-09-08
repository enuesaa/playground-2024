package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	ctx := context.Background()

	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("app").Collection("notes")
	res, err := collection.InsertOne(ctx, bson.D{
		{Key: "name", Value: "aaa"},
		{Key: "value", Value: "hello"},
	})
	if err != nil {
		panic(err)
	}
	id := res.InsertedID
	fmt.Printf("id: %s\n", id)

	var result struct {
		Value string
	}
	filter := bson.D{{Key: "name", Value: "aaa"}}
	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
		panic(err)
	}

	fmt.Printf("value: %s\n", result.Value)
}
