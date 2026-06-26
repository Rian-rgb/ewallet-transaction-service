package repository_test

import "gorm.io/gorm"

func createTransactionTypes(db *gorm.DB) error {
	queries := []string{
		`
		CREATE TYPE transaction_type AS ENUM (
			'TOPUP',
			'PURCHASE',
			'REFUND'
		)
		`,
		`
		CREATE TYPE transaction_status AS ENUM (
			'PENDING',
			'SUCCESS',
			'FAILED'
			'REVERSED'
		)
		`,
	}

	for _, query := range queries {
		if err := db.Exec(query).Error; err != nil {
			return err
		}
	}

	return nil
}
