// Package metrics exposes aggregate Prometheus metrics; it does not record individual requests.
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	Requests = prometheus.NewCounterVec(prometheus.CounterOpts{Name: "rate_limit_requests_total", Help: "Total rate limit checks."}, []string{"domain"})
	Allowed  = prometheus.NewCounterVec(prometheus.CounterOpts{Name: "rate_limit_allowed_total", Help: "Allowed rate limit checks."}, []string{"domain"})
	Rejected = prometheus.NewCounterVec(prometheus.CounterOpts{Name: "rate_limit_rejected_total", Help: "Rejected rate limit checks."}, []string{"domain"})
	Duration = prometheus.NewHistogramVec(prometheus.HistogramOpts{Name: "rate_limit_decision_duration_seconds", Help: "Rate limit decision duration."}, []string{"domain"})
)

func init()                 { prometheus.MustRegister(Requests, Allowed, Rejected, Duration) }
func Handler() http.Handler { return promhttp.Handler() }
