package app

import (
	"endpoints/domain"
	"endpoints/logger"
	"endpoints/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Run() {

	a := adder()
	fmt.Println(a(1))
	fmt.Println(a(2))

	b := func(a int) int {
		var x int
		x = x + a
		return x
	}

	fmt.Println(b(1))
	fmt.Println(b(2))

	df := dynamicFunc(3)

	fmt.Println(df())
	fmt.Println(df())
	fmt.Println(df())

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

	ph := PaymItemHandler{
		service: service.NewPaymItemPostService(domain.NewPaymItemRepositoryDB(dbClient)),
	}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id}", ch.getCustomer).Methods(http.MethodGet)

	router.HandleFunc("/accounts", ah.createAccount).Methods(http.MethodPost)

	router.HandleFunc("/paymitems", ph.PostPaymItem).Methods(http.MethodPost)

	server := os.Getenv("SERVER")
	port := os.Getenv("PORT")
	server = server + ":" + port
	logger.Info("listening on " + server)
	log.Fatal(http.ListenAndServe(server, router))
}

func sanityCheck() {
	if os.Getenv("SERVER") == "" ||
		os.Getenv("PORT") == "" ||
		os.Getenv("DB_PASS") == "" ||
		os.Getenv("DB_NAME") == "" {
		log.Fatal("Envars not defined...")
	}
}

func DBConnection() *sqlx.DB {
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	client, err := sqlx.Open("mysql", "root:"+pass+"@/"+name+"")
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

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum = sum + x
		return sum
	}
}

func dynamicFunc(id int) func() string {
	counter := 0
	x := ""
	switch id {
	case 1:
		return func() string {
			counter++
			x = "uno"
			return x + "-" + strconv.Itoa(counter)
		}
	}
	return func() string {
		x = "otro"
		counter++
		return x + "-" + strconv.Itoa(counter)
	}
}
