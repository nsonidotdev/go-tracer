package tracer

import "context"

func SkipSpan(ctx context.Context, reason string) {
	finish(ctx, finishOptions{
		Status: statusSuccess,
		Reason: reason,
	})
}
