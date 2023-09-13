CREATE TABLE t_spending_history
(
    id                VARCHAR(64) NOT NULL UNIQUE PRIMARY KEY,
    profile_id        VARCHAR(64),
    spending_type_id  VARCHAR(64),
    payment_method_id VARCHAR(64) NULL,
    payment_name      VARCHAR(64) NULL,
    before_balance    DECIMAL,
    spending_amount   DECIMAL,
    after_balance     DECIMAL,
    description       TEXT,
    created_at        DECIMAL     NOT NULL,
    created_by        VARCHAR(64),
    updated_at        DECIMAL     NOT NULL,
    updated_by        VARCHAR(64),
    deleted_at        DECIMAL,
    deleted_by        VARCHAR(64),
    constraint fk_m_profile
        foreign key (profile_id)
            references m_profiles (id)
            on delete cascade
            on update cascade,
    constraint fk_m_spending_type
        foreign key (spending_type_id)
            references m_spending_type (id)
            on delete cascade
            on update cascade,
    constraint fk_m_payment_method
        foreign key (payment_method_id)
            references m_payment_methods (id)
            on delete cascade
            on update cascade
);