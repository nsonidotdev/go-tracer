package tracer

func (s *Span) Success() {
	s.finish(finishOptions{
		Status: statusSuccess,
	})
}
