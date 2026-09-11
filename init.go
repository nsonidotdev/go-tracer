package tracer

import (
	"sync"
	"sync/atomic"
)

type tracerConfig struct {
	tracker *tracker
}

var (
	initOnce      sync.Once
	config        *tracerConfig
	isInitialized atomic.Bool
	isTerminated  atomic.Bool
)

func Init() {
	initOnce.Do(func() {
		tracker := newTracker()
		config = &tracerConfig{tracker: tracker}
		isInitialized.Store(true)
	})
}

// Shutdown terminates all traces. Call this for graceful shutdown
func Shutdown() {
	isTerminated.Store(true)

	if config != nil && config.tracker != nil {
		config.tracker.terminateAll(errProcTerminated.Error())
	}
}
