package util

import (
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
	uuidSatori "github.com/satori/go.uuid"
)

const DayNow = "hari-ini"
const Kemarin = "kemarin"
const MingguNow = "minggu-ini"
const BulanNow = "bulan-ini"

func NewUlid() string {
	return ulid.Make().String()
}

func NewUUID() string {
	return uuidSatori.NewV4().String()
}

func ParseUlid(u string) error {
	if _, err := ulid.Parse(u); err != nil {
		log.Info().Msgf("failed parse ulid | err : %v", err)
		return err
	}

	return nil
}

func ParseUUID(u string) error {
	if _, err := uuid.Parse(u); err != nil {
		log.Info().Msgf("failed parse uuid | err : %v", err)
		return err
	}

	return nil
}
