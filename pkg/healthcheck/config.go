package healthcheck

import "time"

type Config struct {
	Addr            string        `yaml:"addr" env:"ADDR" env-default:":8080"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout" env:"SHUTDOWN_TIMEOUT" env-default:"5s"`
}
