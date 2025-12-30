package httpclient

import (
	"errors"
	"time"
)

type Config struct {
	Timeout      time.Duration
	RetriesMax   int
	RetriesDelay time.Duration
}

func (c *Config) validate() error {
	if c.Timeout <= 0 {
		return errors.New("timeout must be positive")
	}

	if c.RetriesMax < 0 {
		return errors.New("retries max cannot be negative")
	}

	if c.RetriesDelay < 0 {
		return errors.New("retries delay cannot be negative")
	}

	if c.RetriesMax > 0 && c.RetriesDelay == 0 {
		return errors.New("retries delay must be positive when retries are enabled")
	}

	return nil
}
