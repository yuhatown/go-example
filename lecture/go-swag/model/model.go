package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type Model struct {
	client *mongo.Client
	colPersons *mongo.Collection
}

type Person struct {
	Name string `bson:"name"`
	Age int `bson:"age"`
	Pnum string `bson:"pnum"`
}

func NewModel() (*Model, error) {
	r := &Model{}

	var err error
	mgUrl := "mongodb://127.0.0.1:27017"
	if r.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mgUrl)); err != nil {
		return nil, err
	} else if err := r.client.Ping(context.Background(), nil); err != nil {
		return nil, err
	} else {
		db := r.client.Database("go-ready")
		r.colPersons = db.Collection("tPerson")
	}
	
	return r, nil
}

func (p *Model) PostNewperson(name, pnum string, age int) *mongo.InsertOneResult {
	filter := bson.D{{"name", name}, {"pnum", pnum}, {"age", age}}

	result, err := p.colPersons.InsertOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	return result
}

func (p *Model) PutNameAge(name string, age int) *mongo.UpdateResult {
	filter := bson.D{{"name", name}}
	update := bson.D{{"$set", bson.D{{"age", age}}}}

	result, err := p.colPersons.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	return result
}