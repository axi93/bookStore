package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN -It is the object of connection to the DB
var MongoCN = ConnectionDB()

var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:Lv3s5olo@bookstore.9ehba.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

//ConnectionDB - Makes the connection to the Database
func ConnectionDB() *mongo.Client {

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Connection created to the DB")
	return client
}

//CheckConnection - Ping to DB
func CheckConnection() bool {
	err := MongoCN.Ping(context.Background(), nil)
	if err != nil {
		return false
	}
	return true
}
