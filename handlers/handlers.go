package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/CRM-Backend/models"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, "static/index.html")
}

func GetAllCustomers(customers map[int]models.Customer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(customers)
	}
}

func GetCustomerByID(customers map[int]models.Customer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := mux.Vars(r)["id"]

		for _, customer := range customers {
			if customer.ID == id {
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(customer)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
	}
}

func CreateCustomer(customers *map[int]models.Customer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newCustomer models.Customer

		w.Header().Set("Content-Type", "application/json")

		reqBody, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Error reading body request", http.StatusBadRequest)
			return
		}
		json.Unmarshal(reqBody, &newCustomer)

		if newCustomer.ID == "" {
			id := len(*customers)
			id++
			newCustomer.ID = strconv.Itoa(id)
		}

		id, _ := strconv.Atoi(newCustomer.ID)

		if _, ok := (*customers)[id]; ok {
			w.WriteHeader(http.StatusConflict)
		} else {
			(*customers)[id] = newCustomer
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newCustomer)
		}

	}
}

func UpdateCustomerByID(customers *map[int]models.Customer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newCustomer models.Customer

		w.Header().Set("Content-Type", "application/json")

		id := mux.Vars(r)["id"]

		customerID, _ := strconv.Atoi(id)

		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading body request", http.StatusBadRequest)
			return
		}
		json.Unmarshal(reqBody, &newCustomer)

		if id != newCustomer.ID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if _, ok := (*customers)[customerID]; ok {
			(*customers)[customerID] = newCustomer
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newCustomer)
			return
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func DeleteCustomerByID(customers *map[int]models.Customer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := mux.Vars(r)["id"]

		for idx, customer := range *customers {
			if customer.ID == id {
				delete(*customers, idx)
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(customer)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
	}
}
