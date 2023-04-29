package controllers

import (
	"encoding/json"
	"crud-rest-api-in-golang/database"
	"crud-rest-api-in-golang/entities"
	"net/http"
	"github.com/gorilla/mux"
)

// ****************************************************************
// Create
// ****************************************************************

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Defines a new product variable
	var product entities.Product
	// Decodes the Body of the Incoming JSON request and maps it to the newly created product variable.
	json.NewDecoder(r.Body).Decode(&product)
	// Using GORM, we try to create a new product by passing in the parsed product. 
	// This would create a new record in the products table for us.
	database.Instance.Create(&product)
	// Returns the newly created product data back to the client.
	json.NewEncoder(w).Encode(product)
}
