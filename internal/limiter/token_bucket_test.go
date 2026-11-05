package limiter

import (
	"context"
	"testing"
	"time"
)

func TestTokenBucketCapacity(t *testing.T) {
	l := NewTokenBucket()
	p := Policy{Capacity: 1, RefillRate: 1, RefillPeriod: time.Second}
	a, err := l.Allow(context.Background(), "k", p)
	if err != nil || !a.Allowed {
		t.Fatalf("first request: %+v, %v", a, err)
	}
	b, err := l.Allow(context.Background(), "k", p)
	if err != nil || b.Allowed {
		t.Fatalf("second request should be rejected: %+v, %v", b, err)
	}
}
