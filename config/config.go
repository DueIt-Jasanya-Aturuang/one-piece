package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
)

var (
	cfg Config
	do  sync.Once
)

func Get() Config {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf(fmt.Sprintf("cannot read file env : %v", err))
	}

	do.Do(func() {
		if err := viper.Unmarshal(&cfg); err != nil {
			log.Fatalf(fmt.Sprintf("cannot unmarsahl config : %v", err))
		}
	})
	return cfg
}
