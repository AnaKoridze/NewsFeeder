package config

import (
	"strings"
	"github.com/spf13/viper"
)

type AppConfig struct {

	// global
	Port int

	// postgresql
	PostgresqlHost     string
	PostgresqlPort     int
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbName   string
}

func LoadConfig() AppConfig {

	viper.AutomaticEnv()
	viper.SetEnvPrefix("test")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// define default values
	viper.SetDefault("port", 8086)

	viper.SetDefault("db.host", "host")
	viper.SetDefault("db.port", 5432)
	viper.SetDefault("db.user", "user")
	viper.SetDefault("db.pwd", "pwd")
	viper.SetDefault("db.name", "name")


	// load configuration
	cfg := AppConfig{}
	cfg.Port = viper.GetInt("port")

	cfg.PostgresqlHost = viper.GetString("db.host")
	cfg.PostgresqlPort = viper.GetInt("db.port")
	cfg.PostgresqlUser = viper.GetString("db.user")
	cfg.PostgresqlPassword = viper.GetString("db.pwd")
	cfg.PostgresqlDbName = viper.GetString("db.name")

	return cfg
}