.PHONY: build install clean

GO = go

all: build

build:
	@$(GO) build -ldflags="-s -w" -trimpath  -o bin/devnv ./cmd/devnv/main.go

install: build
	@install ./bin/devnv $(shell go env GOPATH)/bin

clean:
	@rm -rf ./bin

