GO := go
GOFLAGS := -buildvcs=false -o ./bin/neon

build:
	$(GO) build $(GOFLAGS) ./cmd/neon

clean:
	rm -f neon

format:
	$(GO) fmt ./...

.PHONY: build clean format
