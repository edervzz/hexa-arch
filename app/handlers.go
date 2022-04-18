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

// -----------------------------------
type AccountHandler struct {
	service service.AccountService
}

func (ah *AccountHandler) createAccount(w http.ResponseWriter, r *http.Request) {
	var createAccountRequest service.AccountCreateRequest

	w.Header().Add("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&createAccountRequest)

	if err != nil {
		logger.Info(err.Error())
		w.WriteHeader(http.StatusBadRequest)
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

// -----------------------------------
type PaymItemHandler struct {
	service service.PaymItemService
}

func (ph PaymItemHandler) PostPaymItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var itemReq service.PaymItemPostRequest

	err := json.NewDecoder(r.Body).Decode(&itemReq)
	if err != nil {
		logger.Info(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Cannot create payment item...")
		return
	}
	itemResp, errs := ph.service.Post(itemReq)
	if errs != nil {
		logger.Info(errs.Message)
		w.WriteHeader(errs.Code)
		json.NewEncoder(w).Encode("Cannot create payment item...")
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(itemResp)
}
