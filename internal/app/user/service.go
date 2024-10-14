package user

import (
	"context"
	"time"

	"github.com/arvinpaundra/technical-test-api-mnc/internal/dto/request"
	"github.com/arvinpaundra/technical-test-api-mnc/internal/dto/response"
	"github.com/arvinpaundra/technical-test-api-mnc/internal/factory"
	"github.com/arvinpaundra/technical-test-api-mnc/internal/model"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/dbutil"
	"go.uber.org/zap"
)

type UserRepository interface {
	FindOne(ctx context.Context, opts ...dbutil.QueryOption) (model.User, error)
	Update(ctx context.Context, user model.User, opts ...dbutil.QueryOption) error
}

type Service struct {
	userRepository UserRepository
	logger         *zap.Logger
}

func NewService(f *factory.Factory) *Service {
	return &Service{
		userRepository: f.UserRepository,
		logger:         f.Logger.With(zap.String("domain", "user")),
	}
}

func (s *Service) UpdateProfile(ctx context.Context, userId string, payload request.UpdateProfile) (response.UpdateProfile, error) {
	_, err := s.userRepository.FindOne(ctx, dbutil.Select("id"), dbutil.Where("id = ?", userId))
	if err != nil {
		s.logger.Error(err.Error())
		return response.UpdateProfile{}, err
	}

	user := model.User{
		FirstName:   payload.FirstName,
		LastName:    payload.LastName,
		Address:     payload.Address,
		UpdatedDate: time.Now(),
	}

	err = s.userRepository.Update(
		ctx,
		user,
		dbutil.Select("first_name", "last_name", "address", "updated_date"),
		dbutil.Where("id = ?", userId),
	)

	if err != nil {
		s.logger.With(zap.Any("updated_user_payload", user)).Error(err.Error())
		return response.UpdateProfile{}, err
	}

	res := response.UpdateProfile{
		UserId:      userId,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Address:     user.Address,
		UpdatedDate: user.UpdatedDate.Format("2006-01-02 15:04:05"),
	}

	return res, nil
}
