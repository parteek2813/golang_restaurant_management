package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func DBInstance() *mongo.Client {
	MongoDb := "mongodb+srv://digitalparteek:35TAP4RZVPtzmaO9@cluster1.jpqpjku.mongodb.net/?retryWrites=true&w=majority"

	fmt.Println(MongoDb)


	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))

	if err != nil {
		log.Fatal(err)
	}


	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongodb")
	return client
}

var Client  *mongo.Client = DBInstance()

func  OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {


	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)


	return collection
}
