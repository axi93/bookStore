package routers

/*
//ModifyBook is the function to create the book record in the DB.
func ModifyBook(w http.ResponseWriter, r *http.Request) {
	var t models.Book

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Inccorrect data"+err.Error(), 400)
		return
	}

	var status bool

	status, err = db.ModifyBook(t, ID)
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
