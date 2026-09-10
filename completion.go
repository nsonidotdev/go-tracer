package tracer

// Entry point for handling completed spans
func handleTraceCompleted(s *Span) {
}

// Checks if every span under the one passed in arguments
// has already finished including the passed span itself
func checkTraceFinished(s *Span) bool {
	spanFinished := s.isFinished()
	childrenFinished := true
	for _, childSpan := range s.children {
		childrenFinished = checkTraceFinished(childSpan)
	}

	return spanFinished && childrenFinished
}
