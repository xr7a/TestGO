package config

import "os"

type Env string

const ServiceEnvVarName = "TEST_ENV"

const (
	Local = Env("local")
	Dev   = Env("dev")
	Prod  = Env("prod")
)

func GetEnv() Env {
	env := os.Getenv(ServiceEnvVarName)
	if env == "" {
		return Local
	}
	return Env(env)
}
