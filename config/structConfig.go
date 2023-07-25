package config

import "time"

type Config struct {
	Application struct {
		Port     string `mapstructure:"PORT"`
		FeWebUrl string `mapstructure:"FE_WEB_URL"`

		GraceFul struct {
			MaxSecond time.Duration `mapstructure:"MAX_SECOND"`
		} `mapstructure:"GRACEFUL"`
	} `mapstructure:"APPLICATION"`

	Default struct {
		DefaultImage string `mapstructure:"DEFAULT_IMAGE"`
		Aes          struct {
			CFBkey   string `mapstructure:"CFB_KEY"`
			CBCkey   string `mapstructure:"CBC_KEY"`
			CBCIVkey string `mapstructure:"CBC_IV_KEY"`
		} `mapstructure:"AES"`
	} `mapstructure:"DEFAULT"`

	DB struct {
		Postgres struct {
			Host string `mapstructure:"HOST"`
			Port int    `mapstructure:"PORT"`
			User string `mapstructure:"USER"`
			Pass string `mapstructure:"PASS"`
			Name string `mapstructure:"NAME"`
		} `mapstructure:"POSTGRESQL"`
	} `mapstructure:"DB"`
}
