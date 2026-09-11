package tracer

import (
	"sync"
	"time"
)

type status string

const (
	statusRunning status = "running"
	statusSuccess status = "success"
	statusFail    status = "fail"
	statusSkip    status = "skip"
)

type Span struct {
	id       string
	name     string
	trace    *trace
	duration time.Duration
	start    time.Time
	end      time.Time
	meta     map[string]string
	status   status
	reason   string
	children []*Span
	parent   *Span
	mu       sync.Mutex
	// After func stop() callback
	stopCancelListener func() bool
}

func getRoot(s *Span) *Span {
	if s.parent == nil {
		return s
	}

	return getRoot(s.parent)
}
