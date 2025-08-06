package configuration

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

const (
	MESSAGE_ERROR_LOADING_ENV   = "Error loading the environment variables: %v"
	MESSAGE_SUCCESS_LOADING_ENV = "Success loading the environment variables"
)

type Configuration struct {
	Port         string
	DbConfig     DatabaseConfig
	JwtAlgorithm string
	JwtSecretKey string
}

type DatabaseConfig struct {
	DbHost     string
	DbUser     string
	DbPassword string
	DbPort     string
	DbSSLMode  string
	DbName     string
}

var config *Configuration

func LoadConfig() *Configuration {
	envFile := filepath.Join(".", ".env")

	err := godotenv.Load(envFile)

	if err != nil {
		fmt.Errorf(MESSAGE_ERROR_LOADING_ENV, err)
		return nil
	}

	fmt.Println(MESSAGE_SUCCESS_LOADING_ENV)

	config = &Configuration{
		Port: os.Getenv("PORT"),
		DbConfig: DatabaseConfig{
			DbHost:     os.Getenv("DB_HOST"),
			DbUser:     os.Getenv("DB_USER"),
			DbPassword: os.Getenv("DB_PASSWORD"),
			DbPort:     os.Getenv("DB_PORT"),
			DbSSLMode:  os.Getenv("DB_SSL_MODE"),
			DbName:     os.Getenv("DB_NAME"),
		},
	}
	return config
}

func GetConfiguration() *Configuration {
	return config
}
