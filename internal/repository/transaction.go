package repository

import (
	"ewallet-transaction/internal/domain/transaction"
	"ewallet-transaction/internal/errs"
	"gorm.io/gorm"
)

type TransactionRepo struct {
	DB *gorm.DB
}

func (r *TransactionRepo) Save(trx *transaction.Entity) error {
	err := r.DB.Save(trx).Error
	if err != nil {
		return errs.Wrap(
			errs.ErrInternal,
			"failed to create transaction",
			err,
		)
	}

	return nil
}

func (r *TransactionRepo) FindByReference(reference string) (*transaction.Entity, error) {

	var resp transaction.Entity

	err := r.DB.
		Where("reference = ?", reference).
		First(&resp).Error

	return &resp, err
}

func (r *TransactionRepo) UpdateStatus(reference string, status string, additionalInfo string) error {

	return r.DB.Exec(
		"UPDATE transactions SET transaction_status = ?, additional_info = ? WHERE reference = ?",
		status,
		additionalInfo,
		reference,
	).Error
}
