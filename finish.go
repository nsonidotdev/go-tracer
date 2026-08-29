package tracer

import (
	"context"
	"time"
)

type finishOptions struct {
	status status
	reason string
}

func finish(ctx context.Context, opts finishOptions) {
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
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.isFinished() {
		return
	}

	handleFinish(s, opts)
}

func handleFinish(s *Span, opts finishOptions) {
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
}

func (s *Span) isFinished() bool {
	return s.status == statusFail || s.status == statusSuccess || s.status == statusSkip
}
