package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	DBHost     string `env:"DB_HOST" env-required:"true"`
	DBPort     int    `env:"DB_PORT" env-required:"true"`
	DBPassword string `env:"DB_PASS" env-required:"true"`
	DBUser     string `env:"DB_USER" env-required:"true"`
	DBName     string `env:"DB_NAME" env-required:"true"`
	SecretKey  string `env:"SECRET_KEY" env-required:"true"`
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

func (c config) GetUrlPostgres() string {
	return fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName)
}
