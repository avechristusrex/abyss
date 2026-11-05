# Architecture

Envoy extracts trusted request facts (identity, device, address, route) and sends descriptors to the external rate-limit service. Abyss resolves those descriptors against local policy and applies a limiter. State is abstracted for a later Valkey-backed distributed implementation.

Prometheus exposes aggregate request, decision, and latency metrics; individual requests are not metric labels or log entries. The current gRPC listener is a compilable transport shell; generate and register Envoy's canonical external rate-limit service before deployment.
