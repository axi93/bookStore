package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Book is the MongoDB database user book
type Book struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title          string             `bson:"title" json:"title"`
	Saga           string             `bson:"saga" json:"saga,omitempty"`
	Author         string             `bson:"author" json:"author"`
	PublishCompany string             `bson:"publishcompany" json:"publishCompany,omitempty"`
	Languaje       string             `bson:"languaje" json:"languaje,omitempty"`
	Edition        string             `bson:"edition" json:"edition,omitempty"`
	Genre          string             `bson:"genre" json:"genre,omitempty"`
}
