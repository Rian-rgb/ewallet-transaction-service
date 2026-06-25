package helper

import (
	"gorm.io/gorm"
	"testing"
)

func SetupTestTX(t *testing.T) *gorm.DB {
	tx := TestDB.Begin()

	t.Cleanup(func() {
		tx.Rollback()
	})

	return tx
}
