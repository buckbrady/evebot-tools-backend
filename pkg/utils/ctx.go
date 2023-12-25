package utils

import (
	"context"
	"time"
)

func NewDBCtx(ctx context.Context, timeout int) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
}
