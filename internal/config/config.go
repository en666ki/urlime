package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env    string `yaml:"env" env-required`
	Server Server `yaml:"server" env-required`
	DB     DB     `yaml:"db" env-required`
	Api    Api    `yaml:"api" env-required`
}

type Server struct {
	Host        string        `yaml:"host" env-required`
	Port        string        `yaml:"port" env-required`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type DB struct {
	Driver   string `yaml:"driver" env-required`
	DB       string `yaml:"db" env-required`
	Host     string `yaml:"host" env-required`
	Name     string `yaml:"name" env-required`
	Port     int    `yaml:"port" env-required`
	Table    string `yaml:"table" env-required`
	SslMode  string `yaml:"ssl_mode" env-default:"disable"`
	User     string `yaml:"user" env-required`
	Password string `yaml:"password" env-required`
}

type Api struct {
	Params Params `yaml:"params" env-required`
}

type Params struct {
	Shorten string `yaml:"shorten" env-required`
	Unshort string `yaml:"unshort" env-required`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("[ERROR] CONFIG_PATH is not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("[ERROR] No such config: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("[ERROR] Cannot read config: %v", err)
	}

	return &cfg
}

func TestConfig() *Config {
	return &Config{
		Env: "local",
		Server: Server{
			Host:        "localhost",
			Port:        "8080",
			Timeout:     time.Duration(4 * time.Second),
			IdleTimeout: time.Duration(60 * time.Second),
		},
		DB: DB{
			Driver:   "postgres",
			DB:       "postgres",
			Host:     "postgres",
			Name:     "local_short",
			Port:     5432,
			Table:    "local_urls",
			User:     "local",
			Password: "local_pwd",
		},
		Api: Api{
			Params: Params{
				Shorten: "url",
				Unshort: "surl",
			},
		},
	}
}
