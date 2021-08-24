package controllers

import (
	"encoding/json"
	"net/http"
	"fmt"
	"time"
	"../models"
	"../database"
	"strconv"
	"github.com/gorilla/mux"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var data models.Order

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprintf(w, "error when body parse %v", data)
		return
	}

	var order models.Order
	database.DB.Where("user_id = ?",data.UserId).Where("product_id = ?", data.ProductId).First(&order)
	
	fmt.Println(order)
	if order.Id != 0 {
		order.Quantity = order.Quantity + data.Quantity
		database.DB.Updates(&order)
		fmt.Fprintf(w, "updated order of user %v", data.UserId)
		return 
	}

	data.CreatedAt = time.Now()
	data.ModifiedAt = time.Now()

	database.DB.Create(&data)
	fmt.Fprintf(w, "created order of user %v", data.UserId)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	var orders []models.Order

	database.DB.Find(&orders)
	json.NewEncoder(w).Encode(orders)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var data models.Order

	database.DB.Where("id = ?", id).First(&data)
	if data.Id == 0 {
		fmt.Fprintf(w, "not found product_id: %v", id)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprintf(w, "err when parse body")
		return 
	}

	data.ModifiedAt = time.Now()

	database.DB.Model(&data).Updates(data)

	fmt.Fprintf(w, "updated order_id: %v", data.Id)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)

	data := models.Order{
		Id: uint(id),
	}

	database.DB.Delete(&data)

	fmt.Fprintf(w, "deleted order_id %v", data.Id)
}