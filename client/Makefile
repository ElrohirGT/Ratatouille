BINARY_NAME=ratatouille

DEFAULT_GOAL:=run

build:
	GOARCH=amd64 GOOS=darwin go build -o ./target/${BINARY_NAME}-darwin ./cmd/Ratatouille/main.go
	GOARCH=amd64 GOOS=linux go build -o ./target/${BINARY_NAME}-linux ./cmd/Ratatouille/main.go
	GOARCH=amd64 GOOS=windows go build -o ./target/${BINARY_NAME}-windows ./cmd/Ratatouille/main.go

run: build
	./target/${BINARY_NAME}-linux

clean:
	go clean
	rm ./target/${BINARY_NAME}-darwin
	rm ./target/${BINARY_NAME}-linux
	rm ./target/${BINARY_NAME}-windows

test:
	go test -v ./...

format:
	gofmt -w .

dep:
	go mod download

lint:
	golangci-lint run --enable-all
