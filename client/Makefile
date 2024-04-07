BINARY_NAME=ratatouille

DEFAULT_GOAL:=run

build:
	GOARCH=amd64 GOOS=darwin go build -o ./target/${BINARY_NAME}-darwin ./main.go
	GOARCH=amd64 GOOS=linux go build -o ./target/${BINARY_NAME}-linux ./main.go
	GOARCH=amd64 GOOS=windows go build -o ./target/${BINARY_NAME}-windows ./main.go

run: build
    ifeq ($(shell uname), Darwin)
		./target/${BINARY_NAME}-darwin
    endif
    ifeq ($(shell uname), Linux)
		./target/${BINARY_NAME}-linux
    endif
    ifeq ($(shell uname), Windows_NT)
		./target/${BINARY_NAME}-windows.exe
    endif

clean:
	go clean
	rm -rf ./target

test:
	go test -v ./...

format:
	gofmt -w .

dep:
	go mod download

lint:
	golangci-lint run --enable-all
