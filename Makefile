GO := go
GOFLAGS := -buildvcs=false -o ./bin/neon

build:
	$(GO) build $(GOFLAGS) ./cmd/neon

clean:
	rm -f neon

format:
	$(GO) fmt ./...

lint:
	$(GO) vet ./...

test:
	$(GO) test -v ./...

.PHONY: build clean format
