package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-required:"true"`
	Storage    string `yaml:"storage"`
	HTTPServer `yaml:"http_server"`
	Pages      `yaml:"pages"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost"`
	Port        string        `yaml:"port" env-default:"8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle-timeout" env-default:"60s"`
}

type Pages struct {
	EqualsPage  string `yaml:"equals"`
	MolarPage   string `yaml:"molar"`
	BalancePage string `yaml:"balance"`
}

func LoadConfig() Config {
	var config Config
	config_path := os.Getenv("CONFIG_PATH")
	if config_path == "" {
		log.Fatal("variable CONFIG_PATH is not set")
	}
	if _, err := os.Stat(config_path); os.IsNotExist(err) {
		log.Fatal("file does not exist")
	}
	if err := cleanenv.ReadConfig(config_path, &config); err != nil {
		log.Fatal("parsing error")
	}
	return config
}
