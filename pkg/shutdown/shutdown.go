package shutdown

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func Wait(parent context.Context, cfg Configs, stopFunc func(ctx context.Context)) {
	notifyCtx, notifyCancel := signal.NotifyContext(parent, os.Interrupt, syscall.SIGTERM)
	defer notifyCancel()

	<-notifyCtx.Done()

	timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer timeoutCancel()

	stopFunc(timeoutCtx)
}
