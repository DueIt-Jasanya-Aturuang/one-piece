package schema

import (
	"time"
)

type ResponseIncomeHistory struct {
	ID                    string    `json:"id"`
	ProfileID             string    `json:"profile_id"`
	IncomeTypeID          string    `json:"income_type_id"`
	IncomeTypeTitle       string    `json:"income_type_title"`
	PaymentMethodID       *string   `json:"payment_method_id"`
	PaymentMethodName     *string   `json:"payment_method_name"`
	PaymentName           *string   `json:"payment_name"`
	IncomeAmount          int       `json:"income_amount"`
	FormatIncomeAmount    string    `json:"format_income_amount"`
	Description           string    `json:"description"`
	TimeIncomeHistory     time.Time `json:"time_income_history"`
	ShowTimeIncomeHistory string    `json:"show_time_income_history"`
}
