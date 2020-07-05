package config

import "os"

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresHost     string
	PostgresPort     string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			PostgresUser:     os.Getenv("POSTGRES_USER"),
			PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
			PostgresHost:     os.Getenv("POSTGRES_HOST"),
			PostgresDB:       os.Getenv("POSTGRES_DB"),
			PostgresPort:     os.Getenv("POSTGRES_PORT"),
		},
	}
}
