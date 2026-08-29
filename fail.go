package tracer

import "context"

func FailSpan(ctx context.Context, reason string) {
	finish(ctx, finishOptions{
		Status: statusFail,
		Reason: reason,
	})
}
