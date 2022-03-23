package domain

import "endpoints/errs"

type Account struct {
	AccountID   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      int
}

type IAccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}

// Get(string) (*Account, errs.AppError)
// UpdateBalance(string, float64) (*Account, errs.AppError)
