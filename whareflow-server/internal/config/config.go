package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string      `yaml:"env" env-default:"local"`
	StoragePath StoragePath `yaml:"storage_path" env-required:"true"`
	HTTPServer  HTTPServer  `yaml:"http_server" env-required:"true"`
	Auth        Auth        `yaml:"auth" env-required:"true"`
	QR          QR          `yaml:"qr" env-required:"true"`
}

type StoragePath struct {
	Postgres Postgres `yaml:"postgres" env-required:"true"`
}

type Postgres struct {
	Password string `yaml:"password" env-required:"true"`
	Host     string `yaml:"host" env-required:"true"`
	Port     string `yaml:"port" env-required:"true"`
	Database string `yaml:"database" env-required:"true"`
}

type HTTPServer struct {
	Address      string        `yaml:"address"`
	Timeout      time.Duration `yaml:"timeout"`
	IddleTimeout time.Duration `yaml:"iddle_timeout"`
}

type Auth struct {
	PasswordSalt       string `yaml:"pass_salt"`
	ExpAccessToken     int    `yaml:"access_token_expiry_hour"`
	ExpRefreshToken    int    `yaml:"refresh_token_expiry_hour"`
	SecretAccessToken  string `yaml:"access_token_secret"`
	SecretRefreshToken string `yaml:"refresh_token_secret"`
}

type QR struct {
	UrlFrontend string `yaml:"url_frontend"`
	PathToFile  string `yaml:"path_to_file"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_WARE_FLOW")
	if configPath == "" {
		log.Fatal("CONFIG PATH not found ")
	}

	//check file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s %s", configPath, err)
	}

	var config Config

	err := cleanenv.ReadConfig(configPath, &config)
	if err != nil {
		log.Fatalf("can't read config: %s", configPath)
	}

	return &config
}
