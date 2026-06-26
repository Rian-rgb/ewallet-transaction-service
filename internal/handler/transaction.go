package handler

import (
	"ewallet-transaction/internal/domain/transaction"
	"ewallet-transaction/internal/dto/transaction_dto"
	"ewallet-transaction/internal/errors"
	"net/http"

	appErrors "github.com/Rian-rgb/ewallet-common-lib/errors"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"github.com/Rian-rgb/ewallet-common-lib/response"
	"github.com/Rian-rgb/ewallet-common-lib/security"
	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	TransactionService transaction.IService
}

// @Summary		Create Transaction
// @Description	Processes a new transaction, updating the user's current balance and generating a transaction record.
// @Tags		Transaction
// @Accept		json
// @Produce		json
//
// @Param		Authorization	header		string										true	"Bearer <token>"
// @Param		request			body		transaction_dto.CreateTransactionRequest	true	"Payload create transaction"
//
// @Success		201	{object}	response.SuccessResponse{data=transaction_dto.CreateTransactionResponse}	"Created"
// @Failure		400	{object}	response.BadRequestResponse													"Bad Request"
// @Failure		401	{object}	response.ErrorResponse														"Unauthorized"
// @Failure		500	{object}	response.ErrorResponse														"Internal Server Error"
//
// @Security	BearerAuth
// @Router		/transaction/create [post]
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
