package repository

import (
	"context"
	"errors"

	"github.com/arvinpaundra/technical-test-api-mnc/internal/model"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/constant"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/dbutil"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Insert(ctx context.Context, user model.User) error {
	err := r.db.WithContext(ctx).
		Model(&model.User{}).
		Create(user).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindOne(ctx context.Context, opts ...dbutil.QueryOption) (model.User, error) {
	var user model.User

	err := r.db.WithContext(ctx).
		Model(&model.User{}).
		Scopes(dbutil.ApplyScopes(opts...)).
		Take(&user).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, constant.ErrUserNotFound
		}

		return model.User{}, err
	}

	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, user model.User, opts ...dbutil.QueryOption) error {
	err := r.db.WithContext(ctx).
		Model(&model.User{}).
		Scopes(dbutil.ApplyScopes(opts...)).
		Updates(user).
		Error

	if err != nil {
		return err
	}

	return nil
}
