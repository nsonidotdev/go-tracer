package tracer

import (
	"context"
)

type ctxKey int

const ctxSpanKey ctxKey = iota

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
