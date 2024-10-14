package transaction

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/arvinpaundra/technical-test-api-mnc/internal/dto/request"
	"github.com/arvinpaundra/technical-test-api-mnc/internal/dto/response"
	"github.com/arvinpaundra/technical-test-api-mnc/internal/factory"
	"github.com/arvinpaundra/technical-test-api-mnc/internal/model"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/constant"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/dbutil"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/util"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type UserRepository interface {
	FindOne(ctx context.Context, opts ...dbutil.QueryOption) (model.User, error)
	Update(ctx context.Context, user model.User, opts ...dbutil.QueryOption) error
}

type TransactionRepository interface {
	Insert(ctx context.Context, transaction model.Transaction) error
	FindOne(ctx context.Context, opts ...dbutil.QueryOption) (model.Transaction, error)
	Update(ctx context.Context, transaction model.Transaction, opts ...dbutil.QueryOption) error
}

type TransactionHistoryRepository interface {
	Insert(ctx context.Context, history model.TransactionHistory) error
	FindOne(ctx context.Context, opts ...dbutil.QueryOption) (model.TransactionHistory, error)
	FindAll(ctx context.Context, opts ...dbutil.QueryOption) ([]model.TransactionHistory, error)
}

type TxBeginner interface {
	Begin(opts ...*sql.TxOptions) error
}

type Service struct {
	userRepository               UserRepository
	transactionRepository        TransactionRepository
	transactionHistoryRepository TransactionHistoryRepository
	logger                       *zap.Logger
}

func NewService(f *factory.Factory) *Service {
	return &Service{
		userRepository:               f.UserRepository,
		transactionRepository:        f.TransactionRepository,
		transactionHistoryRepository: f.TransactionHistoryRepository,
		logger:                       f.Logger.With(zap.String("domain", "transaction")),
	}
}

func (s *Service) Topup(ctx context.Context, userId string, payload request.Topup) (response.Topup, error) {
	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		s.logger.Error(err.Error())
		return response.Topup{}, err
	}

	topup := model.Transaction{
		ID:          util.GetUuid(),
		UserId:      parsedUserId,
		Amount:      payload.Amount,
		Category:    "topup",
		Status:      "PENDING",
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}

	err = s.transactionRepository.Insert(ctx, topup)
	if err != nil {
		s.logger.With(zap.Any("topup_payload", topup)).Error(err.Error())
		return response.Topup{}, err
	}

	latestTransaction, err := s.transactionHistoryRepository.FindOne(
		ctx,
		dbutil.Select("id", "balance_after"),
		dbutil.Where("user_id = ?", parsedUserId),
		dbutil.Order("created_date DESC"),
	)

	if err != nil && !errors.Is(err, constant.ErrTransactionHistoryNotFound) {
		s.logger.Error(err.Error())
		return response.Topup{}, err
	}

	previousBalance := latestTransaction.BalanceAfter
	currentBalance := payload.Amount + previousBalance

	transactionHistory := model.TransactionHistory{
		ID:              util.GetUuid(),
		UserId:          parsedUserId,
		TransactionType: "CREDIT",
		Amount:          payload.Amount,
		BalanceBefore:   previousBalance,
		BalanceAfter:    currentBalance,
		ReferenceId:     topup.ID.String(),
		CreatedDate:     time.Now(),
		UpdatedDate:     time.Now(),
	}

	err = s.transactionHistoryRepository.Insert(ctx, transactionHistory)
	if err != nil {
		s.logger.With(zap.Any("transaction_history_payload", transactionHistory)).Error(err.Error())
		return response.Topup{}, err
	}

	user := model.User{
		Balance: currentBalance,
	}

	err = s.userRepository.Update(ctx, user, dbutil.Where("id = ?", userId))
	if err != nil {
		s.logger.With(zap.Any("user_payload", user)).Error(err.Error())
		return response.Topup{}, err
	}

	res := response.Topup{
		TopupId:       topup.ID.String(),
		AmountTopup:   topup.Amount,
		BalanceBefore: previousBalance,
		BalanceAfter:  currentBalance,
		CreatedDate:   topup.CreatedDate.Format("2006-01-02 15:04:05"),
	}

	return res, nil
}

