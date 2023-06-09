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

// ****************************************************************
// Update
// ****************************************************************

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// MUX extracts the id from the URL and assigns the value to the id variable. 
	productId := mux.Vars(r)["id"]
	// 	Then the code checks if the passed product Id actually exists in the product table.
	if checkIfProductExists(productId) == false {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	// If found, GORM queries the product record to the product variable. 
	// The JSON decoder then converts the request body to a product variable, which is then saved to the database table.
	var product entities.Product
	database.Instance.First(&product, productId)
	json.NewDecoder(r.Body).Decode(&product)
	database.Instance.Save(&product)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// ****************************************************************
// Delete
// ****************************************************************

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//  Extracts the id to be deleted from the request URL. 
	// Checks if the ID is actually available in the product table. 
	productId := mux.Vars(r)["id"]
	if checkIfProductExists(productId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	// create a new product variable.
	var product entities.Product
	// GORM deletes the product by ID.
	database.Instance.Delete(&product, productId)
	// send a message of “Product Deleted Successfully!” to the client
	json.NewEncoder(w).Encode("Product Deleted Successfully!")
}