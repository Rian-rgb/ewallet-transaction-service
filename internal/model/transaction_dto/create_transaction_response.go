ackage transaction_dto

import (
	"ewallet-transaction/internal/domain/transaction"
)

type CreateTransactionResponse struct {
	Reference         string             `json:"reference"`
	TransactionStatus transaction.Status `json:"transactionStatus"`
}
