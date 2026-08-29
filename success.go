package tracer

import "context"

func CompleteSpan(ctx context.Context) {
	finish(ctx, finishOptions{
		Status: statusSuccess,
	})
}
