package validation

import (
	"time"

	"github.com/google/uuid"
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
)

func CreateSpendingHistory(req *domain.RequestCreateSpendingHistory) error {
	err := map[string][]string{}

	if _, err := uuid.Parse(req.ProfileID); err != nil {
		return _error.HttpErrString("invalid profile id", response.CM05)
	}
	if _, errParse := uuid.Parse(req.SpendingTypeID); errParse != nil {
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
		if _, errParse := uuid.Parse(req.PaymentMethodID); errParse != nil {
			err["payment_method_id"] = append(err["payment_method_id"], "invalid payment method id")
		}
	}

	if req.PaymentName != "" {
		paymentName := maxMinString(req.PaymentName, 3, 22)
		if paymentName != "" {
			err["payment_name"] = append(err["payment_name"], paymentName)
		}
	}

	spendingAmount := maxMinNumeric(req.SpendingAmount, 1000, 100000000)
	if spendingAmount != "" {
		err["spending_amount"] = append(err["spending_amount"], spendingAmount)
	}

	description := maxMinString(req.Description, 3, 55)
	if description != "" {
		err["description"] = append(err["description"], description)
	}

	location := maxMinString(req.Location, 3, 32)
	if location != "" {
		err["location"] = append(err["location"], location)
	}

	_, errParse := time.Parse("2006-01-02", req.TimeSpendingHistory)
	if errParse != nil {
		err["time_spending_history"] = append(err["time_spending_history"], "invalid date")
	}

	showTimeSpendingHistory := maxMinString(req.ShowTimeSpendingHistory, 3, 32)
	if showTimeSpendingHistory != "" {
		err["show_time_spending_history"] = append(err["show_time_spending_history"], showTimeSpendingHistory)
	}

	if len(err) != 0 {
		return _error.HttpErrMapOfSlices(err, response.CM06)
	}

	return nil
}

func UpdateSpendingHistory(req *domain.RequestUpdateSpendingHistory) error {
	err := map[string][]string{}

	if _, errParse := uuid.Parse(req.ProfileID); errParse != nil {
		return _error.HttpErrString("invalid profile id", response.CM05)
	}
	if _, errParse := uuid.Parse(req.SpendingTypeID); errParse != nil {
		err["spending_type_id"] = append(err["spending_type_id"], "invalid spending type id")
	}
	if _, errParse := uuid.Parse(req.ID); errParse != nil {
		return _error.HttpErrString(response.CodeCompanyName[response.CM01], response.CM01)
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
		if _, errParse := uuid.Parse(req.PaymentMethodID); errParse != nil {
			err["payment_method_id"] = append(err["payment_method_id"], "invalid payment method id")
		}
	}

	if req.PaymentName != "" {
		paymentName := maxMinString(req.PaymentName, 3, 22)
		if paymentName != "" {
			err["payment_name"] = append(err["payment_name"], paymentName)
		}
	}

	spendingAmount := maxMinNumeric(req.SpendingAmount, 1000, 100000000)
	if spendingAmount != "" {
		err["spending_amount"] = append(err["spending_amount"], spendingAmount)
	}

	description := maxMinString(req.Description, 3, 55)
	if description != "" {
		err["description"] = append(err["description"], description)
	}

	location := maxMinString(req.Location, 3, 32)
	if location != "" {
		err["location"] = append(err["location"], location)
	}

	_, errParse := time.Parse("2006-01-02", req.TimeSpendingHistory)
	if errParse != nil {
		err["time_spending_history"] = append(err["time_spending_history"], "invalid date")
	}

	showTimeSpendingHistory := maxMinString(req.ShowTimeSpendingHistory, 3, 32)
	if showTimeSpendingHistory != "" {
		err["show_time_spending_history"] = append(err["show_time_spending_history"], showTimeSpendingHistory)
	}

	if len(err) != 0 {
		return _error.HttpErrMapOfSlices(err, response.CM06)
	}

	return nil
}
