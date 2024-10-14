package repository

import (
	"context"
	"errors"

	"github.com/arvinpaundra/technical-test-api-mnc/internal/model"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/constant"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/dbutil"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Insert(ctx context.Context, transaction model.Transaction) error {
	err := r.db.WithContext(ctx).
		Model(&model.Transaction{}).
		Create(transaction).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepository) FindOne(ctx context.Context, opts ...dbutil.QueryOption) (model.Transaction, error) {
	var transaction model.Transaction

	err := r.db.WithContext(ctx).
		Model(&model.Transaction{}).
		Scopes(dbutil.ApplyScopes(opts...)).
		Take(&transaction).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Transaction{}, constant.ErrTransactionNotFound
		}

		return model.Transaction{}, err
	}

	return transaction, nil
}

func (r *TransactionRepository) Update(ctx context.Context, transaction model.Transaction, opts ...dbutil.QueryOption) error {
	panic("err")
}
