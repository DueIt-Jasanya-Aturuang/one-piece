package validation

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/src/modules/services"
	"github.com/go-playground/validator/v10"
)

func MsgForTag(tag, param string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "min":
		return fmt.Sprintf("This Field Min Character %s", param)
	case "max":
		return fmt.Sprintf("This Field Max Character %s", param)
	}
	return ""
}

func MustUnique(field validator.FieldLevel, ctx context.Context, db *sql.DB, service *services.PaymentServiceImpl) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		payment, err := service.PaymentRepository.GetPaymentByName(ctx, db, value)
		if err != nil {
			return false
		}

		if payment != nil {
			return false
		}

	}
	return true
}
