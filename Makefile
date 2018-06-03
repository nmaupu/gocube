BIN=bin
BIN_NAME=gocube

all: build

fmt:
	go fmt ./...

deps:
	glide install

build $(BIN)/$(BIN_NAME): $(BIN)
	env CGO_ENABLED=0 go build -o $(BIN)/$(BIN_NAME)

install:
	env CGO_ENABLED=0 go install

clean:
	go clean -i
	rm -rf $(BIN)

test:
	go test $(go list ./... | grep -v /vendor/)

$(BIN):
	mkdir -p $(BIN)

.PHONY: fmt install clean test all release deps
