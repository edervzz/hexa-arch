package service

import (
	"endpoints/domain"
	"endpoints/errs"
	"endpoints/logger"
	"time"
)

// Request-Response
type PaymItemPostRequest struct {
	AccountID int     `json:"accountId"`
	TransType string  `json:"transType"`
	Amount    float32 `json:"amount"`
	Concept   string  `json:"concept"`
}

type PaymItemPostResponse struct {
	DocumentId int
}

// Service Interface
type IPaymItemService interface {
	Post(PaymItemPostRequest) (*PaymItemPostResponse, *errs.AppError)
}

// Service Implementation
type PaymItemService struct {
	repo domain.IPaymItemRepository
}

func (p PaymItemService) Post(request PaymItemPostRequest) (*PaymItemPostResponse, *errs.AppError) {
	t := time.Now().Local()
	s := t.Format("2006-01-02 15:04:05")

	paymItem := domain.PaymItem{
		DocumentId: 0,
		AccountId:  request.AccountID,
		TAmount:    request.Amount,
		TransType:  request.TransType,
		Concept:    request.Concept,
		Status:     3,
		DatePost:   s,
		DateValue:  s,
	}

	documentId, errs := p.repo.Post(paymItem)
	if errs != nil {
		logger.Info(errs.Message)
		return nil, errs
	}
	response := PaymItemPostResponse{
		DocumentId: documentId,
	}
	return &response, nil
}

func NewPaymItemPostService(repo domain.IPaymItemRepository) PaymItemService {
	return PaymItemService{
		repo,
	}
}
