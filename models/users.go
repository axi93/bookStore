package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Users is the MongoDB database user model
type Users struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name         string             `bson:"name" json:"name,omitempty"`
	Surname      string             `bson:"surname" json:"surname,omitempty"`
	DateBirthday time.Time          `bson:"dateBirthday" json:"dateBirthday,omitempty"`
	Email        string             `bson:"email" json:"email"`
	Password     string             `bson:"password" json:"password,omitempty"`
	Avatar       string             `bson:"avatar" json:"avatar,omitempty"`
}