func (s *Service) Payment(ctx context.Context, userId string, payload request.Payment) (response.Payment, error) {
	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		s.logger.Error(err.Error())
		return response.Payment{}, err
	}

	latestTransaction, err := s.transactionHistoryRepository.FindOne(
		ctx,
		dbutil.Select("id", "balance_after"),
		dbutil.Where("user_id = ?", parsedUserId),
		dbutil.Order("created_date DESC"),
	)

	if err != nil && !errors.Is(err, constant.ErrTransactionHistoryNotFound) {
		s.logger.Error(err.Error())
		return response.Payment{}, err
	}

	previousBalance := latestTransaction.BalanceAfter
	currentBalance := previousBalance - payload.Amount

	if currentBalance < 0 {
		s.logger.Error(constant.ErrInsufficientBalance.Error())
		return response.Payment{}, constant.ErrInsufficientBalance
	}

	payment := model.Transaction{
		ID:          util.GetUuid(),
		UserId:      parsedUserId,
		Amount:      payload.Amount,
		Category:    "payment",
		Status:      "PENDING",
		Remarks:     payload.Remarks,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}

	err = s.transactionRepository.Insert(ctx, payment)
	if err != nil {
		s.logger.With(zap.Any("payment_payload", payment)).Error(err.Error())
		return response.Payment{}, err
	}

	transactionHistory := model.TransactionHistory{
		ID:              util.GetUuid(),
		UserId:          parsedUserId,
		TransactionType: "DEBIT",
		Amount:          payload.Amount,
		BalanceBefore:   previousBalance,
		BalanceAfter:    currentBalance,
		ReferenceId:     payment.ID.String(),
		Remarks:         payload.Remarks,
		CreatedDate:     time.Now(),
		UpdatedDate:     time.Now(),
	}

	err = s.transactionHistoryRepository.Insert(ctx, transactionHistory)
	if err != nil {
		s.logger.With(zap.Any("transaction_history_payload", transactionHistory)).Error(err.Error())
		return response.Payment{}, err
	}

	user := model.User{
		Balance: currentBalance,
	}

	err = s.userRepository.Update(ctx, user, dbutil.Select("balance"), dbutil.Where("id = ?", userId))
	if err != nil {
		s.logger.With(zap.Any("user_payload", user)).Error(err.Error())
		return response.Payment{}, err
	}

	res := response.Payment{
		PaymentId:     payment.ID.String(),
		Amount:        payment.Amount,
		Remarks:       payment.Remarks,
		BalanceBefore: previousBalance,
		BalanceAfter:  currentBalance,
		CreatedDate:   payment.CreatedDate.Format("2006-01-02 15:04:05"),
	}

	return res, nil
}

