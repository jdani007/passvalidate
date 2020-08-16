package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Mongo struct gathers connection details
type Mongo struct {
	Database string
	Collection string
	URI string
}

//User is a struct that contains the user credentials
type User struct {
	Firstname string
	Lastname string
	Username string
	Password string
	Email string
}

//InsertCreds function validates the credentials passed by the user.
func (u User)InsertCreds(m Mongo) (interface{}, error) {

	collection := connectMongo(m)

	insertOne, err := collection.InsertOne(context.TODO(),u)
	if err != nil {
		return "", err
	}

	return insertOne, nil
}

//FindCreds retrieves the search criteria from the database
func (u User)FindCreds(search, value string, m Mongo) User {

	collection := connectMongo(m)

	findOne := collection.FindOne(context.TODO(),bson.M{search:value})

	if err := findOne.Decode(&u); err != nil {
		fmt.Println(err)
	}
	
	return u
}

func connectMongo(m Mongo) *mongo.Collection {

	applyURI := options.Client().ApplyURI(m.URI)

	connect, err := mongo.Connect(context.TODO(), applyURI)
	if err != nil {
		fmt.Println(err)
	}

	if err = connect.Ping(context.TODO(), nil); err != nil {
		fmt.Println(err)
	}

	return connect.Database(m.Database).Collection(m.Collection)
}