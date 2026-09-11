package tracer

import (
	"context"
	"time"

	"github.com/nsonidotdev/go-tracer/internal/id"
)

func StartSpan(parent context.Context, name string, meta map[string]string) context.Context {
	if !isInitialized.Load() || isTerminated.Load() {
		return parent
	}

	parentSpan, _ := getCtxSpan(parent)
	if parentSpan != nil && parentSpan.trace.isFinished.Load() {
		return parent
	}

	id, err := id.GenerateRandomID(10)
	if err != nil {
		return parent
	}

	newSpan := &Span{
		id:     id,
		name:   name,
		start:  time.Now(),
		meta:   meta,
		status: statusRunning,
		parent: parentSpan,
	}

	if parentSpan == nil {
		newSpan.trace = newTrace(newSpan)
	} else {
		newSpan.trace = parentSpan.trace
	}

	config.tracker.recordStart(newSpan)

	if parentSpan != nil {
		parentSpan.mu.Lock()
		parentSpan.children = append(parentSpan.children, newSpan)
		parentSpan.mu.Unlock()
	}

	ctx := context.WithValue(parent, ctxSpanKey, newSpan)
	stop := context.AfterFunc(ctx, func() {
		newSpan.finish(finishOptions{status: statusFail, reason: errSpanCtxCancelled.Error()})
	})
	newSpan.stopCancelListener = stop

	return ctx
}
