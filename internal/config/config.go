package config

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	DBHost     string `env:"DB_HOST" env-required:"true"`
	DBPort     int    `env:"DB_PORT" env-required:"true"`
	DBPassword string `env:"DB_PASS" env-required:"true"`
	DBUser     string `env:"DB_USER" env-required:"true"`
	DBName     string `env:"DB_NAME" env-required:"true"`
	SecretKey  string `env:"SECRET_KEY" env-required:"true"`
	LogLevel   string `env:"LOG_LEVEL" env-required:"true"`
}

func NewConfig() (*config, error) {
	var cfg config

	if err := cleanenv.ReadConfig(".env", &cfg); err != nil {
		log.Println(err)

		if err := cleanenv.ReadEnv(&cfg); err != nil {
			return nil, fmt.Errorf("ENV not found")
		}
	}

	return &cfg, nil
}

func (c *config) GetUrlPostgres() string {
	return fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName)
}

func (c *config) GetLogLevel() slog.Level {
	switch c.LogLevel {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		slog.Warn("Invalid log level", "level", c.LogLevel)
		return slog.LevelDebug
	}
}
