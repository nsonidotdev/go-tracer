package tracer

import (
	"time"
)

type finishOptions struct {
	Status status
	Reason string
}

func (s *Span) finish(opts finishOptions) {
	if s.isFinished() {
		// TODO: add warning log to notify that span was already finished
		s.status = opts.Status
		s.reason = opts.Reason
		return
	}

	end := time.Now()
	if s.start.IsZero() {
		// TODO: add warning log to notify that finish was executed
		// with unset `Start` field. This is a bug on developer side
		s.start = time.Now()
	}

	duration := end.Sub(s.start)

	s.end = end
	s.duration = duration
	s.status = opts.Status
	s.reason = opts.Reason

	if s.parent == nil {
		// TODO: remove the span from registry
	}
}

func (s *Span) isFinished() bool {
	return s.status == statusFail || s.status == statusSuccess || s.status == statusSkip
}
