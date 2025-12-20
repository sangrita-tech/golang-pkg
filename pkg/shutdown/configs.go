package shutdown

import "time"

type Configs struct {
	Timeout time.Duration `yaml:"timeout" env:"TIMEOUT" env-default:"5s"`
}
