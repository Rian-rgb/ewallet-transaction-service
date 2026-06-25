package cmd

import (
	"ewallet-transaction/helper"
	"ewallet-transaction/internal/errs"
	"github.com/gin-gonic/gin"
)

func (d *Dependency) MiddlewareRefreshToken(ctx *gin.Context) {
	var (
		log = helper.Logger
	)
	auth := ctx.Request.Header.Get("authorization")
	if auth == "" {
		log.Println("authorization empty")
		helper.SendResponseError(ctx, errs.New(
			errs.ErrUnauthorized,
			"unauthorized",
		))
		ctx.Abort()
		return
	}

	tokenData, err := d.UserClient.ValidateToken(ctx.Request.Context(), auth)
	if err != nil {
		log.Error(err)
		helper.SendResponseError(ctx, errs.New(
			errs.ErrUnauthorized,
			"unauthorized",
		))
		ctx.Abort()
		return
	}

	ctx.Set("external", tokenData)
	ctx.Next()
	return
}
