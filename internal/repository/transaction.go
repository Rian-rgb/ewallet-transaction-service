package repository

import (
	"ewallet-transaction/internal/domain/transaction"
	"gorm.io/gorm"
)

type TransactionRepo struct {
	DB *gorm.DB
}

func (repo *TransactionRepo) Save(entity *transaction.Entity) error {
	return repo.DB.Save(entity).Error
}

//func (repo *TransactionRepo) FindByReference(reference string) (transactionEntity *transaction.Entity, err error) {
//
//	err = repo.DB.
//		Where("reference = ?", reference).
//		First(&transactionEntity).Error
//
//	return transactionEntity, err
//}
//
//func (repo *TransactionRepo) UpdateStatus(reference string, status string, additionalInfo string) error {
//
//	return repo.DB.Exec(
//		"UPDATE transactions SET transaction_status = ?, additional_info = ? WHERE reference = ?",
//		status,
//		additionalInfo,
//		reference,
//	).Error
//}
