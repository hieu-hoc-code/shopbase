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

func CreateCartItem(w http.ResponseWriter, r *http.Request) {
	var data models.CartItem

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprintf(w, "error when body parse %v", data)
		return
	}

	var cartitem models.CartItem
	database.DB.Where("user_id = ?",data.UserId).Where("product_id = ?", data.ProductId).First(&cartitem)
	
	fmt.Println(cartitem)
	if cartitem.Id != 0 {
		cartitem.Quantity = cartitem.Quantity + data.Quantity
		database.DB.Updates(&cartitem)
		fmt.Fprintf(w, "updated cartitem of user %v", data.UserId)
		return 
	}

	data.CreatedAt = time.Now()
	data.ModifiedAt = time.Now()

	database.DB.Create(&data)
	fmt.Fprintf(w, "created cartitem of user %v", data.UserId)
}

func GetCartItems(w http.ResponseWriter, r *http.Request) {
	var cartitems []models.CartItem

	database.DB.Find(&cartitems)
	json.NewEncoder(w).Encode(cartitems)
}

func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var data models.CartItem

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

	fmt.Fprintf(w, "updated cartitem_id: %v", data.Id)
}

func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)

	data := models.CartItem{
		Id: uint(id),
	}

	database.DB.Delete(&data)

	fmt.Fprintf(w, "deleted cartitem_id %v", data.Id)
}