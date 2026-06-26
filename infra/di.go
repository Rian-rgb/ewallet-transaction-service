package infra

import (
	"ewallet-transaction/external/ums"
	"ewallet-transaction/internal/domain/transaction"
	"ewallet-transaction/internal/handler"
	"ewallet-transaction/internal/repository"
	"ewallet-transaction/internal/service"
	pb "github.com/Rian-rgb/ewallet-proto/gen/token_validation/v1"
)

type Dependency struct {
	TransactionHdl transaction.IHandler
	UmsClient      *ums.Client
}

func DependencyInject(appDeps *AppDependencies) *Dependency {

	transactionRepo := &repository.TransactionRepo{
		DB: appDeps.PostgresDB,
	}

	transactionSvc := &service.TransactionService{
		TransactionRepo: transactionRepo,
	}

	transactionHdl := &handler.TransactionHandler{
		TransactionService: transactionSvc,
	}

	pbUmsClient := pb.NewTokenValidationServiceClient(appDeps.GrpcRegistry.UmsConn.Conn)
	umsGrpcClient := ums.NewClient(pbUmsClient)

	return &Dependency{
		TransactionHdl: transactionHdl,
		UmsClient:      umsGrpcClient,
	}
}
