package schema

import (
	"time"

	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

type RequestCreateOrUpdateIncomeHistory struct {
	IncomeTypeID          string `json:"income_type_id"`
	PaymentMethodID       string `json:"payment_method_id"`
	PaymentName           string `json:"payment_name"`
	IncomeAmount          int    `json:"income_amount"`
	Description           string `json:"description"`
	TimeIncomeHistory     string `json:"time_income_history"`
	ShowTimeIncomeHistory string `json:"show_time_income_history"`
}

func (req *RequestCreateOrUpdateIncomeHistory) Validate() error {
	err := map[string][]string{}

	if errParse := util.ParseUlid(req.IncomeTypeID); errParse != nil {
		err["income_type_id"] = append(err["income_type_id"], "invalid income type id")
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
		if errParse := util.ParseUlid(req.PaymentMethodID); errParse != nil {
			err["payment_method_id"] = append(err["payment_method_id"], "invalid payment method id")
		}
	}

	if req.PaymentName != "" {
		paymentName := util.MaxMinString(req.PaymentName, 3, 22)
		if paymentName != "" {
			err["payment_name"] = append(err["payment_name"], paymentName)
		}
	}

	incomeAmount := util.MaxMinNumeric(req.IncomeAmount, 1000, 100000000)
	if incomeAmount != "" {
		err["income_amount"] = append(err["income_amount"], incomeAmount)
	}

	description := util.MaxMinString(req.Description, 3, 55)
	if description != "" {
		err["description"] = append(err["description"], description)
	}

	_, errParse := time.Parse("2006-01-02", req.TimeIncomeHistory)
	if errParse != nil {
		err["time_income_history"] = append(err["time_income_history"], "invalid date")
	}

	showTimeIncomeHistory := util.MaxMinString(req.ShowTimeIncomeHistory, 3, 32)
	if showTimeIncomeHistory != "" {
		err["show_time_income_history"] = append(err["show_time_income_history"], showTimeIncomeHistory)
	}

	if len(err) != 0 {
		return _error.HttpErrMapOfSlices(err, response.CM06)
	}

	return nil
}
