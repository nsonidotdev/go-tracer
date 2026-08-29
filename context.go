package tracer

import (
	"context"
	"errors"
)

var (
	errCtxSpanMistyped = errors.New("could not parse span")
	errCtxSpanNotFound = errors.New("span not found")
)

func getCtxSpan(ctx context.Context) (*Span, error) {
	span := ctx.Value(ctxSpanKey)
	if span == nil {
		return nil, errCtxSpanNotFound
	}

	parsedSpan, ok := span.(*Span)
	if !ok {
		return nil, errCtxSpanMistyped
	}

	return parsedSpan, nil
}
