package config

import (
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() error {
	const ENV_DEV_FILE_NAME = ".env.dev"
	const ENV_PROD_FILE_NAME = ".env.prod"

	if err := godotenv.Load(); err != nil {
		return err
	}

	env := os.Getenv("ENV")
	switch env {
	case "prod":
		if err := godotenv.Load(ENV_PROD_FILE_NAME); err != nil {
			return err
		}

	case "dev":
		if err := godotenv.Load(ENV_DEV_FILE_NAME); err != nil {
			return err
		}
	}

	return nil
}
