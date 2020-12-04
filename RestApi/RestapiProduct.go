package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Item-Product
type Product struct {
	ID    string  `json:"ID"`
	Name  string  `json:"Name"`
	Desc  string  `json:"Desc"`
	Price float64 `json:"Price"`
}

var items []Product

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Fprintf(w, "Endpoints called : HomePage()")
}

//Get items value
func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Function called: getitems()")
	json.NewEncoder(w).Encode(items)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/join")
	fmt.Println("Function called: createProduct()")
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	items = append(items, product)
	json.NewEncoder(w).Encode(product)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/join")
	fmt.Println("Function called: deleteProduct()")
	params := mux.Vars(r)
	_deleteProductAtUid(params["id"])
	json.NewEncoder(w).Encode(items)
}

func _deleteProductAtUid(id string) {
	for index, product := range items {
		if product.ID == id {
			// delete product from slice
			items = append(items[:index], items[index+1:]...)
			break
		}
	}
}

//Update the product details
func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/join")
	fmt.Println("Function called: updateProduct()")
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	params := mux.Vars(r)
	// delete the item at UID
	_deleteProductAtUid(params["id"])
	// create it with new data
	items = append(items, product)
	json.NewEncoder(w).Encode(items)
}

func main() {
	items = append(items, Product{
		ID:    "1",
		Name:  "laptop",
		Desc:  "Lenovo",
		Price: 60000,
	})
	items = append(items, Product{
		ID:    "2",
		Name:  "mobile",
		Desc:  "Dell",
		Price: 95000,
	})
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", homePage).Methods("GET")
	r.HandleFunc("/items", getItems).Methods("GET")
	r.HandleFunc("/items", createProduct).Methods("POST")
	r.HandleFunc("/items/{id}", deleteProduct).Methods("DELETE")
	r.HandleFunc("/items/{id}", updateProduct).Methods("PUT")
	log.Fatal(http.ListenAndServe(":10000", r))
}
