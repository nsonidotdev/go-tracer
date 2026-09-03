package tracer

import "sync"

type tracerConfig struct {
	tracker *tracker
}

var (
	initOnce      sync.Once
	config        *tracerConfig
	isInitialized bool
)

func Init() {
	initOnce.Do(func() {
		tracker := newTracker()
		config = &tracerConfig{tracker: tracker}
		isInitialized = true
	})
}
