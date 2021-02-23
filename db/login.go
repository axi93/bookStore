package db

import "github.com/axi93/bookstore/models"

//TryLogin check login into DB
func TryLogin(email string) (models.Users, bool) {
	usu, find, _ := CheckStillUser(email)
	if find == false {
		return usu, false
	}
	return usu, true

}
