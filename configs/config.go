package configs

type Config struct {

	App struct {
		Env string `envconfig:"APP_ENV"`
	}

	Auth struct {
		Access string `envconfig:"ACCESS_SECRET"`
		Api_key string `envconfig:"API_KEY"`
	}
}