package tracer

import (
	"context"
	"time"
)

type finishOptions struct {
	status status
	reason string
}

func finishByID(id string, opts finishOptions) {
	if !isInitialized.Load() {
		return
	}

	if opts.status == statusRunning {
		return
	}

	config.tracker.mu.Lock()
	s := config.tracker.activeSpans[id]
	// Unlock early because recordFinish will take a lock further
	config.tracker.mu.Unlock()

	if s == nil {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	if s.isFinished() {
		return
	}

	handleFinish(s, opts)
}

func finish(ctx context.Context, opts finishOptions) {
	if !isInitialized.Load() {
		return
	}

	if opts.status == statusRunning {
		return
	}

	s, err := getCtxSpan(ctx)
	if err != nil || s == nil {
		// Context carries no span. Return early
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	if s.isFinished() {
		return
	}

	handleFinish(s, opts)
}

func (s *Span) finish(opts finishOptions) {
	if !isInitialized.Load() {
		return
	}

	if opts.status == statusRunning {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	if s.isFinished() {
		return
	}

	handleFinish(s, opts)
}

// Centralized handler for finishing a span
func handleFinish(s *Span, opts finishOptions) {
	if !isInitialized.Load() {
		return
	}

	if s.stopCancelListener != nil {
		s.stopCancelListener()
		s.stopCancelListener = nil
	}

	end := time.Now()
	if s.start.IsZero() {
		s.start = time.Now()
	}

	duration := end.Sub(s.start)

	s.end = end
	s.duration = duration
	s.status = opts.status
	s.reason = opts.reason

	config.tracker.recordFinish(s.id)
}

func (s *Span) isFinished() bool {
	return s.status == statusFail || s.status == statusSuccess || s.status == statusSkip
}
