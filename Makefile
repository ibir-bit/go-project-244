.PHONY: lint test

lint:
	golangci-lint run

test:
	go test ./... -v -coverprofile=coverage.out