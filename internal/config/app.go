package config

import "fmt"

type AppEnv string

const (
	AppEnvDev  AppEnv = "dev"
	AppEnvProd AppEnv = "prod"
)

func parseAppEnv(value string) (AppEnv, error) {
	env := AppEnv(value)

	switch env {
	case AppEnvDev, AppEnvProd:
		return env, nil
	default:
		return "", fmt.Errorf("invalid APP_ENV: %q", value)
	}
}

type AppConfig struct {
	Name     string
	Env      AppEnv
	Version  string
	RootPath string
}

func loadAppEnv() (AppConfig, error) {
	name, err := requiredEnv("APP_NAME")
	if err != nil {
		return AppConfig{}, err
	}

	rawEnv, err := requiredEnv("APP_ENV")
	if err != nil {
		return AppConfig{}, err
	}

	env, err := parseAppEnv(rawEnv)
	if err != nil {
		return AppConfig{}, err
	}

	version, err := requiredEnv("APP_VERSION")
	if err != nil {
		return AppConfig{}, err
	}

	rootPath, err := requiredEnv("APP_ROOTPATH")
	if err != nil {
		return AppConfig{}, err
	}

	return AppConfig{
		Name:     name,
		Env:      env,
		Version:  version,
		RootPath: rootPath,
	}, nil
}
