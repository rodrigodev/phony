.PHONY: test run

default: build

run: build
	./app phones.txt

build: test
	go build -o ./app ./cmd

test: lint
	go fmt ./...
	go test -vet all ./...

lint: get-linter
	golangci-lint run

get-linter:
    command -v golangci-lint || curl -sfL "https://install.goreleaser.com/github.com/golangci/golangci-lint.sh" | sh -s -- -b $(go env GOPATH)/bin v1.18.0