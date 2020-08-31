test:
	go test -race ./...

test-coverage:
	go test -race -coverprofile=coverage.txt -covermode=atomic -coverpkg=./checks,./config

.PHONY: test test-coverage
