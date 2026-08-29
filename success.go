package tracer

import "context"

func Success(ctx context.Context) {
	finish(ctx, finishOptions{
		Status: statusSuccess,
	})
}
