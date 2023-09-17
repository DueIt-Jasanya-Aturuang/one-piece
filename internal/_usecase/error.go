package _usecase

import (
	"errors"
)

// ProfileIDNotFound for SpendingHistoryUsecaseImpl delete
var ProfileIDNotFound = errors.New("invalid profile id")

// InvalidTimestamp for SpendingHistoryUsecaseImpl GetAllByTimeAndProfileID
var InvalidTimestamp = errors.New("invalid timestamp")

// InvalidPaymentMethodID for SpendingHistoryUsecaseImpl create, update
var InvalidPaymentMethodID = errors.New("invalid payment method id")
