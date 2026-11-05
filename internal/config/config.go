package config

import "os"

type Config struct{ GRPCAddr, HTTPAddr, ValkeyAddr string }

func Load() Config {
	return Config{GRPCAddr: env("ABYSS_GRPC_ADDR", ":9000"), HTTPAddr: env("ABYSS_HTTP_ADDR", ":8080"), ValkeyAddr: env("ABYSS_VALKEY_ADDR", "localhost:6379")}
}
func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
