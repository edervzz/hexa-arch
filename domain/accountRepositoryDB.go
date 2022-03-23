package domain

import (
	"endpoints/errs"
	"endpoints/logger"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (db AccountRepositoryDB) Save(acct Account) (*Account, *errs.AppError) {
	var a Account = Account{}
	a = acct
	insert := "INSERT INTO account (account_id, customer_id, opening_date, account_type, balance, status) VALUES(?,?,?,?,?,?);"
	sqlresult, err := db.client.Exec(insert, a.AccountID, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Info(err.Error())
		return nil, errs.NewUnexpectedError()
	}

	id, err := sqlresult.LastInsertId()
	if err != nil {
		logger.Info(err.Error())
		return nil, errs.NewUnexpectedError()
	}

	acct.AccountID = strconv.FormatInt(id, 10)
	return &acct, nil
}

func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{
		client: dbClient,
	}
}

// func (db AccountRepositoryDB) Get(id string) (*Account, *errs.AppError) {
// 	selection := "SELECT account_id, customer_id, opening_date, account_type, amount, status FROM account WHERE account_id = ?"
// 	a := Account{}
// 	err := db.client.Select(&a, selection, &id)
// 	if err != nil {
// 		return nil, errs.NewErrorNotFound()
// 	}
// 	return &a, nil
// }

// func (db AccountRepositoryDB) UpdateBalance(id string, balance float64) *errs.AppError {
// 	update := "UPDATE account SET balance = ? WHERE account_id = ?"
// 	result, err := db.client.Exec(update, balance, id)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return errs.NewErrorNotFound()
// 	}

// 	_, err = result.RowsAffected()
// 	if err != nil {
// 		return errs.NewErrorNotFound()
// 	}

// 	return nil
// }
