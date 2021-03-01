package db

import (
	"context"
	"fmt"

	"github.com/axi93/bookstore/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

//InsertRegisterBook where its register the book into the DB
func InsertRegisterBook(b models.Book) (string, error) {
	db := MongoCN.Database("bookStore")
	col := db.Collection("book")

	result, err := col.InsertOne(context.TODO(), b)
	if err != nil {
		return "", err
	}

	//To obtain the ID
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	//Retorn the object as string
	return ObjID.String(), nil
}

//InsertRegisterUser where its register the user into the DB
func InsertRegisterUser(u models.Users) (string, error) {

	db := MongoCN.Database("bookStore")
	col := db.Collection("users")
	u.Password, _ = EncryptPassword(u.Password)
	result, err := col.InsertOne(context.TODO(), u)

	fmt.Println(result)
	fmt.Println(err)

	if err != nil {
		return "", err
	}

	//To obtain the ID
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	//Retorn the object as string
	fmt.Println(ObjID)
	return ObjID.String(), nil
}

//CheckStillBook - checks into DB if the book exists
func CheckStillBook(title string, author string, saga string) (models.Book, bool) {

	db := MongoCN.Database("bookStore")
	col := db.Collection("book")

	//Condition in bson format
	condition := bson.M{"title": title, "author": author, "saga": saga}
	var result models.Book
	//It returns a record. The decode converts the results and I place it as a pointer in result.
	err := col.FindOne(context.Background(), condition).Decode(&result)

	if err != nil {
		return result, false
	}
	return result, true

}

//CheckStillUser - checks into DB if the user exists
func CheckStillUser(email string) (models.Users, bool) {
	db := MongoCN.Database("bookStore")
	col := db.Collection("users")

	//Condition in bson format
	condition := bson.M{"email": email}
	var result models.Users
	//It returns a record. The decode converts the results and I place it as a pointer in result.
	err := col.FindOne(context.Background(), condition).Decode(&result)

	if err != nil {
		return result, false
	}
	return result, true
}

//EncryptPassword it's done for encrypt de password. Price is an algrothim and it's de number of times is encrypted the password. In this case 2*8
func EncryptPassword(pass string) (string, error) {
	price := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), price)

	return string(bytes), err
}
