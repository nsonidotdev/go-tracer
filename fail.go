package tracer

import "context"

func Fail(ctx context.Context, reason string) {
	finish(ctx, finishOptions{
		Status: statusFail,
		Reason: reason,
	})
}
