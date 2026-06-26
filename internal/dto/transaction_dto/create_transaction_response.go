package transaction_dto

type CreateTransactionResponse struct {
	Reference         string `json:"reference"`
	TransactionStatus string `json:"transaction_status"`
}
