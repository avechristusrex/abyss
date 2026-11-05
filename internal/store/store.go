// Package store defines state backends for distributed limiter implementations.
package store

import (
	"context"
	"time"
)

type Counter interface {
	Increment(context.Context, string, time.Duration) (uint64, error)
}
