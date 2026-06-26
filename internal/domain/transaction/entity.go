package transaction

import (
	"time"
)

type Entity struct {
	ID                int     `gorm:"primaryKey"`
	UserID            int     `gorm:"column:user_id"`
	Amount            float64 `gorm:"column:amount;type:decimal(15,2)"`
	TransactionType   Type    `gorm:"column:transaction_type;type:transaction_type"`
	TransactionStatus Status  `gorm:"column:transaction_status;type:transaction_status"`
	Reference         string  `gorm:"column:reference;type:varchar(255)"`
	Description       string  `gorm:"column:description;type:varchar(255)"`
	AdditionalInfo    string  `gorm:"column:additional_info;type:text"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (Entity) TableName() string {
	return "transactions"
}
