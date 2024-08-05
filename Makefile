build:
	@go build -o bin/grader cmd/main.go

test:
	go test -v ./...

run: build
	@./bin/grader
