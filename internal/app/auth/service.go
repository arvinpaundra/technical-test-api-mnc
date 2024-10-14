package auth

import (
	"context"
	"errors"
	"time"

	"github.com/arvinpaundra/technical-test-api-mnc/internal/dto/request"
	"github.com/arvinpaundra/technical-test-api-mnc/internal/dto/response"
	"github.com/arvinpaundra/technical-test-api-mnc/internal/factory"
	"github.com/arvinpaundra/technical-test-api-mnc/internal/model"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/constant"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/dbutil"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/util"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Insert(ctx context.Context, user model.User) error
	FindOne(ctx context.Context, opts ...dbutil.QueryOption) (model.User, error)
}

type SessionRepository interface {
	Insert(ctx context.Context, session model.Session) error
	FindOne(ctx context.Context, opts ...dbutil.QueryOption) (model.Session, error)
}

type Service struct {
	userRepository    UserRepository
	sessionRepository SessionRepository
	logger            *zap.Logger
}

func NewService(f *factory.Factory) *Service {
	return &Service{
		userRepository:    f.UserRepository,
		sessionRepository: f.SessionRepository,
		logger:            f.Logger.With(zap.String("domain", "auth")),
	}
}

func (s *Service) Register(ctx context.Context, payload request.Register) (response.Register, error) {
	user, err := s.userRepository.FindOne(ctx, dbutil.Select("id"), dbutil.Where("phone_number = ?", payload.PhoneNumber))
	if err != nil && !errors.Is(err, constant.ErrUserNotFound) {
		s.logger.Error(err.Error())
		return response.Register{}, err
	}

	if user != (model.User{}) {
		s.logger.With(zap.String("phone", payload.PhoneNumber)).Error(constant.ErrPhoneAlreadyExist.Error())
		return response.Register{}, constant.ErrPhoneAlreadyExist
	}

	pin, _ := bcrypt.GenerateFromPassword([]byte(payload.PIN), bcrypt.DefaultCost)

	user = model.User{
		ID:          util.GetUuid(),
		FirstName:   payload.FirstName,
		LastName:    payload.LastName,
		PhoneNumber: payload.PhoneNumber,
		Pin:         string(pin),
		Balance:     0,
		Address:     payload.Address,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}

	err = s.userRepository.Insert(ctx, user)
	if err != nil {
		s.logger.With(zap.Any("user_payload", user)).Error(err.Error())
		return response.Register{}, err
	}

	res := response.Register{
		UserId:      user.ID.String(),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		CreatedDate: user.CreatedDate.Format("2006-01-02 15:04:05"),
	}

	return res, nil
}

func (s *Service) Login(ctx context.Context, payload request.Login) (response.Login, error) {
	user, err := s.userRepository.FindOne(ctx, dbutil.Select("id, pin"), dbutil.Where("phone_number = ?", payload.PhoneNumber))
	if err != nil {
		s.logger.Error(err.Error())
		return response.Login{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Pin), []byte(payload.PIN))
	if err != nil {
		s.logger.Error(err.Error())
		return response.Login{}, constant.ErrUserNotFound
	}

	accessToken, err := util.GenerateJWT(user.ID.String(), time.Hour*12)
	if err != nil {
		s.logger.With(zap.String("jwt_creation", "access_token")).Error(err.Error())
		return response.Login{}, err
	}

	refreshToken, err := util.GenerateJWT(user.ID.String(), time.Hour*72)
	if err != nil {
		s.logger.With(zap.String("jwt_creation", "refresh_token")).Error(err.Error())
		return response.Login{}, err
	}

	session := model.Session{
		ID:           util.GetUuid(),
		UserId:       user.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		CreatedDate:  time.Now(),
		UpdatedDate:  time.Now(),
	}

	err = s.sessionRepository.Insert(ctx, session)
	if err != nil {
		s.logger.With(zap.Any("session_payload", session)).Error(err.Error())
		return response.Login{}, err
	}

	res := response.Login{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return res, nil
}

func (s *Service) Authenticate(ctx context.Context, userId string) (response.Authenticate, error) {
	user, err := s.userRepository.FindOne(ctx, dbutil.Select("id"), dbutil.Where("id = ?", userId))
	if err != nil {
		s.logger.Error(err.Error())
		return response.Authenticate{}, err
	}

	res := response.Authenticate{
		UserId: user.ID.String(),
	}

	return res, nil
}

func (s *Service) RenewRefreshToken(ctx context.Context) error {
	panic("error")
}
