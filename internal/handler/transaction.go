package handler

import (
	"ewallet-transaction/constan"
	"ewallet-transaction/external/user"
	"ewallet-transaction/helper"
	"ewallet-transaction/internal/domain/transaction"
	"ewallet-transaction/internal/errs"
	"ewallet-transaction/internal/model/transaction_dto"
	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	TransactionService transaction.IService
}

// @Summary Create Wallet Transaction  User
// @Description Processes a new wallet transaction, updating the user's current balance and generating a transaction record.
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string  true  "Insert your access external: Bearer <external>"
// @Param        request  body      transaction_dto.CreateTransactionRequest  true  "Payload create wallet transaction user"
// @Success      200      {object}  helper.SuccessResponse{data=transaction_dto.CreateTransactionResponse}
// @Failure      400      {object}  helper.BadRequestResponse
// @Failure      401      {object}  helper.ErrorResponse
// @Failure      500      {object}  helper.ErrorResponse
// @Router       /create [post]
func (api *TransactionHandler) Create(c *gin.Context) {
	var (
		log = helper.Logger
		req transaction_dto.CreateTransactionRequest
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		//log.Error("failed to parse request: ", err)
		helper.SendResponseBadRequest(c, errs.FromValidator(err))
		return
	}

	token, ok := c.Get("external")
	if !ok || token == nil {
		log.Error("failed to get external data")
		helper.SendResponseError(c, errs.New(
			errs.ErrUnauthorized,
			"unauthorized",
		))
		return
	}

	tokenData, ok := token.(user.Token)
	if !ok {
		log.Error("failed to parse external data")
		helper.SendResponseError(c, errs.New(
			errs.ErrUnauthorized,
			"unauthorized",
		))
		return
	}

	if !transaction.Type.IsValid(req.TransactionType) {
		log.Error("invalid transaction type")
		helper.SendResponseError(c, errs.New(
			errs.ErrBadRequest,
			"invalid transaction type",
		))
		return
	}

	result, err := api.TransactionService.CreateTransaction(req.ToModel(int(tokenData.UserID)))
	if err != nil {
		log.Error("failed to create transaction: ", err)
		helper.SendResponseError(c, err)
		return
	}

	resp := transaction_dto.CreateTransactionResponse{
		Reference:         result.Reference,
		TransactionStatus: result.TransactionStatus,
	}

	helper.SendResponseSuccess(c, "Transaction "+constan.MsgCreated, resp)
}
