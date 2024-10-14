package factory

import (
	"github.com/arvinpaundra/technical-test-api-mnc/config"
	"github.com/arvinpaundra/technical-test-api-mnc/internal/repository"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/database"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/logger"

	"go.uber.org/zap"
)

type Factory struct {
	UserRepository               *repository.UserRepository
	SessionRepository            *repository.SessionRepository
	TransactionRepository        *repository.TransactionRepository
	TransactionHistoryRepository *repository.TransactionHistoryRepository
	Logger                       *zap.Logger
}

func NewFactory() *Factory {
	db := database.GetConnection()

	return &Factory{
		// repositories
		UserRepository:               repository.NewUserRepository(db),
		SessionRepository:            repository.NewSessionRepository(db),
		TransactionRepository:        repository.NewTransactionRepository(db),
		TransactionHistoryRepository: repository.NewTransactionHistoryRepository(db),

		// logger
		Logger: logger.NewLogger(config.GetAppEnv()),
	}
}
