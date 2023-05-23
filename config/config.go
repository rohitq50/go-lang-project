package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Configuration struct {
	DBHost string `env:"MONGO_DB_HOST, required"`
	DBName string `env:"MONGO_DB_NAME, required"`
}

func NewConfig(files ...string) *Configuration {

	// Load config from env file
	err := godotenv.Load(files...)

	if err != nil {
		log.Printf("No .env file found: %q\n", files)
	}

	config := Configuration{}

	// Parse env to configuration
	err = env.Parse(&config)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	return &config
}
