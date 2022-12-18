package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name string    `bson:"name"` 
	Age int        `bson:"age"`
	Pnum string    `bson:"pnum"`
}

func main() {
	Mongo_URL := "mongodb://127.0.0.1:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(Mongo_URL))
	if err != nil {
		return
	}
	db := client.Database("go-ready")
	col := db.Collection("tPerson")


	strName := "codz"
	filter := bson.D{{"name", strName}}

	estCount, estCountErr := col.EstimatedDocumentCount(context.TODO())
	if estCountErr != nil {
		panic(estCountErr)
	}
	fmt.Println("Total Document COunt", estCount)
	
	count, err := col.CountDocuments(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println("Filter Document COunt", count)


	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	} ()
}