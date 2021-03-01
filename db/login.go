package db

import (
	"fmt"

	"github.com/axi93/bookstore/models"
	"golang.org/x/crypto/bcrypt"
)

//TryLogin check login into DB
func TryLogin(email string, password string) (models.Users, bool) {
	usu, find := CheckStillUser(email)
	if find == false {
		fmt.Println("Error en el usuario")
		return usu, false
	}
	passwordBytes := []byte(password)
	passwordDB := []byte(usu.Password)
	//Check password from parametrer and from database. Databse Pass is encrypted so decrypt.
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		fmt.Println("Error en la password")
		return usu, false
	}

	return usu, true

}
