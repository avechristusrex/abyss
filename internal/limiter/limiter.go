// Package limiter contains rate-limit algorithms and decision types.
package limiter

import (
	"context"
	"time"
)

type Decision struct {
	Allowed          bool
	Limit, Remaining uint64
	ResetAfter       time.Duration
}
type Limiter interface {
	Allow(context.Context, string, Policy) (Decision, error)
}
type Policy struct {
	Capacity     uint64
	RefillRate   float64
	RefillPeriod time.Duration
}
