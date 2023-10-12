package schema

import (
	"time"
)

type ResponseSpendingHistory struct {
	ID                      string    `json:"id"`
	ProfileID               string    `json:"profile_id"`
	SpendingTypeID          string    `json:"spending_type_id"`
	SpendingTypeTitle       string    `json:"spending_type_title"`
	PaymentMethodID         *string   `json:"payment_method_id"`
	PaymentMethodName       *string   `json:"payment_method_name"`
	PaymentName             *string   `json:"payment_name"`
	BeforeBalance           int       `json:"before_balance"`
	SpendingAmount          int       `json:"spending_amount"`
	FormatSpendingAmount    string    `json:"format_spending_amount"`
	AfterBalance            int       `json:"after_balance"`
	Description             string    `json:"description"`
	TimeSpendingHistory     time.Time `json:"time_spending_history"`
	ShowTimeSpendingHistory string    `json:"show_time_spending_history"`
}
