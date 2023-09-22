CREATE TABLE t_income_history
(
    created_at DECIMAL NOT NULL,
    created_by VARCHAR(64),
    updated_at DECIMAL NOT NULL,
    updated_by VARCHAR(64),
    deleted_at DECIMAL,
    deleted_by VARCHAR(64)
)