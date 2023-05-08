package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:1234@localhost:27017/"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("appdb").Collection("users")

	var result struct {
		Name    string   `bson:"name"`
		Age     int      `bson:"age"`
		Hobbies []string `bson:"hobbies"`
		Address struct {
			Street string `bson:"street"`
			City   string `bson:"city"`
		} `bson:"address"`
	}
	// filter := bson.M{"name": "Test name"}
	filter := bson.D{{"name", "New Name"}}

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(result)

}
