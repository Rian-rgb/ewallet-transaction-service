package transaction

type IRepository interface {
	Save(entity *Entity) error
	//FindByReference(reference string) (transactionEntity *Entity, err error)
	//UpdateStatus(reference string, status string, additionalInfo string) error
}
