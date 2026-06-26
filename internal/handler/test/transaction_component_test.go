package handler_test

import (
	"bytes"
	"encoding/json"
	"ewallet-transaction/constan"
	"ewallet-transaction/external/user"
	"ewallet-transaction/helper"
	"ewallet-transaction/internal/domain/transaction"
	"ewallet-transaction/internal/dto/transaction_dto"
	"ewallet-transaction/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockTransactionService struct {
	CreateTransactionFunc func(tx *transaction.Entity) (*transaction.Entity, error)
}

func (m *MockTransactionService) CreateTransaction(tx *transaction.Entity) (*transaction.Entity, error) {
	if m.CreateTransactionFunc != nil {
		return m.CreateTransactionFunc(tx)
	}
	return nil, nil
}

func TestCreateTransaction_ValidRequestAndToken_ReturnsSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Arrange
	mockService := &MockTransactionService{
		CreateTransactionFunc: func(tx *transaction.Entity) (*transaction.Entity, error) {
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

	router.POST("/create", func(c *gin.Context) {
		mockToken := user.Token{UserID: 99}
		c.Set("external", mockToken)
		c.Next()
	}, hdr.Create)

	reqBody := transaction_dto.CreateTransactionRequest{
		Amount:          10000,
		TransactionType: "TOPUP",
		Description:     "Testing",
	}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var responseBody map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	assert.Contains(t, responseBody["message"], constan.MsgCreated)

	dataField := responseBody["data"].(map[string]interface{})
	assert.Equal(t, "REF-12345", dataField["reference"])
	assert.Equal(t, "PENDING", dataField["transactionStatus"])
}

func TestCreate_InvalidTransactionType_ReturnsBadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Arrange
	mockService := &MockTransactionService{}
	hdr := &handler.TransactionHandler{TransactionService: mockService}

	router := gin.New()
	router.POST("/create", func(c *gin.Context) {
		c.Set("external", user.Token{UserID: 99})
		c.Next()
	}, hdr.Create)

	reqBody := transaction_dto.CreateTransactionRequest{
		TransactionType: "JENIS_PALSU",
	}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreate_MissingExternalToken_ReturnsUnauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Arrange
	helper.SetupLogger()

	mockService := &MockTransactionService{}
	hdr := &handler.TransactionHandler{TransactionService: mockService}

	router := gin.New()
	router.POST("/create", func(c *gin.Context) {
		c.Set("external", nil)
		c.Next()
	}, hdr.Create)

	reqBody := transaction_dto.CreateTransactionRequest{
		Amount:          10000,
		TransactionType: "TOPUP",
		Description:     "Testing",
	}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
