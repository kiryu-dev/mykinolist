package app

import (
	"os"

	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

type dbConfig struct {
	Host     string
	Port     string
	Username string
	DBName   string
	Password string
	SSLMode  string
}

type Config struct {
	ListeningPort       string
	JWTAccessSecretKey  string
	JWTRefreshSecretKey string
	DB                  *dbConfig
}

func LoadConfig(configPath string) (*Config, error) {
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := gotenv.Load(); err != nil {
		return nil, err
	}
	config := &Config{
		ListeningPort:       viper.GetString("port"),
		JWTAccessSecretKey:  os.Getenv("JWT_ACCESS_SECRET_KEY"),
		JWTRefreshSecretKey: os.Getenv("JWT_REFRESH_SECRET_KEY"),
		DB: &dbConfig{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   viper.GetString("db.dbname"),
			SSLMode:  viper.GetString("db.sslmode"),
		},
	}
	return config, nil
}
