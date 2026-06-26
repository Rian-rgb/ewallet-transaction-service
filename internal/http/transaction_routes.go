package http

import (
	"ewallet-transaction/infra"

	"github.com/Rian-rgb/ewallet-common-lib/middleware"

	"github.com/gin-gonic/gin"
)

func registerTransactionRoutes(
	api *gin.RouterGroup,
	dependency *infra.Dependency,
	appDeps *infra.AppDependencies,
) {
	transaction := api.Group("/transaction")
	transaction.Use(
		middleware.AuthMiddleware(
			dependency.UmsClient.ValidateToken,
			*appDeps.RedisRepo,
		))

	transaction.POST(
		"/create",
		dependency.TransactionHdl.Create,
	)
}
