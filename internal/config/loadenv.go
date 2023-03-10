package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() (map[string]string, error) {

	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	log.Println(os.Environ())

	return make(map[string]string, 0), nil
}
