package tracer

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
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

		go listenProcTerm()
	})
}

func listenProcTerm() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	<-sigChan

	fmt.Println("received shutdown signal")
	isTerminated.Store(true)

	// clean up
	if config != nil && config.tracker != nil {
		config.tracker.terminateAll(errProcTerminated.Error())
	}

	os.Exit(0)
}
