package service_test

import (
	"context"
	"ewallet-transaction/internal/domain/transaction"
	"ewallet-transaction/internal/service"
	"github.com/go-openapi/testify/v2/assert"
	"github.com/pkg/errors"
	"testing"
)

type MockTransactionRepo struct {
	SaveFunc func(entity *transaction.Entity) error
}

func (m *MockTransactionRepo) Save(entity *transaction.Entity) error {
	if m.SaveFunc != nil {
		return m.SaveFunc(entity)
	}
	return nil
}

//func (r *MockTransactionRepo) FindByReference(reference string) (*transaction.Entity, error) {
//	return nil, nil
//}
//
//func (r *MockTransactionRepo) UpdateStatus(reference string, status string, additionalInfo string) error {
//	return nil
//}

func TestCreateTransaction_ValidEntity_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()

	dummyEntity := &transaction.Entity{
		UserID:            1,
		Amount:            20000,
		TransactionType:   transaction.Topup,
		TransactionStatus: transaction.Pending,
		Description:       "test",
	}

	mockRepo := &MockTransactionRepo{
		SaveFunc: func(entity *transaction.Entity) error {
			return nil
		},
	}

	svc := &service.TransactionService{
		TransactionRepo: mockRepo,
	}

	// Act
	result, err := svc.CreateTransaction(ctx, dummyEntity)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, dummyEntity, result)
}

func TestCreateTransaction_EmptyEntity_FailedOnRepository(t *testing.T) {
	// Arrange
	ctx := context.Background()

	dummyEntity := &transaction.Entity{}

	mockRepo := &MockTransactionRepo{
		SaveFunc: func(tx *transaction.Entity) error {
			return errors.New("database connection lost")
		},
	}

	svc := &service.TransactionService{
		TransactionRepo: mockRepo,
	}

	// Act
	result, err := svc.CreateTransaction(ctx, dummyEntity)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
}
