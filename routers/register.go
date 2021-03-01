package routers

import (
	"encoding/json"
	"net/http"

	"github.com/axi93/bookstore/db"
	"github.com/axi93/bookstore/models"
)

//RegisterBook is the function to create the book record in the DB.
func RegisterBook(w http.ResponseWriter, r *http.Request) {
	var t models.Book

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Error into the data recived"+err.Error(), 400)
		return
	}
	if len(t.Title) == 0 {
		http.Error(w, "The Title of book is required", 400)
		return
	}
	if len(t.Author) == 0 {
		http.Error(w, "The Author of book is required", 400)
		return
	}
	if len(t.Saga) == 0 {
		http.Error(w, "The Saga of book is required", 400)
		return
	}

	if _, find := db.CheckStillBook(t.Title, t.Author, t.Saga); find == true {
		http.Error(w, "There's exist a book with that Information", 400)
		return
	}
	//Validations - it is checked for proper registration
	status, err := db.InsertRegisterBook(t)
	if err != nil {
		http.Error(w, "An error ocurred at the time of register the error, please try it again"+err.Error(), 400)
		return
	}
	//Validations - check again that it has been registered correctly (problems with mongodb that sometimes does not return true)
	if status == "" {
		http.Error(w, "An error ocurred at the time of register the error, maybe its mongoDB", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

//RegisterUser is the function to create the user record in the DB.
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var u models.Users

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Error into the data recived"+err.Error(), 400)
		return
	}
	if len(u.Email) == 0 {
		http.Error(w, "The email is required", 400)
		return
	}

	if _, find := db.CheckStillUser(u.Email); find == true {
		http.Error(w, "There's exist a user with that email", 400)
		return
	}
	//Validations - it is checked for proper registration
	status, err := db.InsertRegisterUser(u)
	if err != nil {
		http.Error(w, "An error ocurred at the time of register the error, please try it again"+err.Error(), 400)
		return
	}
	//Validations - check again that it has been registered correctly (problems with mongodb that sometimes does not return true)
	if status == "" {
		http.Error(w, "An error ocurred at the time of register the error, maybe its mongoDB", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
