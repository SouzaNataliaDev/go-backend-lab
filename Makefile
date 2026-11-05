.PHONY: test check examples fundamentals
test:
	go test -race ./...
check:
	go vet ./...
	test -z "$$(gofmt -l .)"
examples:
	go run ./cmd/strategy
	go run ./cmd/resilience
	go run ./cmd/workers

fundamentals:
	@set -e; for example in fundamentals/[0-9]*; do go run ./$$example; done
