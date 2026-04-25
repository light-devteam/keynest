package config

import "time"

type TokenConfig struct {
	Ttl time.Duration `yaml:"ttl"`
}
