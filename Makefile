
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

.PHONY: cover
cover:
	$(GO) test -tags $(GOTAGS) -coverprofile=coverage.out ./pkg/...
	$(GO) tool cover -html=coverage.out -o coverage.html
	$(GO) tool cover -func=coverage.out | grep -v 100.0
