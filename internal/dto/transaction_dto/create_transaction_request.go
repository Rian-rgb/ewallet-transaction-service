package transaction_dto

import (
	"ewallet-transaction/internal/domain/transaction"
	"github.com/Rian-rgb/ewallet-common-lib/response"
	"github.com/Rian-rgb/ewallet-common-lib/utils"
	"github.com/go-playground/validator/v10"
)

type CreateTransactionRequest struct {
	Amount          float64          `json:"amount" validate:"required,gt=0" example:"10000"`
	TransactionType transaction.Type `json:"transaction_type" validate:"required,oneof=TOPUP PURCHASE REFUND" example:"TOPUP"`
	Description     string           `json:"description" validate:"required" example:"Top up to buy internet package"`
}

func (req CreateTransactionRequest) Validate() []response.ValidationErrorField {
	v := validator.New()
	err := v.Struct(req)

	return response.MapValidationErrors(err)
}

func (req *CreateTransactionRequest) ToEntity(userID int) *transaction.Entity {
	return &transaction.Entity{
		UserID:            userID,
		Amount:            req.Amount,
		Reference:         utils.GenerateUUID(),
		Description:       req.Description,
		TransactionType:   req.TransactionType,
		TransactionStatus: transaction.Pending,
	}
}