func (s *Service) Transfer(ctx context.Context, userId string, payload request.Transfer) (response.Transfer, error) {
	parsedSenderUserId, err := uuid.Parse(userId)
	if err != nil {
		s.logger.Error(err.Error())
		return response.Transfer{}, err
	}

	senderLatestTransaction, err := s.transactionHistoryRepository.FindOne(
		ctx,
		dbutil.Select("id", "balance_after"),
		dbutil.Where("user_id = ?", parsedSenderUserId),
		dbutil.Order("created_date DESC"),
	)

	if err != nil && !errors.Is(err, constant.ErrTransactionHistoryNotFound) {
		s.logger.Error(err.Error())
		return response.Transfer{}, err
	}

	senderPreviousBalance := senderLatestTransaction.BalanceAfter
	senderCurrentBalance := senderPreviousBalance - payload.Amount

	if senderCurrentBalance < 0 {
		s.logger.Error(constant.ErrInsufficientBalance.Error())
		return response.Transfer{}, constant.ErrInsufficientBalance
	}

	parsedReceiverUserId, err := uuid.Parse(payload.TargetUser)
	if err != nil {
		s.logger.Error(err.Error())
		return response.Transfer{}, err
	}

	senderTransfer := model.Transaction{
		ID:          util.GetUuid(),
		UserId:      parsedSenderUserId,
		TargetUser:  parsedReceiverUserId,
		Amount:      payload.Amount,
		Remarks:     payload.Remarks,
		Category:    "transfer",
		Status:      "PENDING",
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}

	err = s.transactionRepository.Insert(ctx, senderTransfer)
	if err != nil {
		s.logger.With(zap.Any("sender_transfer_payload", senderTransfer)).Error(err.Error())
		return response.Transfer{}, err
	}

	senderTransactionHistory := model.TransactionHistory{
		ID:              util.GetUuid(),
		UserId:          parsedSenderUserId,
		TransactionType: "DEBIT",
		Amount:          payload.Amount,
		BalanceBefore:   senderPreviousBalance,
		BalanceAfter:    senderCurrentBalance,
		Remarks:         payload.Remarks,
		ReferenceId:     senderTransfer.ID.String(),
		CreatedDate:     time.Now(),
		UpdatedDate:     time.Now(),
	}

	err = s.transactionHistoryRepository.Insert(ctx, senderTransactionHistory)
	if err != nil {
		s.logger.With(zap.Any("sender_transaction_history_payload", senderTransactionHistory)).Error(err.Error())
		return response.Transfer{}, err
	}

	senderUpdateBalance := model.User{
		Balance: senderCurrentBalance,
	}

	err = s.userRepository.Update(ctx, senderUpdateBalance, dbutil.Select("balance"), dbutil.Where("id = ?", parsedSenderUserId))
	if err != nil {
		s.logger.With(zap.Any("sender_update_balance_payload", senderUpdateBalance)).Error(err.Error())
		return response.Transfer{}, err
	}

	receiverTransfer := model.Transaction{
		ID:          util.GetUuid(),
		UserId:      parsedReceiverUserId,
		Amount:      payload.Amount,
		Remarks:     payload.Remarks,
		Category:    "transfer",
		Status:      "PENDING",
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}

	err = s.transactionRepository.Insert(ctx, receiverTransfer)
	if err != nil {
		s.logger.With(zap.Any("receiver_transfer_payload", receiverTransfer)).Error(err.Error())
		return response.Transfer{}, err
	}

	receiverLatestTransaction, err := s.transactionHistoryRepository.FindOne(
		ctx,
		dbutil.Select("id", "balance_after"),
		dbutil.Where("user_id = ?", payload.TargetUser),
		dbutil.Order("created_date DESC"),
	)

	if err != nil && !errors.Is(err, constant.ErrTransactionHistoryNotFound) {
		s.logger.Error(err.Error())
		return response.Transfer{}, err
	}

	receiverPreviousBalance := receiverLatestTransaction.BalanceAfter
	receiverCurrentBalance := receiverPreviousBalance + payload.Amount

	receiverTransactionHistory := model.TransactionHistory{
		ID:              util.GetUuid(),
		UserId:          parsedReceiverUserId,
		TransactionType: "CREDIT",
		Amount:          payload.Amount,
		BalanceBefore:   receiverPreviousBalance,
		BalanceAfter:    receiverCurrentBalance,
		Remarks:         payload.Remarks,
		ReferenceId:     receiverTransfer.ID.String(),
		CreatedDate:     time.Now(),
		UpdatedDate:     time.Now(),
	}

	err = s.transactionHistoryRepository.Insert(ctx, receiverTransactionHistory)
	if err != nil {
		s.logger.With(zap.Any("receiver_transaction_history_payload", receiverTransactionHistory)).Error(err.Error())
		return response.Transfer{}, err
	}

	receiverUpdateBalance := model.User{
		Balance: receiverCurrentBalance,
	}

	err = s.userRepository.Update(ctx, receiverUpdateBalance, dbutil.Select("balance"), dbutil.Where("id = ?", parsedReceiverUserId))
	if err != nil {
		s.logger.With(zap.Any("receiver_update_balance_payload", receiverUpdateBalance)).Error(err.Error())
		return response.Transfer{}, err
	}

	res := response.Transfer{
		TransferId:    senderTransfer.ID.String(),
		Amount:        senderTransfer.Amount,
		Remarks:       senderTransfer.Remarks,
		BalanceBefore: senderPreviousBalance,
		BalanceAfter:  senderCurrentBalance,
		CreatedDate:   senderTransfer.CreatedDate.Format("2006-01-02 15:04:05"),
	}

	return res, nil
}

func (s *Service) GetTransactions(ctx context.Context, userId string) ([]response.Transaction, error) {
	transactions, err := s.transactionHistoryRepository.FindAll(
		ctx,
		dbutil.Where("user_id = ?", userId),
		dbutil.Preload("Transaction"),
		dbutil.Order("created_date DESC"),
	)

	if err != nil {
		s.logger.Error(err.Error())
		return []response.Transaction{}, err
	}

	var res []response.Transaction

	for _, transaction := range transactions {
		res = append(res, response.Transaction{
			TransferId:      transaction.ID.String(),
			Status:          transaction.Transaction.Status,
			UserId:          transaction.UserId.String(),
			TransactionType: transaction.TransactionType,
			Amount:          transaction.Amount,
			Remarks:         transaction.Remarks,
			BalanceBefore:   transaction.BalanceBefore,
			BalanceAfter:    transaction.BalanceAfter,
			CreatedDate:     transaction.CreatedDate.Format("2006-01-02 15:04:05"),
		})
	}

	return res, nil
}
