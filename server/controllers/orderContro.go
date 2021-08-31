package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"../database"
	"../models"
)

type Product struct {
	ProductId string `json:"product_id"`
	Quantity  uint   `json:"quantity"`
	Price     int    `json:"price"`
}

type DataOrder struct {
	PaymentId uint      `json:"payment_id"`
	Products  []Product `json:"products"`
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var data DataOrder
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprintf(w, "error when body parser %v", data)
	}

	total := 0
	for i := 0; i < len(data.Products); i++ {
		total = total + int(data.Products[i].Quantity)*data.Products[i].Price
	}

	order := models.Order{
		Total:      int64(total),
		PaymentId:  data.PaymentId,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}

	database.DB.Create(&order)

	for _, v := range data.Products {
		order_detail := models.OrderDetail{
			ProductId:  v.ProductId,
			Quantity:   v.Quantity,
			OrderId:    order.Id,
			CreatedAt:  time.Now(),
			ModifiedAt: time.Now(),
		}
		database.DB.Create(&order_detail)
	}

	fmt.Fprintf(w, "order thanh cong: %v", order)
}
