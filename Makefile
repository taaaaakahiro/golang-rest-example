run: fmt
	go run ./cmd/api/main.go

clean:
	go clean -testcache

test: clean
	go test ./...

fmt:
	go fmt ./...