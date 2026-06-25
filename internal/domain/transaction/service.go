package transaction

type IService interface {
	CreateTransaction(tx *Entity) (*Entity, error)
}
