package shutdown

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func Wait(parentCtx context.Context, cfg Config, stopFunc func(ctx context.Context)) {
	notifyCtx, notifyCancel := signal.NotifyContext(parentCtx, os.Interrupt, syscall.SIGTERM)
	defer notifyCancel()

	<-notifyCtx.Done()

	timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer timeoutCancel()

	stopFunc(timeoutCtx)
}
