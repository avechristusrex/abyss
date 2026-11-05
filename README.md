# Abyss

Abyss is a Go rate-limiter service skeleton designed for Envoy's external rate-limit gRPC integration. Envoy supplies descriptors (request facts); Abyss owns policy and decisions.

## Layout

- `api/proto/ratelimit/v1`: initial rate-limit contract and Buf configuration
- `cmd/abyss`: service entry point
- `internal/{limiter,policy,store,config,metrics,server}`: core boundaries
- `test/{correctness,integration,benchmark}`: test organization
- `deployments/docker`: Valkey development service

The in-memory token bucket is a reference implementation. Valkey and policy resolution are placeholders. The proto currently models the Envoy external rate limit shape; generated code and service registration are intentionally left for the next implementation step.

## Requirements

Go 1.23+, Docker Compose, and Buf (for protobuf generation).

## Development

```sh
make fmt
make test
make build
docker compose -f deployments/docker/compose.yaml up -d
```

HTTP `:8080/healthz` and `:8080/metrics` are exposed; the gRPC transport listens on `:9000`. Configure with `ABYSS_GRPC_ADDR`, `ABYSS_HTTP_ADDR`, and `ABYSS_VALKEY_ADDR`.

## Contract

The protobuf lives at `api/proto/ratelimit/v1/ratelimit.proto`. `buf lint` and `buf generate` use `buf.yaml` and `buf.gen.yaml`. Align generated types with Envoy's canonical service definitions when wiring the adapter.
