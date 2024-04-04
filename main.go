package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/CRM-Backend/handlers"
	"github.com/CRM-Backend/models"
)

var customers map[int]models.Customer

func initialize() {

	customers = make(map[int]models.Customer)

	file, err := os.ReadFile("database/customers.json")
	if err != nil {
		log.Fatalf("Error read json file: %s", err)
	}

	var data []models.Customer

	err = json.Unmarshal(file, &data)

	if err != nil {
		log.Fatalf("Error decoding json data: %s", err)
	}

	for i, customer := range data {
		customers[i+1] = customer
	}
}

func main() {

	initialize()

	router := mux.NewRouter()
	router.HandleFunc("/", handlers.Index).Methods("GET")
	router.HandleFunc("/customers", handlers.GetAllCustomers(customers)).Methods("GET")
	router.HandleFunc("/customers/{id}", handlers.GetCustomerByID(customers)).Methods("GET")
	router.HandleFunc("/customers", handlers.CreateCustomer(&customers)).Methods("POST")
	router.HandleFunc("/customers/{id}", handlers.UpdateCustomerByID(&customers)).Methods("PUT")
	router.HandleFunc("/customers/{id}", handlers.DeleteCustomerByID(&customers)).Methods("DELETE")

	fmt.Println("Server is starting on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
