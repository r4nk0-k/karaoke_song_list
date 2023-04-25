package env

import (
	"github.com/caarlos0/env/v6"
)

type Env struct {
	Mysql mysql `envPrefix:"MYSQL_"`
}

type mysql struct {
	DatabaseName string `env:"DATABASE_NAME"`
	Username     string `env:"USERNAME"`
	Password     string `env:"PASSWORD"`
	Host         string `env:"HOST"`
	Port         int    `env:"PORT"`
}

var parsedEnv = Env{}

func Get() Env {
	return parsedEnv
}

func Parse() error {
	if err := env.Parse(&parsedEnv); err != nil {
		return err
	}

	return nil
}
