package service

import (
	"context"
	"ewallet-transaction/internal/domain/transaction"
	"ewallet-transaction/internal/errors"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
)

type TransactionService struct {
	TransactionRepo transaction.IRepository
}

func (svc *TransactionService) CreateTransaction(
	ctx context.Context,
	transactionEntity *transaction.Entity,
) (*transaction.Entity, error) {

	err := svc.TransactionRepo.Save(transactionEntity)
	if err != nil {
		logger.WithContext(ctx).Error("failed to save transaction: ", err)
		return nil, errors.ErrInternalServerError
	}

	return transactionEntity, nil
}
