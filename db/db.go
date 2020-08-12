package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const database = "user"
const collection = "users"
const mongoURI = "mongodb://127.0.0.1:27017"


//User is a struct that contains the user credentials
type User struct {
	Firstname string
	Lastname string
	Username string
	Password string
	Email string
}


//InsertCreds function validates the credentials passed by the user.
func (user User)InsertCreds() error {

	collection := connectMongo()

	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return  nil
}


//FindCreds retrieves the username and password from the database
func (user User)FindCreds() (User, error) {

	collection := connectMongo()

	var result User

	err := collection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&result)
	if err != nil {
		return User{}, err
	}
	
	return result, nil
}


func connectMongo() *mongo.Collection {

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println(err)
	}

	return client.Database(database).Collection(collection)
}