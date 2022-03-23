package app

import (
	"encoding/json"
	"encoding/xml"
	"endpoints/errs"
	"endpoints/logger"
	"endpoints/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	service service.ICustomerService
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

type AccountHandler struct {
	service service.AccountService
}

func (ah *AccountHandler) createAccount(w http.ResponseWriter, r *http.Request) {
	var createAccountRequest service.AccountCreateRequest

	w.Header().Add("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&createAccountRequest)

	if err != nil {
		logger.Info(err.Error())
		w.WriteHeader(500)
		json.NewEncoder(w).Encode("Error during decoding...")
		return
	}

	result, errApp := ah.service.CreateAccount(createAccountRequest)
	if errApp != nil {
		w.WriteHeader(errApp.Code)
		json.NewEncoder(w).Encode(errApp.Message)
		return
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(result)
}
