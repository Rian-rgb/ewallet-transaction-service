package service

import (
	"ewallet-transaction/internal/domain/transaction"
	"ewallet-transaction/internal/errs"
)

type TransactionService struct {
	TransactionRepo transaction.IRepository
}

func (s *TransactionService) CreateTransaction(tx *transaction.Entity) (*transaction.Entity, error) {

	err := s.TransactionRepo.Save(tx)
	if err != nil {
		return nil, errs.New(
			errs.ErrInternal,
			"failed to create transaction",
		)
	}

	return tx, nil
}
