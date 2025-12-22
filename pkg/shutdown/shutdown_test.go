package shutdown_test

import (
	"context"
	"os"
	"runtime"
	"sync/atomic"
	"syscall"
	"testing"
	"time"

	"github.com/sangrita-tech/golang-pkg/pkg/shutdown"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Wait_ParentCtxCancelled_TriggersStopFunc(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var called atomic.Bool
	delay := 100 * time.Millisecond

	shutdown.Wait(ctx, delay, func(shutdownCtx context.Context) {
		called.Store(true)
	})

	cancel()

	time.Sleep(2 * delay)

	assert.True(t, called.Load())
}

func Test_Wait_SignalReceived_TriggersStopFunc(t *testing.T) {
    if runtime.GOOS == "windows" {
		t.Skip("SIGTERM is not supported on Windows for this test")
	}
    
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var called atomic.Bool
	delay := 100 * time.Millisecond

	shutdown.Wait(ctx, delay, func(shutdownCtx context.Context) {
		called.Store(true)
	})

	time.Sleep(2 * delay)

	p, err := os.FindProcess(os.Getpid())
	require.NoError(t, err)
	require.NoError(t, p.Signal(syscall.SIGTERM))

	assert.True(t, called.Load())
}
