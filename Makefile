GO ?= go
.PHONY: fmt test build run proto lint
fmt:
	$(GO) fmt ./...
test:
	$(GO) test ./...
build:
	$(GO) build ./...
run:
	$(GO) run ./cmd/abyss
proto:
	buf generate
lint:
	buf lint
