package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/khatran/abyss/internal/config"
	"github.com/khatran/abyss/internal/metrics"
	"github.com/khatran/abyss/internal/server"
)

func main() {
	cfg := config.Load()
	mux := http.NewServeMux()
	mux.Handle("/metrics", metrics.Handler())
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok\n"))
	})
	httpServer := &http.Server{Addr: cfg.HTTPAddr, Handler: mux, ReadHeaderTimeout: 5 * time.Second}
	grpcServer := server.NewGRPCServer()
	errCh := make(chan error, 2)
	go func() {
		log.Printf("metrics/health listening on %s", cfg.HTTPAddr)
		errCh <- httpServer.ListenAndServe()
	}()
	go func() {
		log.Printf("gRPC listening on %s", cfg.GRPCAddr)
		errCh <- grpcServer.ListenAndServe(cfg.GRPCAddr)
	}()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	select {
	case s := <-sig:
		log.Printf("shutting down on %s", s)
	case err := <-errCh:
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}
}
