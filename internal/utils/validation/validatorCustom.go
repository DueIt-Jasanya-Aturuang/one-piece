package validation

import (
	"context"
	"database/sql"
	"fmt"

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

func MustUnique(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		ctx := context.Background()
		db := &sql.DB{}
		payment := GetPaymentByName(ctx, db, value)

		if payment {
			return false
		}
	}
	return true
}

func GetPaymentByName(ctx context.Context, db *sql.DB, name string) bool {
	_, err := db.Exec(`set search_path='dueit'`)
	if err != nil {
		return false
	}

	SQL := "SELECT id, name, description, image, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM m_payment_methods WHERE name = $1 LIMIT 1"
	_, err = db.QueryContext(ctx, SQL, name)
	if err != nil {
		return false
	}

	return true
}
