run:
	go run cmd/main.go

gen-mocks:
	mockery --all --keeptree

