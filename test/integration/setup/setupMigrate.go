package setup

import (
	"database/sql"
	"os"

	"github.com/rs/zerolog/log"
)

func Migrate(db *sql.DB) {
	log.Info().Msg("start migrate")
	_, err := db.Exec("CREATE SCHEMA IF NOT EXISTS dueit;")
	if err != nil {
		log.Err(err).Msgf("Failed to create schema dueit: %s", err)
		os.Exit(1)
	}

	// create table m_payment_methods
	createMPaymentMethodTable(db)

	// create table m_profiles
	createMProfilesTable(db)

	// create data m_profiles
	createMProfilesData(db)

	// create table m_spending_type
	createMSpendingTypeTable(db)
	createMDefaultSpendingTypeTable(db)
	createMDefaultSpendingTypeData(db)

	// create table t_spending_history
	createMSpendingHistoryTable(db)

}

func createMPaymentMethodTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS m_payment_methods (
    id 			VARCHAR(64) NOT NULL UNIQUE PRIMARY KEY,
    name        VARCHAR(128),
    description TEXT,
    image       VARCHAR(255) default 'default-icon.png',
    created_at  DECIMAL     NOT NULL,
    created_by  VARCHAR(64),
    updated_at  DECIMAL     NOT NULL,
    updated_by  VARCHAR(64),
    deleted_at  DECIMAL,
    deleted_by  VARCHAR(64)
    )`)
	if err != nil {
		log.Err(err).Msgf("Failed to create table m_payment_methods: %s", err)
		os.Exit(1)
	}
}

func createMProfilesTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS m_profiles (
   	   id VARCHAR(64) NOT NULL UNIQUE PRIMARY KEY,
	   user_id VARCHAR(64),
	   quotes VARCHAR(128),
	   profesi VARCHAR(60),
	   created_at DECIMAL NOT NULL,
	   created_by VARCHAR(64),
	   updated_at DECIMAL NOT NULL,
	   updated_by VARCHAR(64),
	   deleted_at DECIMAL,
	   deleted_by VARCHAR(64)
    )`)
	if err != nil {
		log.Err(err).Msgf("Failed to create table m_profiles: %s", err)
		os.Exit(1)
	}

}

func createMProfilesData(db *sql.DB) {
	// create data profiles
	_, err := db.Exec(`INSERT INTO m_profiles (id, user_id, quotes, profesi, created_at, created_by, updated_at) 
			 VALUES ('profileID1', 'userID1', null, null, 0, 'profileID1', 0)`)
	if err != nil {
		log.Err(err).Msgf("Failed to create data m_profiles: %s", err)
		os.Exit(1)
	}
}

func createMSpendingTypeTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS m_spending_type (
   	id            VARCHAR(64) NOT NULL UNIQUE PRIMARY KEY,
    profile_id    VARCHAR(64),
    title         VARCHAR(64),
    maximum_limit DECIMAL,
    icon 		  VARCHAR(255),
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
    )`)
	if err != nil {
		log.Err(err).Msgf("Failed to create table m_spending_type: %s", err)
		os.Exit(1)
	}
}

func createMDefaultSpendingTypeTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS m_default_spending_type (
   	id            VARCHAR(64) NOT NULL UNIQUE PRIMARY KEY,
    title         VARCHAR(64),
    active        BOOLEAN,
    maximum_limit DECIMAL,
    icon          VARCHAR(64),
    created_at    DECIMAL     NOT NULL,
    created_by    VARCHAR(64),
    updated_at    DECIMAL     NOT NULL,
    updated_by    VARCHAR(64),
    deleted_at    DECIMAL,
    deleted_by    VARCHAR(64)
    )`)
	if err != nil {
		log.Err(err).Msgf("Failed to create table m_spending_type: %s", err)
		os.Exit(1)
	}
}

func createMDefaultSpendingTypeData(db *sql.DB) {
	_, err := db.Exec(`INSERT INTO m_default_spending_type
(id, title, active, maximum_limit, icon, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by)
VALUES('215d0320-fda1-4008-b8d6-f6f6e96b853b', 'makan', true, 100000, 'icon.png', 1684768516, 'admin', 1684768516, NULL, NULL, NULL);`)
	if err != nil {
		log.Err(err).Msgf("Failed to create table m_spending_type: %s", err)
		os.Exit(1)
	}

	_, err = db.Exec(`INSERT INTO m_default_spending_type
(id, title, active, maximum_limit, icon, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by)
VALUES('7101e34a-2c12-4d26-97c4-22c2ee4f4cfa', 'transportasi', true, 200000, 'icon.png', 1684768516, 'admin', 1684768516, NULL, NULL, NULL);`)
	if err != nil {
		log.Err(err).Msgf("Failed to create table m_spending_type: %s", err)
		os.Exit(1)
	}

	_, err = db.Exec(`INSERT INTO m_default_spending_type
(id, title, active, maximum_limit, icon, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by)
VALUES('c20d8e45-c246-4b55-ac1f-80e0a1c5d48c', 'loundry', true, 300000, 'icon.png', 1684768516, 'admin', 1684768516, NULL, NULL, NULL);`)
	if err != nil {
		log.Err(err).Msgf("Failed to create table m_spending_type: %s", err)
		os.Exit(1)
	}
}

func createMSpendingHistoryTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS t_spending_history (
   	id                         VARCHAR(64) NOT NULL UNIQUE PRIMARY KEY,
    profile_id                 VARCHAR(64),
    spending_type_id           VARCHAR(64),
    payment_method_id          VARCHAR(64) NULL,
    payment_name               VARCHAR(64) NULL,
    before_balance             DECIMAL,
    spending_amount            DECIMAL,
    after_balance              DECIMAL,
    description                TEXT,
    location                   VARCHAR(64),
    time_spending_history      TIMESTAMP,
    show_time_spending_history VARCHAR(64),
    created_at                 DECIMAL     NOT NULL,
    created_by                 VARCHAR(64),
    updated_at                 DECIMAL     NOT NULL,
    updated_by                 VARCHAR(64),
    deleted_at                 DECIMAL,
    deleted_by                 VARCHAR(64),
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
    )`)
	if err != nil {
		log.Err(err).Msgf("Failed to create table m_spending_type: %s", err)
		os.Exit(1)
	}
}
