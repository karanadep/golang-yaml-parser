.PHONY: default help

help:           ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

build:
	go build -o bin/parser cmd/main.go

start-docker:
	docker build --target dev . -t go
	docker run -it -v ${PWD}:/work go sh

run:
	go run cmd/main.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm cmd/main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 cmd/main.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 cmd/main.go

all: build run
