package _usecase

import (
	"errors"
)

var NamePaymentExist = errors.New("nama payment sudah tersedia")
var PaymentNotExist = errors.New("payment tidak tersedia")
var SpendingHistoryNotFound = errors.New("spending history tidak ditemukan")
var InvalidSpendingTypeID = errors.New("invalid spending history id")
var TitleSpendingTypeExist = errors.New("title kategori sudah tersedia")
var SpendingTypeNotFound = errors.New("kategori tidak di temukan")
var ProfileIDNotFound = errors.New("invalid profile id")
var InvalidTimestamp = errors.New("invalid timestamp")
var InvalidPaymentMethodID = errors.New("invalid payment method id")
var BalanceNotExist = errors.New("balance tidak ada")
