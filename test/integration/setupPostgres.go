package integration

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/rs/zerolog/log"
)

func Postgres(dockerPool *dockertest.Pool) (*dockertest.Resource, *sql.DB, string) {
	resource, err := dockerPool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "11",
		Env: []string{
			"POSTGRES_PASSWORD=pw",
			"POSTGRES_USER=user_name",
			"POSTGRES_DB=dueit_db",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Err(err).Msgf("Could not start resource: %s", err)
		os.Exit(1)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf("postgres://user_name:pw@%s/dueit_db?sslmode=disable", hostAndPort)

	log.Info().Msgf("Connecting to database on url: %s", databaseUrl)

	err = resource.Expire(120)
	if err != nil {
		log.Err(err).Msgf("Could not connect to docker: %s", err)
		os.Exit(1)
	} // Tell docker to hard kill the container in 120 seconds

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	dockerPool.MaxWait = 120 * time.Second
	var db *sql.DB
	if err = dockerPool.Retry(func() error {
		db, err = sql.Open("postgres", databaseUrl)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Err(err).Msgf("Could not connect to docker: %s", err)
		os.Exit(1)
	}

	return resource, db, databaseUrl
}
