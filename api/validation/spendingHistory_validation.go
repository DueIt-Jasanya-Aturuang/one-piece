package validation

import (
	"time"

	"github.com/google/uuid"
	errResp "github.com/jasanya-tech/jasanya-response-backend-golang"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
)

func CreateSpendingHistory(req *domain.RequestCreateSpendingHistory) error {
	err := map[string][]string{}

	if _, err := uuid.Parse(req.ProfileID); err != nil {
		return errResp.HttpErrString("invalid profile id", errResp.S403)
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
		return errResp.HttpErrMapOfSlices(err, errResp.S400)
	}

	return nil
}

func UpdateSpendingHistory(req *domain.RequestUpdateSpendingHistory) error {
	err := map[string][]string{}

	if _, errParse := uuid.Parse(req.ProfileID); errParse != nil {
		return errResp.HttpErrString("invalid profile id", errResp.S403)
	}
	if _, errParse := uuid.Parse(req.SpendingTypeID); errParse != nil {
		err["spending_type_id"] = append(err["spending_type_id"], "invalid spending type id")
	}
	if _, errParse := uuid.Parse(req.ID); errParse != nil {
		return errResp.HttpErrString("spending history tidak ditemukan", errResp.S404)
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
		return errResp.HttpErrMapOfSlices(err, errResp.S400)
	}

	return nil
}
