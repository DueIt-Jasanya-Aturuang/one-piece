package setup

import (
	"github.com/ory/dockertest/v3"
	"github.com/rs/zerolog/log"
)

func SetupDocker() *dockertest.Pool {
	dockerpool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatal().Msgf("failed create docker pool | err %v", err)
	}
	err = dockerpool.Client.Ping()
	if err != nil {
		log.Fatal().Msgf("failed ping docker pool | err %v", err)
	}

	return dockerpool
}
