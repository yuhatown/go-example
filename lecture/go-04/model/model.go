package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	client     *mongo.Client
	colPersons *mongo.Collection
}

type Person struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
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

func (p *Model) GetNamePerson(n string) Person {

	filter := bson.D{{"name", n}}

	var result Person
	err := p.colPersons.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the name %s\n", n)
	} else if err != nil {
		panic(err)
	}

	fmt.Println(result)
	if jsonData, err := json.MarshalIndent(result, "", "    "); err != nil {
		panic(err)
	} else {
		fmt.Printf("%s\n", jsonData)
	}

	return result
}

func (p *Model) GetPnumPerson(pn string) Person {

	filter := bson.D{{"pnum", pn}}

	var result Person
	err := p.colPersons.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the pnum %s\n", pn)
	} else if err != nil {
		panic(err)
	}

	fmt.Println(result)
	if jsonData, err := json.MarshalIndent(result, "", "    "); err != nil {
		panic(err)
	} else {
		fmt.Printf("%s\n", jsonData)
	}

	return result
}

func (p *Model) PostNewperson(na, pn string, ag int) *mongo.InsertOneResult {

	filter := bson.D{{"name", na}, {"age", ag}, {"pnum", pn}}

	result, err := p.colPersons.InsertOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	return result
}

func (p *Model) DeletePnump(pn string) *mongo.DeleteResult {
	filter := bson.D{{"pnum", pn}}

	result, err := p.colPersons.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	return result
}


func (p *Model) PutPnumAge(pn string, ag int) *mongo.UpdateResult {
	filter := bson.D{{"pnum", pn}}
	update := bson.D{{"$set", bson.D{{"age", ag}}}}

	result, err := p.colPersons.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	return result
}
