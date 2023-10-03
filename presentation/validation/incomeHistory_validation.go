package validation

import (
	"time"

	"github.com/google/uuid"
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"
	"github.com/oklog/ulid/v2"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
)

func CreateIncomeHistory(req *domain.RequestCreateIncomeHistory) error {
	err := map[string][]string{}

	if _, err := uuid.Parse(req.ProfileID); err != nil {
		return _error.HttpErrString("invalid profile id", response.CM05)
	}

	if _, errParse := ulid.Parse(req.IncomeTypeID); errParse != nil {
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
		if _, errParse := ulid.Parse(req.PaymentMethodID); errParse != nil {
			err["payment_method_id"] = append(err["payment_method_id"], "invalid payment method id")
		}
	}

	if req.PaymentName != "" {
		paymentName := maxMinString(req.PaymentName, 3, 22)
		if paymentName != "" {
			err["payment_name"] = append(err["payment_name"], paymentName)
		}
	}

	incomeAmount := maxMinNumeric(req.IncomeAmount, 1000, 100000000)
	if incomeAmount != "" {
		err["income_amount"] = append(err["income_amount"], incomeAmount)
	}

	description := maxMinString(req.Description, 3, 55)
	if description != "" {
		err["description"] = append(err["description"], description)
	}

	_, errParse := time.Parse("2006-01-02", req.TimeIncomeHistory)
	if errParse != nil {
		err["time_income_history"] = append(err["time_income_history"], "invalid date")
	}

	showTimeIncomeHistory := maxMinString(req.ShowTimeIncomeHistory, 3, 32)
	if showTimeIncomeHistory != "" {
		err["show_time_income_history"] = append(err["show_time_income_history"], showTimeIncomeHistory)
	}

	if len(err) != 0 {
		return _error.HttpErrMapOfSlices(err, response.CM06)
	}

	return nil
}

func UpdateIncomeHistory(req *domain.RequestUpdateIncomeHistory) error {
	err := map[string][]string{}

	if _, errParse := uuid.Parse(req.ProfileID); errParse != nil {
		return _error.HttpErrString("invalid profile id", response.CM05)
	}

	if _, errParse := ulid.Parse(req.IncomeTypeID); errParse != nil {
		err["income_type_id"] = append(err["income_type_id"], "invalid income type id")
	}

	if _, errParse := ulid.Parse(req.ID); errParse != nil {
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
		if _, errParse := ulid.Parse(req.PaymentMethodID); errParse != nil {
			err["payment_method_id"] = append(err["payment_method_id"], "invalid payment method id")
		}
	}

	if req.PaymentName != "" {
		paymentName := maxMinString(req.PaymentName, 3, 22)
		if paymentName != "" {
			err["payment_name"] = append(err["payment_name"], paymentName)
		}
	}

	incomeAmount := maxMinNumeric(req.IncomeAmount, 1000, 100000000)
	if incomeAmount != "" {
		err["income_amount"] = append(err["income_amount"], incomeAmount)
	}

	description := maxMinString(req.Description, 3, 55)
	if description != "" {
		err["description"] = append(err["description"], description)
	}

	_, errParse := time.Parse("2006-01-02", req.TimeIncomeHistory)
	if errParse != nil {
		err["time_income_history"] = append(err["time_income_history"], "invalid date")
	}

	showTimeIncomeHistory := maxMinString(req.ShowTimeIncomeHistory, 3, 32)
	if showTimeIncomeHistory != "" {
		err["show_time_income_history"] = append(err["show_time_income_history"], showTimeIncomeHistory)
	}

	if len(err) != 0 {
		return _error.HttpErrMapOfSlices(err, response.CM06)
	}

	return nil
}
