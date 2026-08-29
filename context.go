package tracer

import (
	"context"
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
