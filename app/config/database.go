package config

import (
	"fmt"

	"github.com/IvanSkripnikov/loyalty_system/models"
)

func GetDatabaseConfig() models.DBConfig {
	return models.DBConfig{
		User: Env("MYSQL_USER", "user"),

		Password: Env("MYSQL_PASSWORD", "pass"),

		Protocol: Env("MYSQL_PROT", "tcp"),

		Address: Env("MYSQL_ADDR", "db") + ":3306",

		Database: Env("MYSQL_DATABASE", "test"),
	}
}

func GetDatabaseConnectionString(host string) string {
	config := GetDatabaseConfig()
	netAddr := fmt.Sprintf("%s(%s)", config.Protocol, host+":3306")

	return fmt.Sprintf("%s:%s@%s/%s?timeout=30s", config.User, config.Password, netAddr, config.Database)
}
