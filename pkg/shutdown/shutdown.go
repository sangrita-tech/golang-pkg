package shutdown

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Wait(parentCtx context.Context, delay time.Duration, stopFunc func(shutdownCtx context.Context)) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	go func() {
		defer signal.Stop(ch)

		select {
		case <-ch:
		case <-parentCtx.Done():
		}

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), delay)
		defer shutdownCancel()

		stopFunc(shutdownCtx)
	}()
}
