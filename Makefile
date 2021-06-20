run:
	go run cmd/main.go

gen-mocks:
	mockery --all --keeptree

test-coverage:
	go test ./... -coverprofile coverage.out

.PHONY lint:
	golangci-lint run --config .golangci.yml
