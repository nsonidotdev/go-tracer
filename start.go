package tracer

import (
	"context"
	"errors"
	"time"

	"github.com/nsonidotdev/go-tracer/internal/id"
)

func StartSpan(parent context.Context, name string, meta map[string]string) (context.Context, error) {
	parentSpan, err := getCtxSpan(parent)
	// isNewTrace := errors.Is(err, errCtxSpanNotFound)

	if err != nil && !errors.Is(err, errCtxSpanNotFound) {
		return nil, err
	}

	id, err := id.GenerateRandomID(10)
	if err != nil {
		return nil, err
	}

	newSpan := &Span{
		id:     id,
		name:   name,
		start:  time.Now(),
		meta:   meta,
		status: statusRunning,
		parent: parentSpan,
	}

	if parentSpan != nil {
		parentSpan.mu.Lock()
		parentSpan.children = append(parentSpan.children, newSpan)
		parentSpan.mu.Unlock()
	}

	ctx := context.WithValue(parent, ctxSpanKey, newSpan)
	return ctx, nil
}
