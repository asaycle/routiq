package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"golang.org/x/xerrors"
)

type Config struct {
	Env    string `env:"ENV" env-default:"development"`
	Server ServerConfig
	DB     DBConfig
	App    AppConfig
}

type ServerConfig struct {
	Port string `env:"BACKEND_PORT" env-default:"50051"`
}

type DBConfig struct {
	Scheme   string `env:"DB_SCHEME" env-default:"postgres"`
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Port     string `env:"DB_PORT" env-default:"5432"`
	User     string `env:"DB_USER" env-default:"root"`
	Password string `env:"DB_PSWD" env-default:"root"`
	Database string `env:"DB_NAME" env-default:"routiq"`
}

func (c *DBConfig) GetURL() string {
	return c.Scheme + "://" + c.User + ":" + c.Password + "@" + c.Host + ":" + c.Port + "/" + c.Database + "?sslmode=disable"
}

type AppConfig struct {
	AuthConfig     AuthConfig
	LocationConfig LocationConfig
}

type AuthConfig struct {
	GoogleClientID     string `env:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `env:"GOOGLE_CLIENT_SECRET"`
	RedirectURI        string `env:"REDIRECT_URI"`
	JWTSecret          string `env:"JWT_SECRET"`
}

type LocationConfig struct {
	SourceURL string `env:"ROUTIQ_APP_LOCATION_SOURCE_URL" env-default:"https://nlftp.mlit.go.jp/ksj/gml/codelist/AdminiBoundary_CD.xlsx"`
}

func Load() (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, xerrors.Errorf("Error ReadEnv: %w", err)
	}
	return &cfg, nil
}
