package model

import (
	"errors"
	"fmt"
	"os"

	"github.com/ardanlabs/conf/v3"
	"github.com/joho/godotenv"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variables.
type Config struct {
	DB DBConfig

	ListenPort uint `conf:"env:LISTEN_PORT,required"`

	MigrationPath string `conf:"env:MIGRATION_PATH,required"`

	LogLevel string `conf:"env:LOG_LEVEL,default:debug"`
}

type DBConfig struct {
	DBUser      string `conf:"env:DB_USER,required"`
	DBPassword  string `conf:"env:DB_PASSWORD,required"`
	DBHost      string `conf:"env:DB_HOST,required"`
	DBPort      uint   `conf:"env:DB_PORT,required"`
	DBName      string `conf:"env:DB_NAME,required"`
	TLSDisabled bool   `conf:"env:DB_TLS_DISABLED"`
}

// LoadConfig reads configuration from file or environment variables.
func (c *Config) LoadConfig() error {
	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load()
		if err != nil {
			return fmt.Errorf("failed to load .env file: %w", err)
		}
	}

	_, err := conf.Parse("", c)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			return err
		}

		return err
	}

	return nil
}