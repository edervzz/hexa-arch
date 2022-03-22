package domain

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {
	// var c Customer
	var customers []Customer = []Customer{}

	findAllSql := "select id, name, city, zipcode from customer"
	err := d.client.Select(&customers, findAllSql)
	if err != nil {
		fmt.Println(err.Error())
		return []Customer{}, err
	}
	// rows, _ := d.client.Query(findAllSql)
	// err := sqlx.StructScan(rows, &customers)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return []Customer{}, err
	// }
	// for rows.Next() {
	// 	rows.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode)
	// 	customers = append(customers, c)
	// }
	return customers, nil
}

func (d CustomerRepositoryDB) Find(id int) (*Customer, error) {
	var c *Customer = &Customer{}

	fmt.Print(d)
	findAllSql := "select id, name, city, zipcode from customer where id = ?"
	rows := d.client.QueryRow(findAllSql, id)
	err := rows.Scan(&c.ID_customer, &c.Name, &c.City, &c.Zipcode)
	if err == sql.ErrNoRows {
		fmt.Println(err.Error())
		return nil, errors.New("Customer not found")
	}
	return c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	client, err := sqlx.Open("mysql", "root:eder@/Udemy")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	err = client.Ping()
	if err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}

	return CustomerRepositoryDB{
		client: client,
	}
}
