package tracer

import "context"

func CompleteSpan(ctx context.Context) {
	finish(ctx, finishOptions{
		status: statusSuccess,
	})
}
