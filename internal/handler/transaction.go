package handler

import (
	"ewallet-transaction/internal/domain/transaction"
	"ewallet-transaction/internal/dto/transaction_dto"
	"ewallet-transaction/internal/errors"
	appErrors "github.com/Rian-rgb/ewallet-common-lib/errors"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"github.com/Rian-rgb/ewallet-common-lib/response"
	"github.com/Rian-rgb/ewallet-common-lib/security"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionHandler struct {
	TransactionService transaction.IService
}

// @Summary Create Transaction
// @Description Processes a new transaction, updating the user's current balance and generating a transaction record.
// @Accept       json
// @Produce      json
// @Param        Authorization  header  string  true  "Bearer <token>"
// @Param        request  body      transaction_dto.CreateTransactionRequest  true  "Payload create wallet transaction user"
// @Success      201      {object}  response.SuccessResponse{data=transaction_dto.CreateTransactionResponse}
// @Failure      400      {object}  response.BadRequestResponse
// @Failure      500      {object}  response.ErrorResponse
// @Router       /transaction/create [post]
func (hdl *TransactionHandler) Create(ctx *gin.Context) {
	var (
		req                 transaction_dto.CreateTransactionRequest
		errCodeUnauthorized = appErrors.ErrCodeUnauthorized
		codeBadRequest      = appErrors.ErrCodeBadRequest
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.WithContext(ctx).Error("failed to parse JSON request: ", err)
		response.SendBadRequest(ctx, codeBadRequest, response.InvalidJSONFormatMessage, nil)
		return
	}

	errFields := req.Validate()
	if errFields != nil {
		logger.WithContext(ctx).Warn("request body validation failed")
		response.SendBadRequest(ctx, codeBadRequest, response.InvalidRequestMessage, errFields)
		return
	}

	userData, exists := security.GetGinToken(ctx)
	if !exists {
		logger.WithContext(ctx).Error("token user data no exists: ", userData)
		response.SendError(ctx, errCodeUnauthorized.ToHTTPStatus(), errCodeUnauthorized, response.InvalidTokenMessage)
		return
	}

	transactionEntity := req.ToEntity(userData.UserID)
	result, err := hdl.TransactionService.CreateTransaction(ctx, transactionEntity)
	if err != nil {
		errors.HandleServiceError(ctx, err)
		return
	}

	resp := transaction_dto.CreateTransactionResponse{
		Reference:         result.Reference,
		TransactionStatus: string(result.TransactionStatus),
	}

	response.SendSuccess(ctx, http.StatusCreated, response.SuccessMessage, resp)
}
