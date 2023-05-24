package loadEnv

import (
	"autoSql-cobra/utills/runDir"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() (err error) {
	if err = godotenv.Load(".env.active"); err != nil {
		err = godotenv.Load(fmt.Sprintf("%s/.env.active", runDir.RunDir()))
		if err != nil {
			return err
		}
	}

	activeEnv := os.Getenv("ACTIVE")

	if activeEnv == "dev" {
		if err = godotenv.Load(".env.dev"); err != nil {
			err = godotenv.Load(fmt.Sprintf("%s/.env.dev", runDir.RunDir()))
			if err != nil {
				return err
			}
		}
	}

	if activeEnv == "prod" {
		if err = godotenv.Load(".env.prod"); err != nil {
			err = godotenv.Load(fmt.Sprintf("%s/.env.prod", runDir.RunDir()))
			if err != nil {
				return err
			}
		}
	}

	return err
}
