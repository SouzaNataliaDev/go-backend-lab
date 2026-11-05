.PHONY: test check examples
test:
	go test -race ./...
check:
	go vet ./...
	test -z "$$(gofmt -l .)"
examples:
	go run ./cmd/strategy
	go run ./cmd/resilience
	go run ./cmd/workers
