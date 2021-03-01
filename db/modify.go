package db

import (
	"context"

	"github.com/axi93/bookstore/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
//ModifyBook  allows modify the register of the book
func ModifyBook(b models.Book, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("bookStore")
	col := db.Collection("book")

	//make allos create maps or slices
	register := make(map[string]interface{})
	if len(b.Title) > 0 {
		register["title"] = b.Title
	}
	if len(b.Author) > 0 {
		register["author"] = b.Author
	}
	if len(b.Saga) > 0 {
		register["saga"] = b.Saga
	}
	if len(b.PublishCompany) > 0 {
		register["publishcompany"] = b.PublishCompany
	}
	if len(b.Edition) > 0 {
		register["edition"] = b.Edition
	}
	if len(b.Languaje) > 0 {
		register["languaje"] = b.Languaje
	}

	//We perform the setting of the update record
	updtString := bson.M{
		"$set": register,
	}
	//Indicate the ID of the user
	objID, _ := primitive.ObjectIDFromHex(ID)

	//Add filter for upgrade
	filter := bson.M{"_id": bson.M{"$eq": objID}} //eq=equal

	_, err := col.UpdateOne(ctx, filter, updtString)
	if err != nil {
		return false, err
	}
	return true, nil

}
*/
//ModifyUser allows modify the register of the user
func ModifyUser(u models.Users, ID string) (bool, error) {
	bd := MongoCN.Database("bookstore")
	col := bd.Collection("users")

	//make allos create maps or slices
	register := make(map[string]interface{})
	if len(u.Name) > 0 {
		register["name"] = u.Name
	}
	if len(u.Surname) > 0 {
		register["surname"] = u.Surname
	}
	if len(u.Avatar) > 0 {
		register["avatar"] = u.Avatar
	}
	register["dateBirthday"] = u.DateBirthday

	//We perform the setting of the update record
	updtString := bson.M{
		"$set": register,
	}
	//Indicate the ID of the user
	objID, _ := primitive.ObjectIDFromHex(ID)

	//Add filter for upgrade
	filter := bson.M{"_id": bson.M{"$eq": objID}} //eq=equal

	_, err := col.UpdateOne(context.Background(), filter, updtString)
	if err != nil {
		return false, err
	}
	return true, nil
}
