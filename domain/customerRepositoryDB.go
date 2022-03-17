package domain

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct{}

func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {
	var c Customer
	var customers []Customer
	client, err := sql.Open("mysql", "root:eder@/Udemy")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	findAllSql := "select id, name, city, zipcode from customer"
	rows, err := client.Query(findAllSql)

	for rows.Next() {
		rows.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode)
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDB) Find(id int) (Customer, error) {
	var c Customer

	client, err := sql.Open("mysql", "root:eder@/Udemy")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	findAllSql := "select id, name, city, zipcode from customer where id = ?"
	rows, err := client.Query(findAllSql, id)
	rows.Next()
	rows.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode)
	return c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	return CustomerRepositoryDB{}
}
