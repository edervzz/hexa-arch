package domain

import (
	"endpoints/errs"
	"endpoints/logger"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PaymItem struct {
	DocumentId int     `db:"documentId"`
	AccountId  int     `db:"account_id"`
	TAmount    float32 `db:"tamount"`
	TransType  string  `db:"transType"`
	Status     int     `db:"status"`
	Concept    string  `db:"concept"`
	DatePost   string  `db:"datePost"`
	DateValue  string  `db:"dateValue"`
}

// ----------------------------------------------
type IPaymItemRepository interface {
	Post(PaymItem) (int, *errs.AppError)
}

// ----------------------------------------------
type PaymItemRepositoryDB struct {
	clientdb *sqlx.DB
}

func (db PaymItemRepositoryDB) Post(t PaymItem) (int, *errs.AppError) {
	insert := `INSERT INTO paymitems
	(documentId, account_id, tamount, transType, status, concept, datePost, dateValue)
	VALUES(?,?,?,?,?,?,?,?)`

	sqlRes, err := db.clientdb.Exec(insert,
		t.DocumentId, t.AccountId, t.TAmount, t.TransType, t.Status, t.Concept, t.DatePost, t.DateValue)
	if err != nil {
		logger.Info(err.Error())
		fmt.Println(t)
		return 0, errs.NewBadRequest("Error on save payment Item")
	}
	documentId, err := sqlRes.LastInsertId()
	if err != nil {
		logger.Info(err.Error())
		return 0, errs.NewBadRequest("Error on save payment Item")
	}
	return int(documentId), nil
}

func NewPaymItemRepositoryDB(client *sqlx.DB) PaymItemRepositoryDB {
	return PaymItemRepositoryDB{
		clientdb: client,
	}
}
