package config

type Config struct {
	App AppConfig
}

func Load() (Config, error) {
	app, err := loadAppEnv()
	if err != nil {
		return Config{}, err
	}

	return Config{
		App: app,
	}, nil
}
