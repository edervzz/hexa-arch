package app

import (
	"endpoints/domain"
	"endpoints/logger"
	"endpoints/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Run() {

	sanityCheck()

	router := mux.NewRouter()
	router.StrictSlash(true)

	dbClient := DBConnection()

	ch := CustomerHandler{
		service.NewCustomerService(domain.NewCustomerRepositoryDB(dbClient),
			domain.NewAccountRepositoryDB(dbClient)),
	}

	ah := AccountHandler{
		service: service.NewAccountService(domain.NewAccountRepositoryDB(dbClient)),
	}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id}", ch.getCustomer).Methods(http.MethodGet)

	router.HandleFunc("/accounts", ah.createAccount).Methods(http.MethodPost)

	server := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	server = server + ":" + port
	logger.Info("listening on " + server)
	log.Fatal(http.ListenAndServe(server, router))
}

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" ||
		os.Getenv("DB_USER") == "" ||
		os.Getenv("DB_NAME") == "" {
		log.Fatal("Envars not defined...")
	}
}

func DBConnection() *sqlx.DB {
	user := os.Getenv("DB_USER")
	name := os.Getenv("DB_NAME")
	client, err := sqlx.Open("mysql", "root:"+user+"@/"+name+"")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	err = client.Ping()
	if err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}
	return client
}
