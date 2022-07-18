#!/usr/bin/make -f

PROJECT_NAME := cn.example.api.user

build-linux:
	@env GOOS=linux GOARCH=amd64 go build -o bin/$(PROJECT_NAME)-linux

test-app:
	go test -v ./...

install-deps:
	@if [ ! -f go.mod ]; then \
		go mod init; \
	fi
	go mod tidy