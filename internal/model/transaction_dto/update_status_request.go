package transaction_dto

type UpdateStatusRequest struct {
	Reference         string `json:"reference" validate:"required"`
	TransactionStatus string `json:"transactionStatus" validate:"required"`
	AdditionalInfo    string `json:"additionalInfo"`
}
