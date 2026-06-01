.PHONY: build test verify fmt vet lint modernize modernize-fix tidy update-deps

BINARY :=

build:
	go build -o $(BINARY) .

test:
	go test -race ./...

verify: fmt vet lint modernize

fmt:
	gofmt -l -w .

vet:
	go vet ./...

lint:
	golangci-lint run ./...

modernize:
	go run golang.org/x/tools/gopls/internal/analysis/modernize/cmd/modernize@latest ./...

modernize-fix:
	go run golang.org/x/tools/gopls/internal/analysis/modernize/cmd/modernize@latest -fix ./...

tidy:
	go mod tidy

update-deps:
	go get -u ./...
	go mod tidy
