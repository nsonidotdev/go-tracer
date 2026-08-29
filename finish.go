package tracer

import (
	"context"
	"time"
)

type finishOptions struct {
	Status status
	Reason string
}

func finish(ctx context.Context, opts finishOptions) {
	s, err := getCtxSpan(ctx)
	if err != nil || s == nil {
		// Context carries no span. Return early
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	// TODO: catch context cancel signal (timed out or cancel call)
	// and fail the span

	if s.isFinished() {
		return
	}

	end := time.Now()
	if s.start.IsZero() {
		s.start = time.Now()
	}

	duration := end.Sub(s.start)

	s.end = end
	s.duration = duration
	s.status = opts.Status
	s.reason = opts.Reason
}

func (s *Span) isFinished() bool {
	return s.status == statusFail || s.status == statusSuccess || s.status == statusSkip
}
