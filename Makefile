
GO = go

GOTAGS = git
#GOTAGS = git sql

.PHONY: build
build:
	$(GO) build -tags "$(GOTAGS)" ./cmd/swhid/...

.PHONY: test
test:
	$(GO) test -tags "$(GOTAGS)" ./pkg/swhid/...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: fix
fix:
	golangci-lint run --fix ./...

.PHONY: cover
cover:
	$(GO) test -tags "$(GOTAGS)" -coverprofile=coverage.out ./pkg/...
	$(GO) tool cover -html=coverage.out -o coverage.html
	$(GO) tool cover -func=coverage.out | grep -v 100.0

CC = cc

CFLAGS = -O2 -g -fPIC -shared

sha1.so: ext/misc/sha1.o
	$(CC) -o $@ $(CFLAGS) $^ $(LDFLAGS)

compress.so: LDFLAGS = -lz

compress.so: ext/misc/compress.o
	$(CC) -o $@ $(CFLAGS) $^ $(LDFLAGS)

.PHONY: clean
clean:
	$(RM) ext/misc/*.o
	$(RM) swhid sha1.so compress.so
