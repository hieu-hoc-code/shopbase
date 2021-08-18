package routes

import (
	"log"
	"net/http"

	"../controllers"
	"github.com/gorilla/mux"
)

func Init() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", controllers.Welcome).Methods("GET")
	
	router.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")

	router.HandleFunc("/api/login", controllers.Login).Methods("POST")

	router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))
}
