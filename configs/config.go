package configs

import (
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

// LoadConfig loads the configuration from the config file.
func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.SetConfigFile(".")

	err := viper.ReadInConfig()
	if err != nil {
		// cast parse error to viper.ConfigFileNotFoundError
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg := new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}

	return nil
}

// GetDB returns the database configuration.
func GetDB() DBConfig {
	return cfg.DB
}

// GetServerPort returns the server port as a string.
func GetServerPort() string {
	return cfg.API.Port
}
