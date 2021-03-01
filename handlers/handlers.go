package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/axi93/bookstore/middlew"
	"github.com/axi93/bookstore/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Handling - When you call the API enter here and define the routes that we will handle
func Handling() {
	router := mux.NewRouter()
	router.HandleFunc("/registerBook", middlew.CheckDB(routers.RegisterBook)).Methods("POST")
	router.HandleFunc("/registerUser", middlew.CheckDB(routers.RegisterUser)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")

	router.HandleFunc("/modifyUser", middlew.CheckDB(middlew.ValidJWT(routers.ModifyUser))).Methods("PUT")
	router.HandleFunc("/seeUser", middlew.CheckDB(middlew.ValidJWT(routers.SeeUser))).Methods("GET")
	router.HandleFunc("/seeBook", middlew.CheckDB(middlew.ValidJWT(routers.SeeBook))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
