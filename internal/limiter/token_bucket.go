package limiter

import (
	"context"
	"errors"
	"math"
	"sync"
	"time"
)

var ErrInvalidPolicy = errors.New("invalid token bucket policy")

type bucket struct {
	tokens float64
	last   time.Time
}
type TokenBucket struct {
	mu      sync.Mutex
	buckets map[string]bucket
	now     func() time.Time
}

func NewTokenBucket() *TokenBucket {
	return &TokenBucket{buckets: make(map[string]bucket), now: time.Now}
}
func (t *TokenBucket) Allow(ctx context.Context, key string, p Policy) (Decision, error) {
	if err := ctx.Err(); err != nil {
		return Decision{}, err
	}
	if key == "" || p.Capacity == 0 || p.RefillRate <= 0 || p.RefillPeriod <= 0 {
		return Decision{}, ErrInvalidPolicy
	}
	now := t.now()
	t.mu.Lock()
	defer t.mu.Unlock()
	b, ok := t.buckets[key]
	if !ok {
		b = bucket{tokens: float64(p.Capacity), last: now}
	}
	elapsed := now.Sub(b.last).Seconds()
	if elapsed > 0 {
		b.tokens = math.Min(float64(p.Capacity), b.tokens+elapsed*p.RefillRate/p.RefillPeriod.Seconds())
	}
	b.last = now
	allowed := b.tokens >= 1
	if allowed {
		b.tokens--
	}
	remaining := uint64(math.Floor(b.tokens))
	reset := time.Duration(0)
	if b.tokens < 1 {
		reset = time.Duration((1 - b.tokens) / p.RefillRate * p.RefillPeriod.Seconds() * float64(time.Second))
	}
	t.buckets[key] = b
	return Decision{Allowed: allowed, Limit: p.Capacity, Remaining: remaining, ResetAfter: reset}, nil
}
