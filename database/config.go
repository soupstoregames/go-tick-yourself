package database

type Config struct {
	User     string `env:"DB_USER" default:"soupstoregames"`
	Password string `env:"DB_PASSWORD" default:"twitch2020"`
	Host     string `env:"DB_HOST" default:"postgres"`
	Port     int    `env:"DB_PORT" default:"5432"`
	SSL      bool   `env:"DB_SSL" default:"false"`
}
