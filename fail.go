package tracer

import "context"

func FailSpan(ctx context.Context, reason string) {
	finish(ctx, finishOptions{
		status: statusFail,
		reason: reason,
	})
}
