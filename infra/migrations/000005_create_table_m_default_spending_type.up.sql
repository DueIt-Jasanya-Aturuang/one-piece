CREATE TABLE m_default_spending_type
(
    id            VARCHAR(64) NOT NULL UNIQUE PRIMARY KEY,
    title         VARCHAR(64),
    active        BOOLEAN,
    maximum_limit DECIMAL,
    created_at    DECIMAL     NOT NULL,
    created_by    VARCHAR(64),
    updated_at    DECIMAL     NOT NULL,
    updated_by    VARCHAR(64),
    deleted_at    DECIMAL,
    deleted_by    VARCHAR(64)
);

INSERT INTO m_default_spending_type
(id, title, active, maximum_limit, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by)
VALUES('215d0320-fda1-4008-b8d6-f6f6e96b853b', 'makan', true, 100000, 1684768516, NULL, 1684768516, NULL, NULL, NULL);

INSERT INTO m_default_spending_type
(id, title, active, maximum_limit, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by)
VALUES('7101e34a-2c12-4d26-97c4-22c2ee4f4cfa', 'transportasi', true, 200000, 1684768516, NULL, 1684768516, NULL, NULL, NULL);

INSERT INTO m_default_spending_type
(id, title, active, maximum_limit, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by)
VALUES('c20d8e45-c246-4b55-ac1f-80e0a1c5d48c', 'loundry', true, 300000, 1684768516, NULL, 1684768516, NULL, NULL, NULL);