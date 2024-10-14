package constant

import "errors"

var (
	ErrUserNotFound               = errors.New("user not found")
	ErrSessionNotFound            = errors.New("session not found")
	ErrTransactionNotFound        = errors.New("transaction not found")
	ErrTransactionHistoryNotFound = errors.New("transaction history not found")
	ErrPhoneAndPinNotMatch        = errors.New("phone number and pin doesn't match")
	ErrPhoneAlreadyExist          = errors.New("phone number already registered")
	ErrInsufficientBalance        = errors.New("balance is not enough")
)
