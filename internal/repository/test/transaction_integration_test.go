package test_test

import (
	"ewallet-transaction/internal/domain/transaction"
	"ewallet-transaction/internal/repository"
	"ewallet-transaction/test/helper"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSaveTransaction_ValidEntity_ReturnsSuccess(t *testing.T) {
	// Arrange
	testDB := helper.SetupTestTX(t)
	repo := repository.TransactionRepo{DB: testDB}

	inputTrx := &transaction.Entity{
		ID:                1,
		UserID:            2,
		Amount:            150000,
		TransactionType:   "TOPUP",
		TransactionStatus: "PENDING",
		Description:       "TESTING",
		CreatedAt:         time.Now(),
	}

	// Act
	err := repo.Save(inputTrx)

	var savedTrx transaction.Entity
	errFetch := testDB.First(&savedTrx, "id = ?", inputTrx.ID).Error

	// Assert
	assert.NoError(t, err)
	assert.NoError(t, errFetch)
	assert.Equal(t, inputTrx.Amount, savedTrx.Amount)
	assert.Equal(t, inputTrx.TransactionStatus, savedTrx.TransactionStatus)
}

func TestSaveTransaction_DuplicateID_ReturnsFailed(t *testing.T) {
	// Arrange
	testDB := helper.SetupTestTX(t)
	repo := repository.TransactionRepo{DB: testDB}

	inputTrx := &transaction.Entity{
		ID:                1,
		UserID:            2,
		Amount:            150000,
		TransactionType:   "TOPUP",
		TransactionStatus: "PENDING",
		Description:       "TESTING",
		CreatedAt:         time.Now(),
	}

	// Act
	err := repo.Save(inputTrx)

	errDuplicateID := testDB.Create(inputTrx).Error

	// Assert
	assert.NoError(t, err)
	assert.Error(t, errDuplicateID)
}
