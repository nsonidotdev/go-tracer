package tracer

import "sync/atomic"

type trace struct {
	root       *Span
	isFinished atomic.Bool
}

func newTrace(root *Span) *trace {
	return &trace{
		root: root,
	}
}
