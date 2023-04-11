package loadEnv

import (
	"autoSql-cobra/utills/getRunDir"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() (err error) {
	if err = godotenv.Load(".active-env"); err != nil {
		err = godotenv.Load(fmt.Sprintf("%s/.active-env", getRunDir.RunDir()))
		if err != nil {
			return err
		}
	}

	activeEnv := os.Getenv("active")

	if activeEnv == "dev" {
		if err = godotenv.Load(".env.dev"); err != nil {
			err = godotenv.Load(fmt.Sprintf("%s/.env.dev", getRunDir.RunDir()))
			if err != nil {
				return err
			}
		}
	}

	if activeEnv == "prod" {
		if err = godotenv.Load(".env.prod"); err != nil {
			err = godotenv.Load(fmt.Sprintf("%s/.env.prod", getRunDir.RunDir()))
			if err != nil {
				return err
			}
		}
	}

	return err
}
