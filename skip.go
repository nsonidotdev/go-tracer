package tracer

import "context"

func Skip(ctx context.Context, reason string) {
	finish(ctx, finishOptions{
		Status: statusSuccess,
		Reason: reason,
	})
}
