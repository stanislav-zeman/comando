.PHONY: build
build:
	go build -o bin/ ./...

.PHONY: install
install:
	go install ./cmd/comando

.PHONY: lint
lint:
	go tool github.com/golangci/golangci-lint/v2/cmd/golangci-lint run

.PHONY: test
test:
	go test -race -timeout 1h -coverprofile cp.out ./...

.PHONY: generate
generate:
	go generate ./
