package domain

import (
	"database/sql"
	"errors"
	"fmt"

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

func NewCustomerRepositoryDB(db *sqlx.DB) CustomerRepositoryDB {
	return CustomerRepositoryDB{
		client: db,
	}
}
