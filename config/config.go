package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	//mengumpulkan seluruh konfigurasi di golang
	ENV  string `env:"ENV" envDefault:"dev"`
	PORT string `env:"PORT" envDefault:"8080"`
	JWTConfig JWTConfig `envPrefix:"JWT_"`
	MySQLConfig MySQLConfig `envPrefix:"MYSQL_"`
	SMTPConfig SMTPConfig `envPrefix:"SMTP_"`
}

type SMTPConfig struct{
	Host string `env:"HOST" envDefault:"localhost"` 
	Port int64 `env:"PORT" envDefault:"587"` 
	Username string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
}

type JWTConfig struct{
	SecretKey string `env:"SECRET_KEY"`
}

type MySQLConfig struct{
	Host string `env:"HOST" envDefault:"localhost"`
	Port string `env:"PORT" envDefault:"3306"`
	User string `env:"USER" envDefault:"root"`
	Password string `env:"PASSWORD"`
	Database string `env:"DATABASE"`
}

func NewConfig(path string) (*Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	cfg := new(Config)
	err = env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}