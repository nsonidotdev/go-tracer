package tracer

func (s *Span) Fail(reason string) {
	s.finish(finishOptions{
		Status: statusFail,
		Reason: reason,
	})
}
