package schema

import (
	"time"

	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"
	"github.com/oklog/ulid/v2"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

type RequestCreateOrUpdateSpendingHistory struct {
	SpendingTypeID          string `json:"spending_type_id"`
	PaymentMethodID         string `json:"payment_method_id"`
	PaymentName             string `json:"payment_name"`
	SpendingAmount          int    `json:"spending_amount"`
	Description             string `json:"description"`
	TimeSpendingHistory     string `json:"time_spending_history"`
	ShowTimeSpendingHistory string `json:"show_time_spending_history"`
}

func (req *RequestCreateOrUpdateSpendingHistory) Validate() error {
	err := map[string][]string{}

	if errParse := util.ParseUlid(req.SpendingTypeID); errParse != nil {
		err["spending_type_id"] = append(err["spending_type_id"], "invalid spending type id")
	}

	if req.PaymentName == "" && req.PaymentMethodID == "" {
		err["payment_name"] = append(err["payment_name"], "payment_name or payment_method_id salah satu tidak boleh kosong required")
		err["payment_method_id"] = append(err["payment_method_id"], "payment_name or payment_method_id salah satu tidak boleh kosong required")
	}

	if req.PaymentName != "" && req.PaymentMethodID != "" {
		err["payment_name"] = append(err["payment_name"], "payment_name or payment_method_id pilih salah satu field")
		err["payment_method_id"] = append(err["payment_method_id"], "payment_name or payment_method_id pilih salah satu field")
	}

	if req.PaymentMethodID != "" {
		if _, errParse := ulid.Parse(req.PaymentMethodID); errParse != nil {
			err["payment_method_id"] = append(err["payment_method_id"], "invalid payment method id")
		}
	}

	if req.PaymentName != "" {
		paymentName := util.MaxMinString(req.PaymentName, 3, 22)
		if paymentName != "" {
			err["payment_name"] = append(err["payment_name"], paymentName)
		}
	}

	spendingAmount := util.MaxMinNumeric(req.SpendingAmount, 1000, 100000000)
	if spendingAmount != "" {
		err["spending_amount"] = append(err["spending_amount"], spendingAmount)
	}

	description := util.MaxMinString(req.Description, 3, 55)
	if description != "" {
		err["description"] = append(err["description"], description)
	}

	_, errParse := time.Parse("2006-01-02", req.TimeSpendingHistory)
	if errParse != nil {
		err["time_spending_history"] = append(err["time_spending_history"], "invalid date")
	}

	showTimeSpendingHistory := util.MaxMinString(req.ShowTimeSpendingHistory, 3, 32)
	if showTimeSpendingHistory != "" {
		err["show_time_spending_history"] = append(err["show_time_spending_history"], showTimeSpendingHistory)
	}

	if len(err) != 0 {
		return _error.HttpErrMapOfSlices(err, response.CM06)
	}

	return nil
}
