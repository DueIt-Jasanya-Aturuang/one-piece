CREATE TABLE m_spending_type
(
    id            VARCHAR(64) NOT NULL UNIQUE PRIMARY KEY,
    profile_id    VARCHAR(64),
    title         VARCHAR(64),
    maximum_limit DECIMAL,
    created_at    DECIMAL     NOT NULL,
    created_by    VARCHAR(64),
    updated_at    DECIMAL     NOT NULL,
    updated_by    VARCHAR(64),
    deleted_at    DECIMAL,
    deleted_by    VARCHAR(64),
    constraint fk_m_profile
        foreign key (profile_id)
            references m_profiles (id)
            on delete cascade
            on update cascade
);