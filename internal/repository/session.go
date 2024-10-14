package repository

import (
	"context"
	"errors"

	"github.com/arvinpaundra/technical-test-api-mnc/internal/model"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/constant"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/dbutil"
	"gorm.io/gorm"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (r *SessionRepository) Insert(ctx context.Context, session model.Session) error {
	err := r.db.WithContext(ctx).
		Model(&model.Session{}).
		Create(session).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (r *SessionRepository) FindOne(ctx context.Context, opts ...dbutil.QueryOption) (model.Session, error) {
	var session model.Session

	err := r.db.WithContext(ctx).
		Model(&model.Session{}).
		Scopes(dbutil.ApplyScopes(opts...)).
		Take(&session).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Session{}, constant.ErrSessionNotFound
		}

		return model.Session{}, err
	}

	return session, nil
}
