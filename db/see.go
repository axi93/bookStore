package db

import (
	"context"
	"fmt"

	"github.com/axi93/bookstore/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SeeUser where you see the user recorded into the DB
func SeeUser(ID string) (models.Users, error) {
	db := MongoCN.Database("bookStore")
	col := db.Collection("users")

	var profile models.Users

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(context.TODO(), condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		fmt.Println("Register not find")
		return profile, err
	}
	return profile, nil

}

//SeeBook where you see the book recorded into the DB
func SeeBook(ID string) (models.Book, error) {
	db := MongoCN.Database("bookStore")
	col := db.Collection("book")

	var profile models.Book

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(context.TODO(), condition).Decode(&profile)
	if err != nil {
		fmt.Println("Register not find")
		return profile, err
	}
	return profile, nil

}
