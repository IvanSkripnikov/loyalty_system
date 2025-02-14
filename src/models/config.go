package models

import (
	"os"
	"strconv"
)

type Config struct {
	Database           Database
	RedirectUrl        string
	RedirectServiceUrl string
}

func LoadConfig() (*Config, error) {
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}

	return &Config{
		Database: Database{
			Address:  os.Getenv("DB_ADDRESS"),
			Port:     dbPort,
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DB:       os.Getenv("DB_NAME"),
		},
		RedirectUrl:        os.Getenv("REDIRECT_URL"),
		RedirectServiceUrl: os.Getenv("REDIRECT_SERViCE_URL"),
	}, nil
}

func GetRequiredVariables() []string {
	return []string{
		// Обязательные переменные окружения для подключения к БД сервиса
		"DB_ADDRESS", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME",

		// Обязательные переменные для редиректа в сервис авторизации
		"REDIRECT_URL", "REDIRECT_SERViCE_URL",
	}
}
