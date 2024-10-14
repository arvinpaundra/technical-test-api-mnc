package repository

import (
	"context"
	"errors"

	"github.com/arvinpaundra/technical-test-api-mnc/internal/model"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/constant"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/dbutil"
	"gorm.io/gorm"
)

type TransactionHistoryRepository struct {
	db *gorm.DB
}

func NewTransactionHistoryRepository(db *gorm.DB) *TransactionHistoryRepository {
	return &TransactionHistoryRepository{db: db}
}

func (r *TransactionHistoryRepository) Insert(ctx context.Context, history model.TransactionHistory) error {
	err := r.db.WithContext(ctx).
		Model(&model.TransactionHistory{}).
		Create(history).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionHistoryRepository) FindOne(ctx context.Context, opts ...dbutil.QueryOption) (model.TransactionHistory, error) {
	var history model.TransactionHistory

	err := r.db.WithContext(ctx).
		Model(&model.TransactionHistory{}).
		Scopes(dbutil.ApplyScopes(opts...)).
		Take(&history).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.TransactionHistory{}, constant.ErrTransactionHistoryNotFound
		}

		return model.TransactionHistory{}, err
	}

	return history, nil
}

func (r *TransactionHistoryRepository) FindAll(ctx context.Context, opts ...dbutil.QueryOption) ([]model.TransactionHistory, error) {
	var histories []model.TransactionHistory

	err := r.db.WithContext(ctx).
		Model(&model.TransactionHistory{}).
		Scopes(dbutil.ApplyScopes(opts...)).
		Find(&histories).
		Error

	if err != nil {
		return []model.TransactionHistory{}, err
	}

	return histories, nil
}
