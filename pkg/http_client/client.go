package httpclient

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

func New(cfg *Config) (*retryablehttp.Client, error) {
	if cfg == nil {
		return nil, errors.New("httpclient -> config is nil")
	}

	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("httpclient -> failed to validate config -> %w", err)
	}

	client := retryablehttp.NewClient()

	client.HTTPClient.Timeout = cfg.Timeout
	client.RetryMax = cfg.RetriesMax

	if cfg.RetriesMax > 0 {
		client.RetryWaitMin = cfg.RetriesDelay
		client.RetryWaitMax = cfg.RetriesDelay
		client.Backoff = func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
			return cfg.RetriesDelay
		}
	}

	return client, nil
}
