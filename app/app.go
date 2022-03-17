package app

import (
	"database/sql"
	"endpoints/domain"
	"endpoints/service"
	"fmt"
	"log"
	"net/http"
	"os"

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

	test, _ := sql.Open("mysql", "root:eder@/Udemy")
	err := test.Ping()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	port := "8000"
	server := "localhost:" + port
	fmt.Println("listening on", port)
	log.Fatal(http.ListenAndServe(server, router))
}
