package tracer

func (s *Span) Skip(reason string) {
	s.finish(finishOptions{
		Status: statusSuccess,
		Reason: reason,
	})
}
