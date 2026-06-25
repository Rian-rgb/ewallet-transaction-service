package cmd

import (
	"ewallet-transaction/external/user"
	"ewallet-transaction/helper"
	"ewallet-transaction/internal/domain/transaction"
	"ewallet-transaction/internal/handler"
	"ewallet-transaction/internal/repository"
	"ewallet-transaction/internal/service"
)

type Dependency struct {
	TransactionHandler transaction.IHandler
	UserClient         *user.UserServiceClient
}

func dependencyInject() Dependency {

	transactionRepo := &repository.TransactionRepo{
		DB: helper.DB,
	}

	// Transaction
	transactionService := &service.TransactionService{
		TransactionRepo: transactionRepo,
	}
	transactionHandler := &handler.TransactionHandler{
		TransactionService: transactionService,
	}

	umsAddress := helper.GetEnv("UMS_GRPC_ADDRESS", "localhost:7000")
	userClient := user.NewUserServiceClient(umsAddress)

	return Dependency{
		TransactionHandler: transactionHandler,
		UserClient:         userClient,
	}
}
