package tracer

import "sync"

type tracker struct {
	// map of span pointers by id key
	activeSpans map[string]*Span
	mu          sync.Mutex
}

func newTracker() *tracker {
	return &tracker{
		activeSpans: make(map[string]*Span),
	}
}

func (t *tracker) recordStart(s *Span) {
	t.mu.Lock()
	defer t.mu.Unlock()
	existing := t.activeSpans[s.id]
	if existing != nil {
		return
	}

	t.activeSpans[s.id] = s
}

func (t *tracker) recordFinish(id string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	delete(t.activeSpans, id)
}

func (t *tracker) terminateAll(reason string) {
	t.mu.Lock()
	activeIDs := make([]string, 0, len(t.activeSpans))
	for id := range t.activeSpans {
		activeIDs = append(activeIDs, id)
	}
	t.mu.Unlock()

	finishOptions := finishOptions{status: statusFail, reason: reason}
	for _, ID := range activeIDs {
		finishByID(ID, finishOptions)
	}
}
