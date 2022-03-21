package app

import (
	"encoding/json"
	"encoding/xml"
	errs "endpoints/err"
	"endpoints/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.service.GetAllCustomers()

	if err != nil {
		notFound := errs.NewErrorNotFound()
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(notFound.Code)
		json.NewEncoder(w).Encode(notFound.AsMessage())
		return
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func (ch *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	customer, err := ch.service.GetCustomer(id)

	if err != nil {
		notFound := errs.NewErrorNotFound()
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(notFound.Code)
		json.NewEncoder(w).Encode(notFound.AsMessage())
		return
	}

	if r.Header.Get("Content-Type") == "application/xml" {

		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customer)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}

type Customer struct {
	Name    string `json:"fullname" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("hola mundo")
}

func getAllCustomers1(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Eder", City: "EDOMEX", Zipcode: "52928"},
		{Name: "Sheila", City: "EDOMEX", Zipcode: "55027"},
		{Name: "Osmar", City: "EDOMEX", Zipcode: "11450"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func getCustomer1(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Eder", City: "EDOMEX", Zipcode: "52928"},
		{Name: "Sheila", City: "EDOMEX", Zipcode: "55027"},
		{Name: "Osmar", City: "EDOMEX", Zipcode: "11450"},
	}
	vars := mux.Vars(r)

	var customer Customer = Customer{}

	for i := range customers {
		if customers[i].Name == vars["customer_id"] {
			customer = customers[i]
		}
	}

	fmt.Println(r.Header.Get("Content-Type"))

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customer)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}
