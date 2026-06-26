package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"ewallet-transaction/infra"
	"ewallet-transaction/internal/domain/transaction"
	"ewallet-transaction/internal/dto/transaction_dto"
	"ewallet-transaction/internal/handler"
	"github.com/Rian-rgb/ewallet-common-lib/security"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockTransactionService struct {
	CreateTransactionFunc func(transactionEntity *transaction.Entity) (*transaction.Entity, error)
}

func (m *MockTransactionService) CreateTransaction(ctx context.Context, transactionEntity *transaction.Entity) (*transaction.Entity, error) {
	if m.CreateTransactionFunc != nil {
		return m.CreateTransactionFunc(transactionEntity)
	}
	return nil, nil
}

func TestCreateTransaction_ValidRequestAndToken_ReturnsSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Arrange
	mockService := &MockTransactionService{
		CreateTransactionFunc: func(transactionEntity *transaction.Entity) (*transaction.Entity, error) {
			return &transaction.Entity{
				Reference:         "REF-12345",
				TransactionStatus: "PENDING",
			}, nil
		},
	}

	hdr := &handler.TransactionHandler{
		TransactionService: mockService,
	}

	router := gin.New()

	router.POST("/transaction/create", func(ctx *gin.Context) {
		mockToken := security.Token{UserID: 1}
		security.SetGinToken(ctx, mockToken)
		ctx.Next()
	}, hdr.Create)

	reqBody := transaction_dto.CreateTransactionRequest{
		Amount:          10000,
		TransactionType: "TOPUP",
		Description:     "Testing",
	}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/transaction/create", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)

	var responseBody map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	dataField := responseBody["data"].(map[string]interface{})
	assert.Equal(t, "REF-12345", dataField["reference"])
	assert.Equal(t, "PENDING", dataField["transaction_status"])
}

func TestCreate_InvalidTransactionType_ReturnsBadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Arrange
	infra.InitLogger()

	mockService := &MockTransactionService{}
	hdr := &handler.TransactionHandler{TransactionService: mockService}

	router := gin.New()
	router.POST("/transaction/create", func(ctx *gin.Context) {
		mockToken := security.Token{UserID: 1}
		security.SetGinToken(ctx, mockToken)
		ctx.Next()
	}, hdr.Create)

	reqBody := transaction_dto.CreateTransactionRequest{
		TransactionType: "JENIS_PALSU",
	}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/transaction/create", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreate_MissingToken_ReturnsUnauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Arrange
	infra.InitLogger()

	mockService := &MockTransactionService{}
	hdr := &handler.TransactionHandler{TransactionService: mockService}

	router := gin.New()
	router.POST("/transaction/create", func(ctx *gin.Context) {
		ctx.Next()
	}, hdr.Create)

	reqBody := transaction_dto.CreateTransactionRequest{
		Amount:          10000,
		TransactionType: "TOPUP",
		Description:     "Testing",
	}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/transaction/create", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
