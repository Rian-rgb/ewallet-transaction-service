package http

import (
	"ewallet-transaction/infra"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	dependency *infra.Dependency,
	appDeps *infra.AppDependencies,
) {
	api := router.Group("/api/v1")

	registerTransactionRoutes(
		api,
		dependency,
		appDeps,
	)

	registerSwaggerRoutes(router)
}
