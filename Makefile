all: test

test:
	@go test -race -count=1 -v ./... && echo '\n\n--- All tests pass --- --- All tests pass ---\n'

001:
	go run 001/main.go 1337 10

.PHONY: all test 001
