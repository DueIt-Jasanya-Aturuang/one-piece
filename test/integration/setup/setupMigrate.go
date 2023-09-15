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
