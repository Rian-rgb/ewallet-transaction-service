package transaction

import "context"

type IService interface {
	CreateTransaction(
		ctx context.Context,
		transactionEntity *Entity,
	) (*Entity, error)
}
