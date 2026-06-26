package repository_test

import (
	"ewallet-transaction/internal/domain/transaction"
	"ewallet-transaction/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSaveTransaction_ValidEntity_ReturnsSuccess(t *testing.T) {
	// Arrange
	repo := repository.TransactionRepo{DB: globalTestDB}

	dummyEntity := &transaction.Entity{
		ID:                1,
		UserID:            2,
		Amount:            150000,
		TransactionType:   "TOPUP",
		TransactionStatus: "PENDING",
		Description:       "TESTING",
		CreatedAt:         time.Now(),
	}

	// Act
	err := repo.Save(dummyEntity)

	var savedEntity transaction.Entity
	errFetch := globalTestDB.First(&savedEntity, "id = ?", dummyEntity.ID).Error

	// Assert
	assert.NoError(t, err)
	assert.NoError(t, errFetch)
	assert.Equal(t, dummyEntity.Amount, savedEntity.Amount)
	assert.Equal(t, dummyEntity.TransactionStatus, savedEntity.TransactionStatus)
}

func TestSaveTransaction_DuplicateID_ReturnsFailed(t *testing.T) {
	// Arrange
	repo := repository.TransactionRepo{DB: globalTestDB}

	dummyEntity := &transaction.Entity{
		ID:                1,
		UserID:            2,
		Amount:            150000,
		TransactionType:   "TOPUP",
		TransactionStatus: "PENDING",
		Description:       "TESTING",
		CreatedAt:         time.Now(),
	}

	// Act
	err := repo.Save(dummyEntity)

	errDuplicateID := globalTestDB.Create(dummyEntity).Error

	// Assert
	assert.NoError(t, err)
	assert.Error(t, errDuplicateID)
}
