package service

import (
	"fraud-analytics/internal/model"
	"fraud-analytics/internal/repository"
)

type TransfersService interface {
	FraudTransfers() []model.Transfer
}

type TransfersServiceImpl struct {
	repo repository.Repository
}

func NewTransfersServiceImpl(repo repository.Repository) *TransfersServiceImpl {
	return &TransfersServiceImpl{
		repo: repo,
	}
}

func (s *TransfersServiceImpl) FraudTransfers() []model.Transfer {
	return s.repo.GetTransfers()
}
