package config

type (
	Config struct {
		//	application configurations
		AppPort   string `mapstructure:"app_port"`
		AppMode   string `mapstructure:"app_mode"`
		JWTSecret string `mapstructure:"jwt_secret"`

		// postgres
		Postgres struct {
			DSN string `mapastructure:"dsn"`
		} `mapstructure:"postgres"`
	}
)
