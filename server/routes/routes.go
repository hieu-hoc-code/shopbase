package routes

import (
	"log"
	"net/http"

	"../controllers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Init() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", controllers.Welcome).Methods("GET")

	router.HandleFunc("/api/register", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")

	router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":3000", handler))
}
