package entities

import "database/sql"

type Payment struct {
	Id          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Image       string         `json:"image"`
	CreatedAt   int64          `json:"created_at"`
	CreatedBy   string         `json:"created_by"`
	UpdatedAt   int64          `json:"updated_at"`
	UpdatedBy   sql.NullString `json:"updated_by"`
	DeletedAt   sql.NullInt64  `json:"deleted_at"`
	DeletedBy   sql.NullString `json:"deleted_by"`
}
