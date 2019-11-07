package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var portVar int
var hostnameVar string

func init() {
	flag.IntVar(&portVar, "port", 27500, "mongodb port number to connect to")
	flag.StringVar(&hostnameVar, "hostname", "", "mongodb hostname to connect to (default to os.Hostname())")
}

type listCommands struct {
	Commands bson.M
}

func main() {
	flag.Parse()

	commands, err := runListCommands(getClient(hostnameVar, portVar))
	if err != nil {
		fmt.Printf("Failed to run listCommands against the admin db with the given client. Err: %v\n", err)
	} else {
		fmt.Printf("Successful execution of listCommands returned: %v\n", commands)
	}
}

func getClient(hostname string, port int) *mongo.Client {
	client, err := mongo.NewClient(generateClientOptions(hostname, port))
	if err != nil {
		panic(fmt.Sprintf("Failed to create a mongodb client. Err: %v\n", err))
	}

	if err = client.Connect(context.TODO()); err != nil {
		client.Disconnect(context.TODO())
		panic(fmt.Sprintf("Client failed to connect to mongod. Err: %v\n", err))
	}

	return client
}

func generateClientOptions(hostname string, port int) *options.ClientOptions {
	var err error
	if hostname == "" {
		hostname, err = os.Hostname()
		if err != nil {
			panic(fmt.Sprintf("No hostname was supplied and running os.Hostname failed with %v\n", err))
		}
	}

	fmt.Printf("Proceeding with hostname '%s' and port '%d'\n", hostname, port)

	opts := options.ClientOptions{}
	opts.ApplyURI(fmt.Sprintf("mongodb://%s:%d/?connect=direct", hostname, port))
	return &opts
}

func runListCommands(client *mongo.Client) (listCommands, error) {
	var commands listCommands

	db := client.Database("admin")
	err := db.RunCommand(context.TODO(), bson.D{{"listCommands", 1}}).Decode(&commands)
	if err != nil {
		return listCommands{}, err
	}

	return commands, nil
}
