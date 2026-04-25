package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env   string      `yaml:"env"`
	GRPC  GRPCConfig  `yaml:"grpc"`
	Token TokenConfig `yaml:"token"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type TokenConfig struct {
	Ttl time.Duration `yaml:"ttl"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exists:" + path)
	}
	var config Config
	if err := cleanenv.ReadConfig(path, &config); err != nil {
		panic("failed to read config")
	}
	return &config
}

func fetchConfigPath() string {
	var res string
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}
