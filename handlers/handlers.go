package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/axi93/bookstore/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Handling - When you call the API enter here and define the routes that we will handle
func Handling() {
	router := mux.NewRouter()
	router.HandleFunc("/registerBook", routers.RegisterBook).Methods("POST")
	//router.HandleFunc("/modifyBook", routers.ModifyBook).Methods("PUT")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
