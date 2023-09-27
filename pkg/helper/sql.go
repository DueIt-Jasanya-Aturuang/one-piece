package helper

import (
	"database/sql"
)

func NewNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}

	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func GetNullString(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}

	return nil
}

func NewNullBool(b bool) sql.NullBool {
	if !b {
		return sql.NullBool{}
	}

	return sql.NullBool{
		Bool:  b,
		Valid: true,
	}
}

func GetNullBool(b sql.NullBool) *bool {
	if b.Valid {
		return &b.Bool
	}

	return nil
}

func NewNullInt64(i int64) sql.NullInt64 {
	if i == 0 {
		return sql.NullInt64{}
	}

	return sql.NullInt64{
		Int64: i,
		Valid: true,
	}
}

func GetNullInt64(i sql.NullInt64) *int64 {
	if i.Valid {
		return &i.Int64
	}

	return nil
}
