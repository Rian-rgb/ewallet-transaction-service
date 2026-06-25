package test_test

import (
	"errors"
	"ewallet-transaction/internal/domain/transaction"
	"ewallet-transaction/internal/errs"
	"ewallet-transaction/internal/service"
	"github.com/go-openapi/testify/v2/assert"
	"testing"
)

type MockTransactionRepo struct {
	SaveFunc func(tx *transaction.Entity) error
}

// Save mengimplementasikan interface transaction.IRepository
func (m *MockTransactionRepo) Save(tx *transaction.Entity) error {
	if m.SaveFunc != nil {
		return m.SaveFunc(tx)
	}
	return nil
}

func (r *MockTransactionRepo) FindByReference(reference string) (*transaction.Entity, error) {
	return nil, nil
}

func (r *MockTransactionRepo) UpdateStatus(reference string, status string, additionalInfo string) error {
	return nil
}

func TestCreateTransaction_ValidEntity_Success(t *testing.T) {
	// Arrange
	dummyTx := &transaction.Entity{
		UserID:            1,
		Amount:            20000,
		TransactionType:   transaction.Topup,
		TransactionStatus: transaction.Pending,
		Description:       "test",
	}

	mockRepo := &MockTransactionRepo{
		SaveFunc: func(tx *transaction.Entity) error {
			return nil
		},
	}

	svc := &service.TransactionService{
		TransactionRepo: mockRepo,
	}

	// Act
	result, err := svc.CreateTransaction(dummyTx)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, dummyTx, result)
}

func TestCreateTransaction_EmptyEntity_FailedOnRepository(t *testing.T) {
	// Arrange
	dummyTx := &transaction.Entity{}

	mockRepo := &MockTransactionRepo{
		SaveFunc: func(tx *transaction.Entity) error {
			return errors.New("database connection lost")
		},
	}

	service := &service.TransactionService{
		TransactionRepo: mockRepo,
	}

	expectedErrMsg := errs.New(errs.ErrInternal, "failed to create transaction")

	// Act
	result, err := service.CreateTransaction(dummyTx)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedErrMsg.Error(), err.Error())
}
