package shutdown

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func Wait(parentCtx context.Context, cfg Config, stopFunc func(ctx context.Context)) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(sigCh)

	select {
	case <-parentCtx.Done():
		return
	case <-sigCh:
	}

	timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer timeoutCancel()
	stopFunc(timeoutCtx)
}
