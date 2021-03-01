package routers

import (
	"encoding/json"

	"fmt"
	"net/http"

	"github.com/axi93/bookstore/db"
	"github.com/axi93/bookstore/models"
)

/*
is the function to create the book record in the DB.
func ModifyBook(w http.ResponseWriter, r *http.Request) {
	var t models.Book

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Inccorrect data"+err.Error(), 400)
		return
	}

	var status bool

	status, err = db.ModifyBook(t, IDBook)
	if err != nil {
		http.Error(w, "An error ocurred when we try to modify the register. Try again"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Imposible modify the register of the user", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
*/
//ModifyUser is the function to create the user record in the DB.
func ModifyUser(w http.ResponseWriter, r *http.Request) {
	var u models.Users

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Inccorrect data"+err.Error(), 400)
		return
	}

	var status bool

	status, err = db.ModifyUser(u, IDUser)
	if err != nil {
		http.Error(w, "An error ocurred when we try to modify the register. Try again"+err.Error(), 400)
		return
	}

	fmt.Println(status)
	if status == false {
		http.Error(w, "Imposible modify the register of the user", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
