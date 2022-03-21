package app

import (
	"endpoints/domain"
	"endpoints/logger"
	"endpoints/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	router := mux.NewRouter()
	router.StrictSlash(true)

	ch := CustomerHandler{
		service.NewCustomerService(domain.NewCustomerRepositoryDB()),
	}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id}", ch.getCustomer).Methods(http.MethodGet)

	router.HandleFunc("/customers1/v2", getAllCustomers1).Methods(http.MethodGet)
	router.HandleFunc("/customers1", greet).Methods(http.MethodGet)
	router.HandleFunc("/customer1/{customer_id}", getCustomer1).Methods(http.MethodGet)

	port := "8000"
	server := "localhost:" + port
	logger.Info("listening on " + server)
	log.Fatal(http.ListenAndServe(server, router))
}
