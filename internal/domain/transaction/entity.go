package transaction

import (
	"time"
)

type Entity struct {
	ID                int       `json:"-"`
	UserID            int       `json:"userId" validate:"required"`
	Amount            float64   `json:"amount" gorm:"column:amount;type:decimal(15,2)" validate:"required"`
	TransactionType   Type      `json:"transactionType" gorm:"column:transaction_type;type:transaction_type" validate:"required,oneof=TOPUP PURCHASE REFUND"`
	TransactionStatus Status    `json:"transactionStatus" gorm:"column:transaction_status;type:transaction_status" validate:"required,oneof=PENDING SUCCESS FAILED REVERSED"`
	Reference         string    `json:"reference" gorm:"column:reference;type:varchar(255)"`
	Description       string    `json:"description" gorm:"column:description;type:varchar(255)" validate:"required"`
	AdditionalInfo    string    `json:"additionalInfo" gorm:"column:additional_info;type:text"`
	CreatedAt         time.Time `json:"-"`
	UpdatedAt         time.Time `json:"-"`
}

func (Entity) TableName() string {
	return "transactions"
}
