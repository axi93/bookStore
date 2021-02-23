package db

import (
	"context"
	"time"

	"github.com/axi93/bookstore/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertRegisterBook where its register the book into the DB
func InsertRegisterBook(b models.Book) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("bookStore")
	col := db.Collection("book")

	result, err := col.InsertOne(ctx, b)
	if err != nil {
		return "", false, err
	}

	//To obtain the ID
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	//Retorn the object as string
	return ObjID.String(), true, nil
}

//CheckStillBook - checks into DB if the book exists
func CheckStillBook(title string) (models.Book, bool, string) {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("bookStore")
	col := db.Collection("books")

	//Condition in bson format
	condition := bson.M{"title": title}
	var result models.Book
	//It returns a record. The decode converts the results and I place it as a pointer in result.
	err := col.FindOne(ctx, condition).Decode(&result)
	//In the third parameter, we get the objectID and here it is converted to string
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID

}

//InsertRegisterUser where its register the user into the DB
func InsertRegisterUser(u models.Users) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("bookStore")
	col := db.Collection("users")

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	//To obtain the ID
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	//Retorn the object as string
	return ObjID.String(), true, nil
}

//CheckStillUser - checks into DB if the user exists
func CheckStillUser(email string) (models.Users, bool, string) {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("bookStore")
	col := db.Collection("users")

	//Condition in bson format
	condition := bson.M{"email": email}
	var result models.Users
	//It returns a record. The decode converts the results and I place it as a pointer in result.
	err := col.FindOne(ctx, condition).Decode(&result)
	//In the third parameter, we get the objectID and here it is converted to string
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID

}
