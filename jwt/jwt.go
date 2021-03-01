package jwt

import (
	"time"

	"github.com/axi93/bookstore/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//GenerateJWT generates the encrypt with JWT
func GenerateJWT(u models.Users) (string, error) {
	myKey := []byte("TheBookStoreOptions")
	//The data part
	paylod := jwt.MapClaims{
		"email":         u.Email,
		"name":          u.Name,
		"surname":       u.Surname,
		"date_birthday": u.DateBirthday,
		"_id":           u.ID.Hex(),
		"exp":           time.Now().Add(time.Hour * 24).Unix(),
	}
	//We choose something to encrypt with
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, paylod)
	//The signature String is passed
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
