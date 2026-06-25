package transaction

type IRepository interface {
	Save(trx *Entity) error
	FindByReference(reference string) (*Entity, error)
	UpdateStatus(reference string, status string, additionalInfo string) error
}
