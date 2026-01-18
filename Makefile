
GO = go

.PHONY: build
build:
	$(GO) build ./cmd/swhid/...

.PHONY: test
test:
	$(GO) test ./pkg/swhid/...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: fix
fix:
	golangci-lint run --fix ./...
