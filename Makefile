.PHONY: build-all build-win-amd64 build-win-i386 build-linux-amd64 build-linux-i386

SRC_DIR=$(shell pwd)/src
BIN_DIR=$(shell pwd)/bin

build-all: build-win-amd64 build-win-i386 build-linux-amd64 build-linux-i386

build-win-amd64:
	GOOS=windows GOARCH=amd64 go build -ldflags -H=windowsgui -o $(BIN_DIR)/win-amd64.exe $(SRC_DIR)/*.go

build-win-i386:
	GOOS=windows GOARCH=amd64 go build -ldflags -H=windowsgui -o $(BIN_DIR)/win-i386.exe $(SRC_DIR)/*.go

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/linux-amd64 $(SRC_DIR)/*.go

build-linux-i386:
	GOOS=linux GOARCH=386 go build -o $(BIN_DIR)/linux-i386 $(SRC_DIR)/*.go
