package config

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	Dev  = "dev"
	Prod = "prod"
)

func InitEnv() error {
	const DevEnvFileName = ".env.dev"
	const ProdEnvFileName = ".env.prod"

	if err := godotenv.Load(); err != nil {
		return err
	}

	env := os.Getenv("ENV")
	switch env {
	case Dev:
		if err := godotenv.Load(DevEnvFileName); err != nil {
			return err
		}

	case Prod:
		if err := godotenv.Load(ProdEnvFileName); err != nil {
			return err
		}
	}

	return nil
}

func IsDev() bool {
	return os.Getenv("ENV") == Dev
}

func IsProd() bool {
	return os.Getenv("ENV") == Prod
}
