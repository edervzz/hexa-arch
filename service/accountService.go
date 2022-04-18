package service

import (
	"endpoints/domain"
	"endpoints/errs"
	"strconv"
	"time"
)

type AccountCreateRequest struct {
	CustomerId  int     `json:"customerId"`
	AccountType string  `json:"accountType"`
	Amount      float64 `json:"amount"`
}

type AccountCreateResponse struct {
	AccountID string
	Status    int
}

type AccountService struct {
	account domain.IAccountRepository
}

type IAccountService interface {
	CreateAccount(AccountCreateRequest) (*AccountCreateResponse, *errs.AppError)
}

func (a AccountService) CreateAccount(req AccountCreateRequest) (*AccountCreateResponse, *errs.AppError) {
	if req.Amount < 5000 {
		return nil, errs.NewErrorMinimalBalance()
	}
	t := time.Now().Local()
	s := t.Format("2006-01-02")
	newAccount := domain.Account{
		AccountID:   "",
		CustomerId:  strconv.Itoa(req.CustomerId),
		OpeningDate: s,
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      1,
	}

	result, err := a.account.Save(newAccount)
	if err != nil {
		errs.NewUnexpectedError()
		return nil, err
	}

	response := AccountCreateResponse{
		AccountID: result.AccountID,
		Status:    result.Status,
	}

	return &response, nil
}

func NewAccountService(account domain.IAccountRepository) AccountService {
	return AccountService{
		account: account,
	}
}
