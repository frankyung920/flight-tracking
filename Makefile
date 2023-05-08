.PHONY: unit-test lint

unit-test:
	go test -race -tags=unit -covermode=atomic -coverprofile=coverage.out ./... 
	go tool cover -html=coverage.out -o ./coverage.html

lint:
	@golangci-lint --color always run --config .golangci.yaml
