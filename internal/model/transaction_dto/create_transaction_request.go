package transaction_dto

import (
	"ewallet-transaction/helper"
	"ewallet-transaction/internal/domain/transaction"
)

type CreateTransactionRequest struct {
	Amount          float64          `json:"amount" binding:"required,gt=0"`
	TransactionType transaction.Type `json:"transaction_type" binding:"required,oneof=TOPUP PURCHASE REFUND"`
	Description     string           `json:"description" binding:"required"`
}

func (req *CreateTransactionRequest) ToModel(userID int) *transaction.Entity {
	return &transaction.Entity{
		UserID:            userID,
		Amount:            req.Amount,
		Reference:         helper.GenerateReference(),
		Description:       req.Description,
		TransactionType:   req.TransactionType,
		TransactionStatus: transaction.Pending,
	}
}
