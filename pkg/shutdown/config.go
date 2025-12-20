package shutdown

import "time"

type Config struct {
	Timeout time.Duration `yaml:"timeout" env:"TIMEOUT" env-default:"5s"`
}
