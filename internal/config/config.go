package config

import (
	"AnlossAPI/pkg/env"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	Env          string `yaml:"env" env-default:"local"`
	RecordsPath  string `yaml:"records_path" env-default:"./storage/records.db"`
	StudentsPath string `yaml:"students_path" env-required:"true"`
	HTTPServer   `yaml:"http_server"`
}

type HTTPServer struct {
	Port        string `yaml:"address" env-default:"8080"`
	Timeout     string `yaml:"timeout" env-default:"4s"`
	IdleTimeout string `yaml:"idle_timeout" env-default:"5m"`
}

func MustLoad() *Config {
	configPath := env.GetValue("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("config path not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("config path does not exist:", configPath)
	}

	var config Config

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatal("can`t read config", err)
	}

	return &config
}
