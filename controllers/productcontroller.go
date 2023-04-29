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

// ****************************************************************
// Get By ID
// ****************************************************************

// The below function takes in the productId and queries against the database if the ID is found in the table. If there are no records found for the ID, GORM would return the ID as 0 for which the entire function, in turn, would return false. If the Product Id is found in the table, a true flag will be returned. 
func checkIfProductExists(productId string) bool {
	var product entities.Product
	database.Instance.First(&product, productId)
	if product.ID == 0 {
		return false
	}
	return true
}
func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]//Gets the Product Id from the Query string of the request.
	if checkIfProductExists(productId) == false {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product entities.Product //Creates a new product variable.
	database.Instance.First(&product, productId)//With the help of GORM, the product table is queried with the product Id. This would fill in the product details to the newly created product variable.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)//encode the product variable and send it back to the client.
}

// ****************************************************************
// Get All
// ****************************************************************

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []entities.Product//define an empty new list of products.
	database.Instance.Find(&products)//Maps all the available products into the product list variable.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)//write an HTTP Status code of 200 OK to the Header of the response.
	json.NewEncoder(w).Encode(products)	// encodes the products variable and returns it back to the client.
}

