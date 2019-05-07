PWD := $(shell pwd)

all: build build-win

build:
	@echo "Building split-csv binary to './split-csv'"
	@go build ./cmd/main.go
	@mv $(PWD)/main $(PWD)/split-csv

build-win:
	@echo "Building split-csv binary to './split-csv.exe'"
	@env GOOS=windows GOARCH=amd64 go build ./cmd/main.go
	@mv $(PWD)/main.exe $(PWD)/split-csv.exe