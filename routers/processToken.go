package routers

import (
	"errors"

	"github.com/axi93/bookstore/db"
	"github.com/axi93/bookstore/models"
	"github.com/dgrijalva/jwt-go"
)

//Email value of Email used in all the EndPoints
var Email string

//IDUser is the ID returned from the model, it will be used in all the Endpoints
var IDUser string

//ProcessToken process token for extract their values
func ProcessToken(tk string) (*models.Claim, bool, error) {
	myKey := []byte("TheBookStoreOptions")
	claims := &models.Claim{}

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, find := db.CheckStillUser(claims.Email)
		if find == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, find, nil
	}
	if !tkn.Valid {
		return claims, false, errors.New("Token format Incorrect again")
	}
	return claims, false, err
}
