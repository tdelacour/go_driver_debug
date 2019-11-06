package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	opts := options.ClientOptions{}
	client, err := mongo.NewClient(&opts)
	fmt.Printf("CLIENT: %+v\n", client)
	fmt.Printf("ERR: %v\n", err)
}
