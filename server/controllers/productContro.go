package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../database"
	"../models"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var data models.Product

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprintf(w, "error when body parse %v", data)
		return
	}

	id := data.Id

	var product models.Product
	database.DB.Where("id = ?", id).First(&product)

	if product.Id != "" {
		fmt.Fprintf(w, "product_id exists")
		return
	}

	database.DB.Create(&data)
	fmt.Fprintf(w, "created product %v", data)
}
