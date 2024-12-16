package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-required:"true"`
	Driver     string `yaml:"driver"`
	Dns        string `yaml:"dns"`
	Root       string `yaml:"root"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost"`
	Port        string        `yaml:"port" env-default:"8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle-timeout" env-default:"60s"`
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
