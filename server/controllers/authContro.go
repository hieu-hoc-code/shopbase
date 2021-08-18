package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../database"
	"../models"
)

type LoginData struct {
	Email    string
	Password string
}

func Login(w http.ResponseWriter, r *http.Request) {
	var data LoginData

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprintf(w, "err when body parse %v", data)
		return
	}

	var user models.User

	database.DB.Where("email = ?", data.Email).First(&user)

	if user.Email == "" {
		fmt.Fprintf(w, "email not found!")
		return
	}

	if user.Password != data.Password {
		fmt.Fprintf(w, "password incorrect!")
		return
	}

	fmt.Fprintf(w, "login success")
}
