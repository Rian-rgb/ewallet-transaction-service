package cmd

import (
	_ "ewallet-transaction/docs"
	"ewallet-transaction/helper"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
)

// @title           E-Wallet API - Mobile/Web
// @version         0.0
// @description     API Service untuk m_transaction dompet digital pengguna (Transaction).
// @description     Fitur mencakup: Transaction.
// @description     <br/><b>Developer:</b> Muhammad Brilian Satria Utama
// @description     <b>Environment:</b> Development
// @contact.name    API Support
// @contact.email   rian3903@gmail.com
// @license.name    Internal Use Only
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html
// @host            localhost:8082
// @BasePath        /transaction/v1
// @schemes http https
func ServeHTTP() {
	dependency := dependencyInject()

	r := gin.Default()

	v1 := r.Group("/transaction/v1")

	authorized := v1
	authorized.Use(dependency.MiddlewareRefreshToken)
	{
		v1.POST("/create", dependency.MiddlewareRefreshToken, dependency.TransactionHandler.Create)
		v1.POST("/update-status/:reference", dependency.MiddlewareRefreshToken)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := helper.GetEnv("PORT", "8082")
	err := r.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
