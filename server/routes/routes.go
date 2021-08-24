package routes

import (
	"log"
	"fmt"
	"net/http"

	"../controllers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Init() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", controllers.Welcome).Methods("GET")

	// auth
	router.HandleFunc("/api/register", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")

	// product
	router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products", controllers.AllProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", controllers.GetProduct).Methods("GET")
	router.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("DELETE")

	// order
	router.HandleFunc("/api/cartitems", controllers.CreateCartItem).Methods("POST")
	router.HandleFunc("/api/cartitems/{id}", controllers.GetCartItems).Methods("GET")
	router.HandleFunc("/api/cartitems/{id}", controllers.UpdateCartItem).Methods("PUT")
	router.HandleFunc("/api/cartitems/{id}", controllers.DeleteCartItem).Methods("DELETE")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5000"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)
	fmt.Println("server is running http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", handler))
}
