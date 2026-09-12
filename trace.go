package tracer

import (
	"sync/atomic"

	"github.com/nsonidotdev/gotrail/internal/id"
)

type trace struct {
	id         id.TraceID
	root       *Span
	isFinished atomic.Bool
}

func newTrace(root *Span) *trace {
	// TODO: handle possible error
	id, _ := id.GenerateTraceID()

	return &trace{
		id:   id,
		root: root,
	}
}
